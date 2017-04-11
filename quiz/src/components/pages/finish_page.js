import React, {Component} from 'react';
import {Panel} from 'react-bootstrap';

class FinishPage extends Component {

  // Display thank you message
  render() {
    return (
      <Panel className="panel-success" header="Quiz Complete">
        <div className="panel-body fixed-panel">
          <h3>Thank You</h3>
        </div>
        <div className="panel-footer"><br/></div>
      </Panel>
    );
  }
}

export default FinishPage
