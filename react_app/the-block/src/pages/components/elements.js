import styled from "styled-components";
import '../../assets/fonts/FontFaces.css'

/* Buttons */

export const Butt = styled.button`
    text-decoration: none;
    font-family: "SemiBold";
    font-size: 16px;
    border: 0px;
    padding: 1.5% 5% 1.5% 5%;
    border-radius: 14px;
    color: #20321D;
    border: 1px solid rgba(255, 255, 255, 0.1);
    background: #A2C795;

    transition: .5s;
    cursor: pointer;
    pointer-events: auto;

    :hover {
        border: 1px solid rgba(32,50,29, 0.8);
        box-shadow: 0px 2px 2px rgba(0, 0, 0, 0.2);
        transform: translateY(-3px);
`;

/* elements for login and signup pages */

export const Auth_Form_container = styled.div`
    text-align: center;
    font-size: 140%;
    border-radius: 20px;
    width: 60%;
    padding-top: 15px;
    padding-bottom: 40px;
    background-color: white;
    color: black;
    font-family: "Semibold";
    box-shadow: -12px 10px 0px 0px rgba(0, 0, 0, .25);
`

export const Auth_Input = styled.input`
    font-family: "Semibold";
    font-size: 15px;
    padding: 12px;
    margin-bottom: 20px;
    border-radius: 10px;
    border: none;
    width: 70%;
    background-color: #A2C795;
    color: 1px solid rgba(32,50,29, 0.8);
    text-align: center;

`
export const Divider = styled.hr`
    border-top: 0.5px solid #bbb;
    border-radius: 5px;
    width: 90%;
    margin-bottom: 30px;
    border-color: #8D90B5;
`