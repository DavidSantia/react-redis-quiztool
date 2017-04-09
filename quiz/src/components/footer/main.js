import React, {Component} from 'react';
import {Image} from 'react-bootstrap';

// Displays connection status icon on the left, and pagination text on the right
class Footer extends Component {
  render() {
    let {connected, text} = this.props;

    // Disconnected status
    let img_src = "/images/disconnected.png";
    let status = " Not Connected";

    // Connected status
    if (connected) {
      status = " Connected to Server";
      img_src = "/images/connected.png";
    }

    return (
      <div className="footer navbar-default navbar-fixed-bottom">
        <div className="container-fluid footer-container">
	        <Image src={img_src} height="24" width="24" rounded />
          {status}
          <span className="pull-right">{text}</span>
        </div>
      </div>);
  }
}

Footer.propTypes = {
  connected: React.PropTypes.bool.isRequired,
  text: React.PropTypes.string.isRequired
}

export default Footer
