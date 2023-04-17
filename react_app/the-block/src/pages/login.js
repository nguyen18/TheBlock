import React from "react";
import { Wrapper } from './components/wrapper';
import { Content } from './components/content';
import LoginForm from "./components/login_form";
import './css/login.css'

const LoginPage = () => {
    return (
      <Wrapper>
        <Content>
            <a href="/">back</a>
            <div id="LoginForm">
                <LoginForm></LoginForm>
            </div>
            
        </Content>
      </Wrapper>
    );
  };
    
  export default LoginPage;