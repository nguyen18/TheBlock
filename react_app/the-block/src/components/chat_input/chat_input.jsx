import React, {Component} from "react";
import './chat_input.scss'

class ChatInput extends Component {
    render() {
        return (
            <div className='ChatInput'>
                <input onKeyDown={this.props.send} placeholder="Type a message... Hit Enter to send"></input>
            </div>
        )
    }
}

export default ChatInput;