import React, { Component } from 'react';
import {Panel, Button} from 'react-bootstrap';
import Footer from './components/footer/main';
import Socket from './socket';

class QuizRoutes extends Component {
  constructor(props) {
    super(props);
    this.state = {
      routes: {
        '/default': () => props.defaultPage(),
        '/quiz/:quizId/details': (quizId) => props.quizDetails(quizId),
        '/quiz/:quizId/:qNum': (quizId, qNum) => props.questionPage(quizId, qNum)
      },
      data: {},
      currentQ: 0,
      done: false
    };

    // Initialize router
    this.router = Router(this.state.routes);
    this.router.init();

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
  setConnected() {
    console.log("Redis server Ready");
    this.props.setRootState({connected: true});
  }
  onDisconnect() {
    console.log("Connection closed");
    this.props.setRootState({connected: false});
  }
  onSuccess(data) {
    let quizId = 1;
    if (data == "PONG") {
      // Server replied to PING, so is connected
      this.setConnected();

      // Request quiz meta-data
      this.socket.send("HGETALL", "quiz:" + quizId);
    } else if (data.title != null) {
      // Got quiz meta-data
      this.props.setRootState(data);

      // Display quiz details start page
      this.props.quizDetails(quizId);
    } else {
      console.log("Got Reply: ", data);
    }
  }
  onError(data) {
    console.log("Got Error:", data);
  }

  getQuestionNumber() {
    let {questions} = this.props;
    let {currentQ, done} = this.state;
    if (done) {
      return "Finished Quiz";
    }
    if (currentQ > 0 && parseInt(questions, 10) > 0) {
      return "Question " + String(currentQ) + " of " + questions;
    }
    return "";
  }

  incrementPage() {
    let {questions} = this.props;
    let q = this.state.currentQ;
    q = q + 1;

    if (q <= parseInt(questions, 10)) {
      this.setState({currentQ: q});
    } else {
      this.setState({done: true});
      console.log("Done with quiz");
      return;
    }
  }

  render() {
    let footer_text = this.getQuestionNumber();

    return(
      <Footer connected={this.props.connected} text={footer_text} />
    );
  }
}

QuizRoutes.propTypes = {
  connected: React.PropTypes.bool.isRequired,
  categories: React.PropTypes.string.isRequired,
  questions: React.PropTypes.string.isRequired,
  setRootState: React.PropTypes.func.isRequired,
  defaultPage: React.PropTypes.func.isRequired,
  quizDetails: React.PropTypes.func.isRequired,
  questionPage: React.PropTypes.func.isRequired
}

export default QuizRoutes;
