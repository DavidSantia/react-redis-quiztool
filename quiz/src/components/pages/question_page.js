import React, {Component} from 'react';
import {Panel, Button, Radio} from 'react-bootstrap';

class QuestionPage extends Component {
  constructor(props) {
    super(props);
    this.state = {
      answer: "",
      selected: false,
      data: [
        {name: "q1", value: "Apple"},
        {name: "q2", value: "Orange"},
        {name: "q3", value: "Watermelon"}
      ]
    }
  }
  onFormSubmit(event) {
    event.preventDefault();
    let {submitAnswer} = this.props;
    let {answer} = this.state;
    if (answer != "") {
      submitAnswer(answer);
    }
  }
  onSelect(event) {
    event.preventDefault();
    console.log("Radio Select:", event);
  }

  render() {
    let {ready, disable, categories, questions, nextQuestion} = this.props;
    let {answer, data} = this.state;
    let noSelection = (answer == "");
    let buttons = data.map(q => {
        return (
          <Radio key={q.name} name="multiChoice" type="radio" value={q.name}> {q.value} </Radio>
        );
    });
    console.log("Answer =", answer);
    //onChange={event => this.onSelect(event)}

    return (
      <form onSubmit={event => this.onFormSubmit(event)}>
        <Panel bsStyle="primary" header="Fruit Question" className="form-group">
          <div className="radio">
            {buttons}
           </div>
          <br/>
          <Button bsStyle="primary" type="submit" disabled={disable || noSelection}>Submit</Button>
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
