import React, {Component} from 'react';
import {Panel, Button, Image} from 'react-bootstrap';

class QuestionPage extends Component {
  constructor(props) {
    super(props);
    this.state = {
      answer: "",
      selected: false,
      toggle: true,
      category: "Favorites",
      correctAnswers: "",
      question: "What is your favorite fruit?",
      newData: {
        "Category": "Terminology",
        "CorrectAnswers": "2",
        "Question": "What is the scientific study of plant life?",
        "A1": "Dryad Science",
        "A2": "Botany",
        "A3": "Nutrition",
        "A4": "Greenery",
        "A5": "Plantology"
      },
      multipleChoice: [
        {name: "q1", value: "Apple"},
        {name: "q2", value: "Orange"},
        {name: "q3", value: "Watermelon"}
      ]
    }
  }

  // onComponentDidMount(){
  //   let {newData} = this.state;
  //   let multipleChoice = [];
  //
  //   // Gather category, question, correct answer, and multiple choices
  //   for (var k in newData) {
  //     if (newData.hasOwnProperty(k)) {
  //       if (k == "Category") {
  //         this.setState({category: newData[k]});
  //       } else if (k == "CorrectAnswers") {
  //         this.setState({correctAnswers: newData[k]});
  //       } else if (k == "Question") {
  //         this.setState({question: newData[k]});
  //       } else {
  //         multipleChoice.push({name: k, value: newData[k]});
  //       }
  //     }
  //     this.setState({multipleChoice});
  //     console.log("State is:", this.state);
  //   }
  // }
  
  onFormSubmit(event) {
    event.preventDefault();
    let {submitAnswer} = this.props;
    let {answer} = this.state;
    if (answer != "") {
      submitAnswer(answer);
    }
  }

  onSelect(event, value) {
    event.preventDefault();
    this.setState({answer: value});
    console.log("Radio Select:", value);
  }

  render() {
    let {ready, connected, categories, questions, nextQuestion} = this.props;
    let {answer, multipleChoice, category, question} = this.state;
    let disabled = (!connected || answer == "");

    // Generate multiple-choice radio buttons
    let buttons = multipleChoice.map(q => {
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
        <Panel className="form-group" bsStyle="primary" header={category}>
          <div className="panel-body fixed-panel">
            <label className="control-label">{question}</label>
            {buttons}
            <br/>
          </div>
          <div className="panel-footer">
            <Button bsStyle="primary" type="submit" disabled={disabled}>Submit</Button>
          </div>
        </Panel>
      </form>);
  }
}

QuestionPage.propTypes = {
  connected: React.PropTypes.bool.isRequired,
  categories: React.PropTypes.string.isRequired,
  questions: React.PropTypes.string.isRequired,
  submitAnswer: React.PropTypes.func.isRequired
}

export default QuestionPage;
