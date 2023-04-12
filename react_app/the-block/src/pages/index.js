// path: '/'
import React from 'react';
import { Wrapper } from '../components/wrapper';
import { Content } from '../components/content';
import Spline from '@splinetool/react-spline';
  
const Home = () => {
  return (
    <Wrapper>
        <Spline scene="https://prod.spline.design/FNtEG2IT78qP98ex/scene.splinecode" />
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
