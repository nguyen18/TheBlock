import React from 'react';
import { BrowserRouter as Router, Routes, Route }
    from 'react-router-dom';
import AboutPage from './about';
import Home from './home';
import TeamPage from './team';
import TechStackPage from './tech_stack';

export const Routepusher = () => {
    return (
        <Router>
            <Routes>
                <Route exact path='/' element={<Home />} />
                <Route path='/about' element={<AboutPage />} />
                <Route path='/team' element={<TeamPage />} />
                <Route path='/techstack' element={<TechStackPage />} />
            </Routes>
        </Router>
    )
}
