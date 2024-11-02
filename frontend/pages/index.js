import React, { useState } from 'react';
import NavBar from '../components/NavBar';


export default function Header() {

    return (
        <div className='main_grid'>
            <div class="logo">
                <h1>EASY</h1>
            </div>

            <div className='page'>Page</div>
            
            <NavBar />

            

        </div>
    );
}    