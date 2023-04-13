import React from "react";
import { Logo, Nav, NavLink, NavMenu }
    from "./NavbarElements";

const Navbar = () => {
    return (
        <>
            <Nav>
                <Logo><p>THE BLOCK</p></Logo>
                <NavMenu>
                    <NavLink to="/" activeStyle>
                        home
                    </NavLink>
                    <NavLink to="/about" activeStyle>
                        about
                    </NavLink>
                    <NavLink to="/team" activeStyle>
                        team
                    </NavLink>
                    <NavLink to="/techstack" activeStyle>
                        tech stack
                    </NavLink>
                </NavMenu>
            </Nav>
        </>
    );
};

export default Navbar;