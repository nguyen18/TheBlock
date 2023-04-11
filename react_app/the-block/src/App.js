import React from 'react';
import './css/app.css'
import Spline from '@splinetool/react-spline';
import styled from 'styled-components';
import Navbar from './components/navbar/navbar';
import { BrowserRouter as Router, Routes, Route }
    from 'react-router-dom';
import Home from './pages/index';

function App() {
    return (
        <Router>
            <Navbar />
            <Routes>
                <Route exact path='/' exact element={<Home />} />
            </Routes>
        </Router>

        //<Spline scene="https://prod.spline.design/FNtEG2IT78qP98ex/scene.splinecode" />
    )
}

export default App;