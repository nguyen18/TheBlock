import { FaBars } from "react-icons/fa";
import { NavLink as Link } from "react-router-dom";
import styled from "styled-components";
import "../../../assets/fonts/FontFaces.css";
  
export const Nav = styled.nav`
  background: transparent;
  height: 85px;
  display: flex;
  justify-content: space-between;
  padding: 0.2rem calc((100vw - 1000px) / 2);
  z-index: 12;
`;
  
export const NavLink = styled(Link)`
  color: #808080;
  display: flex;
  align-items: center;
  text-decoration: none;
  padding: 0 1rem;
  height: 100%;
  cursor: pointer;

  :hover{
    text-decoration: underline;
  }
  &.active {
    color: #7BB666/*#4d4dff*/;
  }
`;
  
export const Bars = styled(FaBars)`
  display: none;
  color: #808080;
  @media screen and (max-width: 768px) {
    display: block;
    position: absolute;
    top: 0;
    right: 0;
    transform: translate(-100%, 75%);
    font-size: 1.8rem;
    cursor: pointer;
  }
`;
  
export const Logo = styled.div`
  display: flex;
  align-items: center;
  float: left;
  margin-left: 2%;
  height: 85px;
  font-family: "SemiBold";
  color: black;
  @media screen and (max-width: 768px) {
    margin-left: 8%;
  }
`
export const LogoText = styled.a`
  text-decoration: none;
  color: black;

  :hover {
    color: #7BB666/*#4d4dff*/;
  }
`

export const NavMenu = styled.div`
  display: flex;
  align-items: center;
  margin-right: 2%;
  float: right;
  font-family: "Medium";
  @media screen and (max-width: 768px) {

  }
`;