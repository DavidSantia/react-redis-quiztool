import React, {Component} from 'react';
import {Image} from 'react-bootstrap';

class Footer extends Component {
  render() {
    let {connected, text} = this.props;
    let img_src = "/images/disconnected.png";
    let status = " Not Connected";
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
  connected: React.PropTypes.bool.isRequired
}

export default Footer
