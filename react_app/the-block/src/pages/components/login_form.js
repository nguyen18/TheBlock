import React from "react";
import styled from "styled-components";
import { Butt } from "./elements";
import './css/login_form.css'

const LoginForm = () => {
    return (
        <Form_container>
            <p id="LoginTitle">LOGIN</p>
            <form>
                <div>
                    <Input type="text" name="uname" placeholder="enter username..." required />
                </div>
                <div className="input-container">
                    <Input type="password" name="pass" placeholder="enter password..." required />
                </div>
                <div className="button-container">
                    <Butt id="SubmitButton" type="button|submit|reset">Submit</Butt>
                </div>
            </form>
            <Divider/>
            <Butt id="SignupButton" as="a" href="/signup">Sign-Up</Butt>
        </Form_container>
    );
};

const Form_container = styled.div`
    text-align: center;
    font-size: 130%;
    border: solid 1px;
    border-radius: 5px;
    width: 60%;
    padding-top: 15px;
    padding-bottom: 40px;
`

const Input = styled.input`
    padding: 10px;
    margin-bottom: 20px;
    border-radius: 5px;
    width: 70%;
`

const Divider = styled.hr`
    border-top: 1px solid #bbb;
    border-radius: 5px;
    width: 90%;
    margin-bottom: 30px;
`

export default LoginForm;
