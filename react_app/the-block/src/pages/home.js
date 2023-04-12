// path: '/'
import React from 'react';
import { Wrapper } from './components/wrapper';
import { Content } from './components/content';
import Spline from '@splinetool/react-spline';
import Navbar from './components/navbar/navbar';
import './css/home.css';
  
const Home = () => {
  return (
    <Wrapper>
        <Navbar></Navbar>
      <Content>
        <h1>Collaborate with people</h1>
        <p>
        imagine a world where it was easy to make friends. a place where you can meet people––not curated profiles.
        </p>
      </Content>
    </Wrapper>
  );
};
  
export default Home;
