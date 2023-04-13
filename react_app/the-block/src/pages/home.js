// path: '/'
import React from 'react';
import { Wrapper } from './components/wrapper';
import { Content } from './components/content';
//import Spline from '@splinetool/react-spline';
import Navbar from './components/navbar/navbar';
import './css/home.css';
  
const Home = () => {
  return (
    <Wrapper>
        <Navbar></Navbar>
      <Content>
        <h1 id='WelcomeText'>
        imagine a world where it was easy to make human connections. a place 
        where you can meet people––not curated profiles.        </h1>
      </Content>
    </Wrapper>
  );
};
  
export default Home;
