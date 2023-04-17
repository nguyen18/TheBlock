import React from "react";
import styled from "styled-components";
import { Butt, Auth_Input, Auth_Form_container, Divider } from "./elements";
import './css/signup_form.css'

const SignupForm = () => {
    return (
        <Auth_Form_container>
            <p id="SignupTitle">SIGN-UP</p>
            <form>
                <div>
                    <Auth_Input type="text" name="uname" placeholder="enter username..." required />
                </div>
                <div className="input-container">
                    <Auth_Input type="password" name="pass" placeholder="enter password..." required />
                </div>
                <div className="input-container">
                    <Auth_Input type="password" name="pass" placeholder="confirm password..." required />
                </div>
                <div className="button-container">
                    <Butt id="SubmitButton" type="button|submit|reset">Submit</Butt>
                </div>
            </form>
            <Divider/>
            <Butt id="SignupButton" as="a" href="/login">Login</Butt>
        </Auth_Form_container>
    );
};

export default SignupForm;
