import React, { Component } from 'react';
import ReactDOM from 'react-dom';
import Socket from './socket';


class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      quizId: 1,
      title: "",
      questions: 0,
      categories: 0,
      quizData: {},
      connected: false,
      showModal: false
    };
  }

  // This section communicates handles the Redis connection
  componentDidMount() {
    let socket = this.socket = new Socket();

    // Route Redis responses
    socket.on("success", (data) => this.onSuccess(data));
    socket.on("error", (data) => this.onError(data));

    // Route internal actions
    socket.on("connect", () => this.onConnect());
    socket.on("disconnect", () => this.onDisconnect());
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
          <div className="col-md-3">
            <h2>{this.state.title}</h2>
            <p>Categories: {this.state.categories}</p>
            <p>Questions: {this.state.questions}</p>
            <img src="/images/connected.png" />
          </div>
        </div>
      </div>
    );
  }
}

ReactDOM.render(<App />, document.getElementById('root'));
