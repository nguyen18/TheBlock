import React from "react";
import { Logo, LogoText, Nav, NavLink, NavMenu }
    from "./NavbarElements";

const Navbar = () => {
    return (
        <>
            <Nav>
                <Logo><LogoText href="/">THE BLOCK</LogoText></Logo>
                <NavMenu>
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