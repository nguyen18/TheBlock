import React from "react";
import { Wrapper } from './components/wrapper';
import { Content } from './components/content';
import SignupForm from "./components/signup_form";
import './css/signup.css'

const SignupPage = () => {
    return (
      <Wrapper>
        <Content>
        <a href="/">back</a>
            <div id="SignupForm">
                <SignupForm></SignupForm>
            </div>
            
        </Content>
      </Wrapper>
    );
  };
    
  export default SignupPage;