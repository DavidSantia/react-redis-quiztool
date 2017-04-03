import React, { Component } from 'react';
import ReactDOM from 'react-dom';

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      title: "My Quiz",
      questions: 4,
      categories: 2,
      quizdata: {},
      connected: false
    };
  }

  // This section communicates handles the Redis connection
  componentDidMount() {
    // Connect to Redis
    var Redis = require("ioredis");
    var redis = new Redis();

    redis.get('quiz:1 title', function (err, result) {
        console.log("Got", reply);
    });
  }

  render() {
    return (
      <div className="app">
        <div className="row">
          <div className="col-md-3">
            <h2>{this.state.title}</h2>
            <p>Categories: {this.state.categories}</p>
            <p>Questions: {this.state.questions}</p>
          </div>
        </div>
      </div>
    );
  }
}


ReactDOM.render(<App />, document.getElementById('root'));
