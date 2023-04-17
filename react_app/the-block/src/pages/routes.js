import React from 'react';
import { BrowserRouter as Router, Routes, Route }
    from 'react-router-dom';
import AboutPage from './about';
import Home from './home';
import TeamPage from './team';
import TechStackPage from './tech_stack';
import LoginPage from './login';
import SignupPage from './signup';

export const Routepusher = () => {
    return (
        <Router>
            <Routes>
                <Route exact path='/' element={<Home />} />
                <Route path='/about' element={<AboutPage />} />
                <Route path='/team' element={<TeamPage />} />
                <Route path='/techstack' element={<TechStackPage />} />
                <Route path='/login' element={<LoginPage />} />
                <Route path='/signup' element={<SignupPage />} />
            </Routes>
        </Router>
    )
}
