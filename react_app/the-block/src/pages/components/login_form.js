import React from "react";
import styled from "styled-components";
import { Butt, Input, Form_container, Divider } from "./elements";
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

export default LoginForm;
