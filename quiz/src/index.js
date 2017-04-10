import React, { Component } from 'react';
import {PageHeader} from 'react-bootstrap';
import ReactDOM from 'react-dom';
import DefaultPage from './components/pages/default_page';
import StartPage from './components/pages/start_page';
import QuestionPage from './components/pages/question_page';
import QuizRoutes from './quiz_routes'

class QuizTool extends Component {
  constructor(props) {
    super(props);
    this.state = {
      connected: false,
      began: false,
      quizId: 0,
      title: "",
      categories: "0",
      questions: "0",
      quizData: {},
      showModal: false
    };
    this.currentPage = "Loading...";
    this.header = "Welcome to QuizTool";
  }
  componentDidMount() {
    this.defaultPage();
  }

  // The following Page hooks get called by client router
  defaultPage() {
    console.log("Default page");
    this.currentPage = (<DefaultPage />);
    this.forceUpdate();
  }
  quizDetailsPage(quizId) {
    this.header = this.state.title + " Quiz";
    console.log("Details Page for Quiz", quizId);
    this.currentPage = (
      <StartPage
        {...this.state}
        setRootState={data => this.setRootState(data)}
        questionPage={(quizId, qNum) => this.questionPage(quizId, qNum)}/>
    );
    this.forceUpdate();
  }
  questionPage(quizId, qNum) {
    console.log("viewPage Quiz:", quizId, " Question:" + qNum);
    this.currentPage = (
      <QuestionPage
        {...this.state}
        submitAnswer={(answer) => this.submitAnswer(answer)}/>
    );
    this.forceUpdate();
  }
  
  submitAnswer(answer) {
    console.log("Submitted answer: ", answer);
  }

  setRootState(data) {
    //console.log("Set root state: ", data);
    this.setState(data);
  }

  render() {
    return (
      <div className="app">
        <div className="row">
          <div className="col-sm-12">
            <PageHeader>{this.header}</PageHeader>
            {this.currentPage}
            <QuizRoutes
              {...this.state}
              setRootState={(data) => this.setRootState(data)}
              defaultPage={() => this.defaultPage()}
              quizDetailsPage={(id) => this.quizDetailsPage(id)}
              questionPage={(id, qNum) => this.questionPage(id, qNum)}/>
          </div>
        </div>
      </div>
    );
  }
}

ReactDOM.render(<QuizTool />, document.getElementById('root'));
