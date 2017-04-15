import React, {Component} from 'react';
import {Button,Modal} from 'react-bootstrap';

// Displays connection status icon on the left, and pagination text on the right
class Answer extends Component {
  render() {
    let {show, text, nextQ} = this.props;

    return (
      <div className="modal-container" style={{height: 200}}>
        <Modal
          show={show}
          onHide={nextQ}
          container={this}
          aria-labelledby="contained-modal-title">
          <Modal.Header closeButton>
            <Modal.Title id="contained-modal-title">Answer</Modal.Title>
          </Modal.Header>
          <Modal.Body>
            {text}
          </Modal.Body>
          <Modal.Footer>
            <Button onClick={nextQ}>Continue</Button>
          </Modal.Footer>
        </Modal>
      </div>);
  }
}

Answer.propTypes = {
  show: React.PropTypes.bool.isRequired,
  text: React.PropTypes.object.isRequired,
  nextQ: React.PropTypes.func.isRequired
}

export default Answer
