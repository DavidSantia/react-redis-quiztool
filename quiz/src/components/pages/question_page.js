import React, {Component} from 'react';
import {Panel, Button, Image} from 'react-bootstrap';

class QuestionPage extends Component {
  constructor(props) {
    super(props);
    this.state = {
      answer: "",
      selected: false,
      toggle: true,
      data: [
        {name: "q1", value: "Apple"},
        {name: "q2", value: "Orange"},
        {name: "q3", value: "Watermelon"}
      ]
    }
    console.log("Questions:", this.props.QuestionPage)
  }
  onFormSubmit(event) {
    event.preventDefault();
    let {submitAnswer} = this.props;
    let {answer} = this.state;
    if (answer != "") {
      submitAnswer(answer);
    }
  }

  sleepFor(ms){
    var now = new Date().getTime();
    while(new Date().getTime() < now + ms){ /* do nothing */ }
  }
  onSelect(event, value) {
    event.preventDefault();
    this.setState({answer: value});
    console.log("Radio Select:", value);
  }

  render() {
    let {ready, disable, categories, questions, nextQuestion} = this.props;
    let {answer, data, toggle} = this.state;
    let noSelection = (answer == "");

    let buttons = data.map(q => {
      let selected = (answer == q.name);
      let img_src = "/images/unselected.png";
      if (selected) {
        img_src = "/images/selected.png";
      }
      return (
        <div key={q.name} className="radio">
          <label key={q.name} >
            <Button key={q.name} bsStyle="link" bsSize="xsmall" propType="radio" name="multipleChoice"
                   onClick={event => this.onSelect(event, q.name)} value={q.name}>
              <Image src={img_src} height="24" width="24"  />
            </Button>
            {q.value}
          </label>
        </div>
      );
    });

    return (
      <form onSubmit={event => this.onFormSubmit(event)}>
        <Panel className="form-group" bsStyle="primary" header="Fruit Question">
          <div className="panel-body fixed-panel">
            <label className="control-label">Select the best answer</label>
            {buttons}
            <br/>
          </div>
          <div className="panel-footer">
            <Button bsStyle="primary" type="submit" disabled={disable || noSelection}>Submit</Button>
          </div>
        </Panel>
      </form>);
  }
}

QuestionPage.propTypes = {
  ready: React.PropTypes.bool.isRequired,
  disable: React.PropTypes.bool.isRequired,
  categories: React.PropTypes.string.isRequired,
  questions: React.PropTypes.string.isRequired,
  submitAnswer: React.PropTypes.func.isRequired
}

export default QuestionPage;
