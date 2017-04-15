import React, { Component } from 'react';
import {PageHeader} from 'react-bootstrap';
import ReactDOM from 'react-dom';
import DefaultPage from './components/pages/default_page';
import FinishPage from './components/pages/finish_page';
import StartPage from './components/pages/start_page';
import QuestionPage from './components/pages/question_page';
import QuizRoutes from './quiz_routes'
import Socket from './socket';

class QuizTool extends Component {
  constructor(props) {
    super(props);
    this.state = {
      began: false,
      quizId: "1",
      title: "",
      totalCorrect: 0,
      categories: "0",
      questions: "0"
    };
    this.currentPage = "Loading...";
    this.header = "Welcome to QuizTool";

    // Initialize websocket for Redis
    this.socket = new Socket();
  }
  componentDidMount() {
    this.defaultPage();
  }

  // The following Page hooks get called by client router
  defaultPage() {
    console.log("[defaultPage]");
    this.currentPage = (<DefaultPage />);
    this.forceUpdate();
  }
  quizDetailsPage(quizId) {
    this.header = this.state.title + " Quiz";
    console.log("[quizDetailsPage] quiz:", quizId);
    this.currentPage = (
      <StartPage
        {...this.state}
        setRootState={data => this.setRootState(data)}
        questionPage={(quizId, qNum) => this.questionPage(quizId, qNum)}/>
    );
    this.forceUpdate();
  }
  questionPage(quizId, qNum) {
    console.log("[questionPage] quiz:", quizId, " question:" + qNum);
    this.currentPage = (
      <QuestionPage
        {...this.state}
        totalQs={this.state.questions}
        socket={this.socket}
        submitAnswer={(answer, correctA, text) => this.submitAnswer(answer, correctA, text)}
        finishPage={() => this.finishPage()}/>
    );
    this.forceUpdate();
  }
  finishPage() {
    console.log("[finishPage]");
    let {totalCorrect, questions} = this.state;
    let text = (<ul>
        <h4><li>You answered {totalCorrect} out of {questions} questions correctly</li></h4>
      </ul>);
    this.currentPage = (<FinishPage text={text} />);
    this.forceUpdate();
  }
  
  submitAnswer(answer, correctA) {
    let {totalCorrect} = this.state;
    if (answer == correctA) {
      this.setState({totalCorrect: totalCorrect + 1});
    }
    console.log("[Submitted answer: " + answer +", Correct answer: " + correctA + "]");
  }

  setRootState(data) {
    //console.log("Set root state: ", data);
    this.setState(data);
  }

  onConnect() {
    console.log("Connecting to Redis server");
    this.socket.send("PING", null);
  }
  onDisconnect() {
    console.log("Connection closed");
    this.setState({connected: false});
  }

  gotConnected() {
    console.log("Redis server Ready");
    this.setState({connected: true});

    this.socket.send("HGETALL", "quiz:" + this.state.quizId);
  }

  onSuccess(data) {
    if (data == "PONG") {
      // Server replied to PING, so is ready
      this.gotConnected();
    } else if (data.title != null) {
      // Server replied to HGETALL for quiz meta-data, so set in state
      this.setState(data)
    } else {
      console.log("Got Reply: ", data);
    }
  }
  
  onError(data) {
    console.log("Got Error:", data);
  }


  render() {
    return (
      <div className="app">
        <div className="row">
          <div className="col-sm-12">
            <PageHeader>{this.header}</PageHeader>
            {this.currentPage}
            <QuizRoutes
              {...this.state}
              socket={this.socket}
              totalQs={this.state.questions}
              setRootState={(data) => this.setRootState(data)}
              defaultPage={() => this.defaultPage()}
              quizDetailsPage={(id) => this.quizDetailsPage(id)}
              questionPage={(id, qNum) => this.questionPage(id, qNum)}/>
          </div>
        </div>
      </div>
    );
  }
}

ReactDOM.render(<QuizTool />, document.getElementById('root'));

