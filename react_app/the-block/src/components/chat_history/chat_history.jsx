import React, {Component} from "react";
import './chat_history.scss';
import Message from '../message/message';

class ChatHistory extends Component{
    render(){
        console.log(this.props.ChatHistory);
        this.props.ChatHistory.map(msg=><Message key={msg.timeStamp} message={msg.data}/>);

        return(
            <div className='ChatHistory'>
                <h2>Chat History</h2>
                {messages}
            </div>
        );
    };
}

export default ChatHistory;