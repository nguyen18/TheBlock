import styled from "styled-components";
import '../../assets/fonts/FontFaces.css'
import React, { useState } from 'react';
import { ReactComponent as Clear } from '../../assets/images/svg/clear.svg';

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
