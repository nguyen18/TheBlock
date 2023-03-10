import React, {Component} from "react";
import './message.scss';

class Message extends Conponent{
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