import React, {Component} from 'react';
import {Panel, Button} from 'react-bootstrap';

class StartPage extends Component {

  onBegin(event) {
    event.preventDefault();
    let {setAppState} = this.props;
    
    setAppState({began: true});
    console.log("Start button pressed");
  }
  
  render() {
    let {ready, disable, categories, questions} = this.props;
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
            <li><strong>Questions</strong>: {questions}</li>
            <li><strong>Categories</strong>: {categories}</li>
          </ul>
          <Button bsStyle="primary" disabled={disable} onClick={event => this.onBegin(event)}>Begin</Button>
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
  setAppState: React.PropTypes.func.isRequired
}

export default StartPage
