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
          Bring your team together and build your community by using our
          cross-platform app that lets you collaborate via chat, voice and by
          sharing and storing unlimited media files. A world of topics is
          waiting for you. Join the private beta.
        </p>
      </Content>
    </Wrapper>
  );
};
  
export default Home;