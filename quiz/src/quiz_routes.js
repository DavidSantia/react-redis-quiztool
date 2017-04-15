import React, { Component } from 'react';
import {Panel, Button} from 'react-bootstrap';
import Footer from './components/footer/footer';

class QuizRoutes extends Component {
  constructor(props) {
    super(props);
    this.state = {
      routes: {
        '/default': () => props.defaultPage(),
        '/quiz/:quizId/details': (quizId) => props.quizDetailsPage(quizId),
        '/quiz/question/next': () => props.quizDetailsPage(quizId),
        '/quiz/:quizId/:qNum': (quizId, qNum) => props.questionPage(quizId, qNum)
      },
      data: {},
      currentQ: 0
    };

    // Initialize router
    let options = {notfound: () => props.defaultPage()};
    this.router = Router(this.state.routes).configure(options);
    this.router.init();

    console.log("Setting socket");
  }

  componentDidMount() {
    this.dumpRoutes();
    let {socket} = this.props;

    // Route Redis responses
    socket.on("success", (data) => this.handleInit(data));
    socket.on("error", (data) => this.onError(data));
    // Route internal actions
    socket.on("connect", () => this.onConnect());
    socket.on("disconnect", () => this.onDisconnect());
  }
  
  dumpRoutes() {
    console.log("Available routes:");
    let base = window.location.href.split("#")[0];
    for (var route in this.state.routes) {
      let url = route.replace(/:[a-zA-Z]+/g, "1");
      console.log("• " + base + "#" + url);
    }
  }

  onConnect() {
    let {socket} = this.props;
    console.log("[Connecting to server]");
    socket.send("PING", null);
  }
  onDisconnect() {
    console.log("Connection closed");
    this.props.setRootState({questions: "0"});
  }
  onError(data) {
    console.log("Got Error:", data);
  }

  handleInit(data) {
    let {socket} = this.props;
    let quizId = "1";
    if (data == "PONG") {
      // Server replied to PING, so is connected
      console.log("[Redis server Ready]");

      // Request quiz meta-data
      socket.send("HGETALL", "quiz:" + quizId);
    } else if (data.title != null) {
      // Got quiz meta-data
      this.props.setRootState(data);

      // Display quiz details page
      this.props.quizDetailsPage(quizId);
    } else {
      console.log("Error, unexpected: ", data);
    }
  }

  render() {
    let {questions} = this.props;
    let connected = (questions != "0");
    
    let footer_text = "";
    if (connected) {
      footer_text = "Total Questions: " + questions;
    }

    return(
      <Footer connected={connected} text={footer_text} />
    );
  }
}

QuizRoutes.propTypes = {
  categories: React.PropTypes.string.isRequired,
  totalQs: React.PropTypes.string.isRequired,
  socket: React.PropTypes.object.isRequired,
  setRootState: React.PropTypes.func.isRequired,
  defaultPage: React.PropTypes.func.isRequired,
  quizDetailsPage: React.PropTypes.func.isRequired,
  questionPage: React.PropTypes.func.isRequired
}

export default QuizRoutes;
