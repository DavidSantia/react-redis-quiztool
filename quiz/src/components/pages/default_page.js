import React, {Component} from 'react';
import {Panel} from 'react-bootstrap';

class DefaultPage extends Component {

  // Default message until server connects
  render() {
    return (
      <Panel header="Quiz not available">
        <div className="panel-body fixed-panel">Waiting for connection to server...</div>
        <div className="panel-footer"><br/></div>
      </Panel>
    );
  }
}

export default DefaultPage
