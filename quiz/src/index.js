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
      header: "Welcome to QuizTool",
      currentPage: (<DefaultPage />),
      connected: false,
      began: false,
      quizId: 0,
      title: "",
      categories: "0",
      questions: "0",
      quizData: {},
      showModal: false
    };
  }
  componentDidMount() {
    this.defaultPage();
  }

  defaultPage() {
    console.log("Default page");
    this.setState({currentPage: (<DefaultPage />)});
  }
  quizDetails(quizId) {
    let header = this.state.title + " Quiz";
    this.setState({header, quizId});
    console.log("Details Page for Quiz", quizId);
    let currentPage = (
      <StartPage
        {...this.state}
        setRootState={data => this.setRootState(data)}
        questionPage={(quizId, qNum) => this.questionPage(quizId, qNum)}/>
    );
    this.setState({currentPage});
  }
  questionPage(quizId, qNum) {
    console.log("viewPage Quiz:", quizId, " Question:" + qNum);
    let currentPage = (
      <QuestionPage
        {...this.state}
        submitAnswer={(answer) => this.submitAnswer(answer)}/>
    );
    this.setState({currentPage});
  }
  
  submitAnswer(answer) {
    console.log("Submitted answer: ", answer);
  }

  setRootState(data) {
    //console.log("Set root state: ", data);
    this.setState(data);
  }

  render() {
    let {header, currentPage} = this.state;
    return (
      <div className="app">
        <div className="row">
          <div className="col-sm-12">
            <PageHeader>{header}</PageHeader>
            {currentPage}
            <QuizRoutes
              {...this.state}
              setRootState={(data) => this.setRootState(data)}
              defaultPage={() => this.defaultPage()}
              quizDetails={(id) => this.quizDetails(id)}
              questionPage={(id, qNum) => this.questionPage(id, qNum)}/>
          </div>
        </div>
      </div>
    );
  }
}

ReactDOM.render(<QuizTool />, document.getElementById('root'));
