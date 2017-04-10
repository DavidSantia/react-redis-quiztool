import React, { Component } from 'react';
import ReactDOM from 'react-dom';
import {Panel, Button} from 'react-bootstrap';
import Socket from './socket';

class QuizRoutes extends Component {
  constructor(props) {
    super(props);
    this.state = {
      connected: false,
      ready: false,
      routes: {
        '/default': () => this.default(),
        '/quiz/:quizId/details': (quizId) => this.details(quizId),
        '/quiz/:quizId/:questionId': (quizId, questionId) => this.viewPage(quizId, questionId)
      },
      quizId: 1,
      data: {}
    };

    // Initialize router
    this.router = Router(this.state.routes);
    this.router.init();

    // Initialize Redis to send and receive commands
    this.socket = new Socket();
  }

  default () {
    console.log("Default page");
  }
  details(quizId) {
    console.log("Details for Quiz", quizId);
  }
  viewPage(quizId, questionId) {
    console.log("viewPage Quiz:", quizId, " Question: " + questionId);
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

  render() {
    let {connected, ready, routes} = this.state;
    let ctext = "false"; if (connected) {ctext = "true";}
    let rtext = "false"; if (ready) {rtext = "true";}

    let page = (
      <ul>
        <li>Connected: {ctext}</li>
        <li>Ready: {rtext}</li>
      </ul>
    );

    let list = [];
    let i = 1;
    for (var route in routes) {
      let url = route.replace(/:[a-zA-Z]+/g, "1");
      url = url.replace(/^\//i, "#/");
      let k = "route" + String(i);
      list.push(
        <li key={k}>{k}: <a key={k} href={url}>{route}</a></li>
      );
      i++;
    }
    // list.push(<User key={id} id={id} name={users[id]} {...this.props} />);


    return (
      <div className="app">
        <div className="row">
          <div className="col-sm-12">
            <h4>Status</h4>
            {page}
            <br/>
            <h4>Route test links</h4>
            <ul>
              {list}
            </ul>
          </div>
        </div>
      </div>
    );
  }
}

ReactDOM.render(<QuizRoutes />, document.getElementById('root'));
