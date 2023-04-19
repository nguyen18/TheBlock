import React from "react";
import { Wrapper } from './components/wrapper';
import { Content } from './components/content';
import LoginForm from "./components/login_form";
import './css/login.css'

const LoginPage = () => {
    return (
      <Wrapper id="body">
        <Content>
            <a href="/">home</a>
            <div id="LoginForm">
                <LoginForm></LoginForm>
            </div>
            
        </Content>
      </Wrapper>
    );
  };
    
  export default LoginPage;