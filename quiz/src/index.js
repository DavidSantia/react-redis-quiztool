import React, { Component } from 'react';
import {PageHeader} from 'react-bootstrap';
import ReactDOM from 'react-dom';
import StartPage from './components/start_page/main';
import Footer from './components/footer/main'
import Socket from './socket';

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      ready: false,
      connected: false,
      quizId: 1,
      title: "",
      categories: "",
      questions: "",
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
  startQuiz() {
    console.log("Start button pressed");
  }
  onError(data) {
    console.log("Got Error:", data);
  }

  render() {
    let {title, ready, connected} = this.state;
    let header = title + " Quiz";
    if (!ready) {
      header = "Welcome to QuizTool";
    }
    return (
      <div className="app">
        <div className="row">
          <div className="col-sm-12">
            <PageHeader>{header}</PageHeader>
            <StartPage
              {...this.state}
              disable={!connected}
              startQuiz={() => this.startQuiz()} />
            <Footer connected={connected} />
          </div>
        </div>
      </div>
    );
  }
}

ReactDOM.render(<App />, document.getElementById('root'));
