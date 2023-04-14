import React from "react";
import { Wrapper } from './components/wrapper';
import { Content } from './components/content';
import './css/login.css'

const LoginPage = () => {
    return (
      <Wrapper>
        <Content>
            <p>login page</p>
            <div id="LoginForm">
                <p id="LoginText">Log In</p>
            </div>
        </Content>
      </Wrapper>
    );
  };
    
  export default LoginPage;