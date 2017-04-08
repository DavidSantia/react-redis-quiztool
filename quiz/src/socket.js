import {EventEmitter} from 'events';

// Connection to Redis Websocket
class Socket {
  constructor() {
    let ws = this.ws = new WebSocket('ws://localhost:4000');
    this.ee = new EventEmitter();
    ws.onmessage = event => this.message(event);
    ws.onopen = event => this.open(event);
    ws.onclose = event => this.close(event);
  }
  message(event) {
    const payload = event.data;
    try {
      const msg = JSON.parse(payload);
      console.log("DEBUG Received from Redis: ", payload);
      this.ee.emit(msg.command, msg.data);
    }
    catch(err) {
      this.ee.emit("error", err);
    }
  }
  open(event) {
    this.ee.emit("connect");
  }
  close(event) {
    this.ee.emit("disconnect");
    console.log("Websocket closed");
  }
  send(name, data) {
    // Send message to the server
    const msg = {command: name, data: data};
    const payload = JSON.stringify(msg);
    console.log("DEBUG Sending to Redis: ", payload);
    this.ws.send(payload);
  }
  on(name, fn) {
    // Route message from the server
    this.ee.on(name, fn);
  }
}

export default Socket;
