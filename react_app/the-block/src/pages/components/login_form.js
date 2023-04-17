import React from "react";
import styled from "styled-components";
import { Butt, Auth_Input, Auth_Form_container, Divider } from "./elements";
import './css/login_form.css'

const LoginForm = () => {
    return (
        <Auth_Form_container>
            <p id="LoginTitle">LOGIN</p>
            <form>
                <div>
                    <Auth_Input type="text" name="uname" placeholder="enter username..." required />
                </div>
                <div className="input-container">
                    <Auth_Input type="password" name="pass" placeholder="enter password..." required />
                </div>
                <div className="button-container">
                    <Butt id="SubmitButton" type="button|submit|reset">Submit</Butt>
                </div>
            </form>
            <Divider/>
            <Butt id="SignupButton" as="a" href="/signup">Sign-Up</Butt>
        </Auth_Form_container>
    );
};

export default LoginForm;
