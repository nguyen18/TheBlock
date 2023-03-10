import React, {Component} from "react";
import Header from './components/header/header';
import ChatHistory from "./components/chat_history/chat_history";
import ChatInput from "./components/chat_input/chat_input";
import './App.css';
import {connect, sendMsg} from './api';

class App extends Component {
  constructor(props){
    super(props);
    this.state = {
      ChatHistory: []
    }
  }

  componentDidMount() {
    connect((msg) => {
      console.log("New Message")
      this.setState(prevState => ({
        chatHistory : [...prevState.chatHistory, msg]
      }))
      console.log(this.state);
    })
  }

  render() {
    return(
      <div className="App">
        <Header/>
        <ChatHistory chatHistory={this.state.chatHistory}/>
        <ChatInput send={this.send}/>
      </div>
    )
  }
}

export default App;