import React, {Component} from 'react';
import {Panel, Button, Image} from 'react-bootstrap';
import Answer from '../answer/answer';

class QuestionPage extends Component {
  constructor(props) {
    super(props);
    this.state = {
      answer: "",
      modalText: (<br/>),
      showModal: false,
      selected: false,
      done: false,
      category: "Connection Issue",
      questions: "0",
      correctA: "",
      correctText: "",
      explanation: "",
      question: "Waiting for server...",
      currC: 0,
      currQ: 0,
      mulChoice: []
    }
  }

  componentDidMount() {
    let {totalQs, quizId, socket} = this.props;
    if (totalQs == "0") {
      return;
    }
    socket.on("nextquestion", () => this.incrementQuestion());

    // Initialize for first question
    this.setState({currQ: 1});

    // Get first category
    this.incrementCategory();
  }

  incrementCategory() {
    let {quizId, categories, socket, finishPage} = this.props;
    let {currC} = this.state;

    // Clear previous route
    socket.removeAll("success");
    currC++;
    if (currC <= categories) {
      // Request category info
      socket.on("success", (data) => this.getCategory(data));
      this.setState({currC, currQ: 1});
      socket.send("HGETALL", "quiz:" + quizId + ":c:" + currC);
    } else {
      this.setState({done: true});
      console.log("[Completed Quiz]");
      finishPage();
    }
  }

  getCategory(data) {
    let {quizId, socket} = this.props;
    let {currC} = this.state;
    //console.log("DEBUG Category: " + data["category"]);
    this.setState(data);

    // Clear category route
    socket.removeAll("success");

    // Request first question
    socket.on("success", (data) => this.getQuestion(data));
    this.setState({currQ: 1});
    socket.send("HGETALL", "quiz:" + quizId + ":c:" + currC + ":q:1");
  }

  // Question sample data
  // data = {
  //    "Category": "Terminology",
  //    "CorrectAnswers": "2",
  //    "Question": "What is the scientific study of plant life?",
  //    "A1": "Dryad Science",
  //    "A2": "Botany"
  //  };

  getQuestion(data) {
    let {questions, currQ, category} = this.state;
    let correctA = "";
    let correctText = "";
    let question = "";
    let explanation = "";
    let mulChoice = [];
    //console.log("DEBUG Displaying", currQ, "of", questions, "in Category '" + category +"'");

    if (currQ <= questions) {
      // Extract category, question, correct answer, and multiple choices
      for (var k in data) {
        if (k == "Category") {
          // redundant field
        } else if (k == "CorrectAnswers") {
          correctA = "A" + data[k];
        } else if (k == "Explanation") {
          explanation = data[k];
        } else if (k == "Question") {
          question = data[k];
        } else {
          mulChoice.push({name: k, value: data[k]});
        }
      }
      // Second pass to lookup correct answer text
      for (var k in data) {
        if (k == correctA) {
          correctText = data[k];
        }
      }
      //console.log("DEBUG Question: " + question);
      this.setState({correctA, correctText, explanation, question, mulChoice});
    } else {
      console.log("Error: Unexpected state");
    }
  }

  onSelect(event, value) {
    event.preventDefault();
    this.setState({answer: value});
    //console.log("Radio Select:", value);
  }

  onFormSubmit(event) {
    event.preventDefault();
    let {submitAnswer} = this.props;
    let {answer, correctA, explanation} = this.state;
    let text = (<div><h4 className="text-success">Correct</h4>
          <p className="text-info">{explanation}</p></div>);
    if (answer != "") {
      if (answer != correctA) {
        text = (<div><h4 className="text-danger">Incorrect</h4>
          <p>The correct answer is {this.state.correctText}</p>
          <p className="text-info">{explanation}</p></div>);
      }
      submitAnswer(answer, correctA, text);

      // Clear response after submitting, and show Modal
      this.setState({answer: "", modalText: text, showModal: true});
    }
  }

  onNextQ() {
    this.setState({showModal: false});
    this.incrementQuestion();
  }

  incrementQuestion() {
    let {quizId, socket} = this.props;
    let {currC, currQ, questions} = this.state;
    currQ++;

    if (currQ <= questions) {
      this.setState({currQ});
      socket.send("HGETALL", "quiz:" + quizId + ":c:" + currC + ":q:" + currQ);
    } else {
      //console.log("DEBUG Finished Category");
      this.incrementCategory();
    }
  }

  getFooterText() {
    let {totalQs} = this.props;
    let {currQ, questions, done} = this.state;
    if (done) {
      return "Finished Quiz";
    }
    if (currQ > 0 && questions > 0) {
      return "Question " + String(currQ) + " of " + questions;
    }
    return "";
  }


  render() {
    let {ready, categories, totalQs, nextQuestion} = this.props;
    let {answer, mulChoice, category, question} = this.state;
    let disabled = (totalQs == "0" || answer == "");
    let footer_text = this.getFooterText();

    // Generate multiple-choice radio buttons
    let buttons = mulChoice.map(q => {
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
      <div>
        <form onSubmit={event => this.onFormSubmit(event)}>
          <Panel className="form-group" bsStyle="primary" header={category}>
            <div className="panel-body fixed-panel">
              <label className="control-label">{question}</label>
              {buttons}
              <br/>
            </div>
            <div className="panel-footer">
              <div className="btn btn-outline" disabled>{footer_text}</div>
              <span className="pull-right">
                <Button bsStyle="primary" type="submit" disabled={disabled}>Submit</Button>
              </span>
            </div>
          </Panel>
        </form>
        <Answer
          show={this.state.showModal}
          text={this.state.modalText}
          nextQ={() => this.onNextQ()}/>
      </div>);
  }
}

QuestionPage.propTypes = {
  quizId: React.PropTypes.string.isRequired,
  categories: React.PropTypes.string.isRequired,
  totalQs: React.PropTypes.string.isRequired,
  socket: React.PropTypes.object.isRequired,
  submitAnswer: React.PropTypes.func.isRequired,
  finishPage: React.PropTypes.func.isRequired
}

export default QuestionPage;
