// path: '/'
import React from 'react';
import { Wrapper } from './components/wrapper';
import { Content } from './components/content';
//import Spline from '@splinetool/react-spline';
import Navbar from './components/navbar/navbar';
//import './css/home.css';
  
const TeamPage = () => {
  return (
    <Wrapper>
        <Navbar></Navbar>
      <Content>
      </Content>
    </Wrapper>
  );
};
  
export default TeamPage;