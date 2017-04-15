import React, {Component} from 'react';
import {Panel} from 'react-bootstrap';

class FinishPage extends Component {

  // Display thank you message
  render() {
    let {text} = this.props;
    
    return (
      <Panel className="panel-success" header="Quiz Complete">
        <div className="panel-body fixed-panel">
          <h3>Thank You</h3>
          {text}
        </div>
        <div className="panel-footer"><br/></div>
      </Panel>
    );
  }
}

FinishPage.propTypes = {
  text: React.PropTypes.object.isRequired
}

export default FinishPage
