import React, { Component } from 'react';
import {PageHeader} from 'react-bootstrap';
import ReactDOM from 'react-dom';
import StartPage from './components/pages/start_page';
import QuestionPage from './components/pages/question_page';
import Footer from './components/footer/main'
import Socket from './socket';

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      connected: false,
      ready: false,
      began: false,
      done: false,
      quizId: 1,
      title: "",
      categories: "",
      questions: "",
      currentQ: 0,
      quizData: {},
      showModal: false
    };

    // Initialize Redis to send and receive commands
    this.socket = new Socket();
  }

  componentDidMount() {
    // Route Redis responses
    this.socket.on("success", (data) => this.onSuccess(data));
    this.socket.on("error", (data) => this.onError(data));
    // Route internal actions
    this.socket.on("connect", () => this.onConnect());
    this.socket.on("disconnect", () => this.onDisconnect());
  }

  onConnect() {
    console.log("Connecting to Redis server");
    this.socket.send("PING", null);
  }
  onDisconnect() {
    console.log("Connection closed");
    this.setState({connected: false});
  }
  onSuccess(data) {
    if (data == "PONG") {
      // Server replied to PING, so is connected
      console.log("Redis server Ready");
      this.setState({connected: true});
      // Request quiz meta-data
      this.socket.send("HGETALL", "quiz:" + this.state.quizId);
    } else if (data.title != null) {
      // Got quiz meta-data, so quiz is ready
      this.setState(data);
      this.setState({ready: true});
    } else {
      console.log("Got Reply: ", data);
    }
  }
  onError(data) {
    console.log("Got Error:", data);
  }
  setAppState(data) {
    console.log("In set App state: ", data);
    this.setState(data);
  }

  getQuestionNumber() {
    let {ready, currentQ, questions, done} = this.state;
    if (done) {
      return "Finished Quiz";
    }
    if (currentQ > 0 && parseInt(questions, 10) > 0 && ready) {
      return "Question " + String(currentQ) + " of " + questions;
    }
    return "";
  }
  submitAnswer(answer) {
    let q = this.state.currentQ;
    q = q + 1;
    if (q <= parseInt(this.state.questions, 10)) {
      this.setState({currentQ: q});
    } else {
      this.setState({done: true});
      console.log("Done with quiz");
      return;
    }
    console.log("Submitted answer: ", answer);
  }

  render() {
    let {title, connected, ready, began, currentQ} = this.state;
    // Set header and footer
    let header = title + " Quiz";
    let footer_text = this.getQuestionNumber();
    if (!ready) {
      header = "Welcome to QuizTool";
      footer_text = "";
    }
    // Set start or question page
    let page = "";
    if (began) {
      page = (
        <QuestionPage
          {...this.state}
          disable={!connected}
          submitAnswer={(answer) => this.submitAnswer(answer)} />
      );
    } else {
      page = (
        <StartPage
          {...this.state}
          disable={!connected}
          setAppState={data => this.setAppState(data)}/>
      );
    }
    return (
      <div className="app">
        <div className="row">
          <div className="col-sm-12">
            <PageHeader>{header}</PageHeader>
            {page}
            <Footer connected={connected} text={footer_text} />
          </div>
        </div>
      </div>
    );
  }
}

ReactDOM.render(<App />, document.getElementById('root'));
