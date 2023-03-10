import React, {Component} from "react";
import './message.css';

class Message extends Component{
    constructor(props){
        super(props);
        let temp= JSON.parse(this.props.message);
        this.state={
            message: temp
        }
    }
//study what we are rendering here
    render() {
        return(
            <div className='Message'>
                {this.state.message.body}
            </div>
        );
    };
}

export default Message;