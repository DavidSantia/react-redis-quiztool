import React, {Component} from 'react';
import {Panel, Button} from 'react-bootstrap';

class StartPage extends Component {

  // If we want a quiz timer, add it here
  onBegin(event) {
    event.preventDefault();
    let {setRootState, questionPage, quizId} = this.props;
    setRootState({began: true});
    questionPage(quizId, 1);
    console.log("[Begin Quiz]");
  }

  // Summarize quiz details once connected
  render() {
    let {ready, categories, questions} = this.props;
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
          <Button bsStyle="primary" disabled={questions == "0"}
                  onClick={event => this.onBegin(event)}>Begin Quiz</Button>
        </div>
      </Panel>
    );
  }
}

StartPage.propTypes = {
  quizId: React.PropTypes.string.isRequired,
  categories: React.PropTypes.string.isRequired,
  questions: React.PropTypes.string.isRequired,
  setRootState: React.PropTypes.func.isRequired,
  questionPage: React.PropTypes.func.isRequired
}

export default StartPage
