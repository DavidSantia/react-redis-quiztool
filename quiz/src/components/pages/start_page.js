import React, {Component} from 'react';
import {Panel, Button} from 'react-bootstrap';

class StartPage extends Component {

  // Default message until server connects
  componentDidMount() {
    this.body = (
      <Panel header="Quiz not available">
        <div className="panel-body fixed-panel">Waiting for connection to server...</div>
        <div className="panel-footer"></div>
      </Panel>
    );
  }

  // Summarize quiz details once connected
  onReady() {
    let {ready, disable, categories, questions} = this.props;
    return (
      <Panel bsStyle="primary" header="Quiz Details">
        <div className="panel-body fixed-panel">
          <h3><ul>
            <li>Questions: {questions}</li>
            <br/>
            <li>Categories: {categories}</li>
          </ul></h3>
        </div>
        <div className="panel-footer">
          <Button bsStyle="primary" disabled={disable}
                  onClick={event => this.onBegin(event)}>Begin Quiz</Button>
        </div>
      </Panel>
    );
  }

  // If we want a quiz timer, add it here
  onBegin(event) {
    event.preventDefault();
    let {setAppState} = this.props;
    setAppState({began: true, currentQ: 1});
    console.log("[Begin Quiz]");
  }

  render() {
    if (this.props.ready) {
      this.body = this.onReady();
    }
    return (
      <div>
        {this.body}
      </div>);
  }
}

StartPage.propTypes = {
  ready: React.PropTypes.bool.isRequired,
  disable: React.PropTypes.bool.isRequired,
  categories: React.PropTypes.string.isRequired,
  questions: React.PropTypes.string.isRequired,
  setAppState: React.PropTypes.func.isRequired
}

export default StartPage
