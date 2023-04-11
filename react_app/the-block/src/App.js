import React from 'react';
import './css/app.css'
import Spline from '@splinetool/react-spline';
import styled from 'styled-components';
import Navbar from './components/navbar/navbar';
import { BrowserRouter as Router, Routes, Route }
    from 'react-router-dom';
import Home from './pages/index';
import { Wrapper } from './components/wrapper';

function App() {
    return (
        <Router>
            <Navbar />
            <Routes>
                <Route exact path='/' element={<Home />} />
            </Routes>
        </Router>
    )
}

export default App;