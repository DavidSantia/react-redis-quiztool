import React, {Component} from 'react';
import {Image} from 'react-bootstrap';

class Footer extends Component {
  render() {
    let {connected} = this.props;
    let img_src = "/images/disconnected.png";
    let text = " Not";
    if (connected) {
      text = "";
      img_src = "/images/connected.png";
    }
    return (
      <div className="footer navbar-default navbar-fixed-bottom">
        <div className="container-fluid">
	        <Image src={img_src} height="24" width="24" thumbnail />
          {text} Connected to Server
        </div>
      </div>);
  }
}

Footer.propTypes = {
  connected: React.PropTypes.bool.isRequired
}

export default Footer
