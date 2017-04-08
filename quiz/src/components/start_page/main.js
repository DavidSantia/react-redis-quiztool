import React, {Component} from 'react';
import {Panel, Button} from 'react-bootstrap';

class StartPage extends Component {
  render() {
    let {ready, disable, categories, questions, startQuiz} = this.props;
    let body = (
      <Panel header="Waiting for connection to server">
        <div>
          <br/>
          <br/>
          <br/>
          <br/>
        </div>
      </Panel>
    );
    if (ready) {
      body = (
        <Panel bsStyle="primary" header="Quiz Details">
          <ul>
            <li><strong>Categories</strong>: {categories}</li>
            <li><strong>Questions</strong>: {questions}</li>
          </ul>
          <Button bsStyle="primary" disabled={disable} onclick={startQuiz}>Begin</Button>
        </Panel>
      );
    }
    return (
      <div>
        {body}
      </div>);
  }
}

StartPage.propTypes = {
  ready: React.PropTypes.bool.isRequired,
  disable: React.PropTypes.bool.isRequired,
  categories: React.PropTypes.string.isRequired,
  questions: React.PropTypes.string.isRequired,
  startQuiz: React.PropTypes.func.isRequired
}

export default StartPage
