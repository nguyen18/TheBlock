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
                    <NavLink to="/contact" activeStyle>
                        team
                    </NavLink>
                    <NavLink to="/blogs" activeStyle>
                        tech stack
                    </NavLink>
                </NavMenu>
            </Nav>
        </>
    );
};

export default Navbar;