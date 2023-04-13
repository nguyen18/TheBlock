import styled from "styled-components";

export const Content = styled.div`
    color: black;
    display: grid;
    margin-left: 10%;
    margin-right: 10%;
    z-index: 12;
    padding: 0.2rem calc((100vw - 1000px) / 2);

    .spline {
        position: absolute;
        margin: 0;
        top: 0;
        right: 0;
        width: 1200px;
        height: 1000px;
    
        @media (max-width: 1024px) {
          transform: scale(0.8) translateX(200px);
          transform-origin: top;
        }
        @media (max-width: 800px) {
          transform: scale(0.7) translateX(600px);
        }
        @media (max-width: 600px) {
          transform: scale(0.5) translateX(-100px);
          right: auto;
          left: 50%;
          margin-left: -600px;
        }
        @media (max-width: 375px) {
          transform: scale(0.45) translateX(-50px);
        }
`;
