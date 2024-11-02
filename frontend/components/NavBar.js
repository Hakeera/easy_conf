import React, { useState } from 'react';

export default function NavBar() {
    const [isVisible, setIsVisible] = useState(true);

    const toggleVisibility = () => {
        setIsVisible(!isVisible);
    };

    return (
        <div className={`navbar ${isVisible ? 'visible' : 'hidden'}`}>
            {isVisible && (
                <div>
                    <div className="navbar_item">
                        <h1>Clientes</h1>
                    </div>
                    <div className="navbar_item">
                        <h1>Produtos</h1>
                    </div>
                    <div className="navbar_item">
                        <h1>Pedidos</h1>
                    </div>
                    <div className="navbar_item">
                        <h1>Contas</h1>
                    </div>
                    <div className="navbar_item">
                        <h1>Fornecimento</h1>
                    </div>
                    <div className="navbar_item">
                        <h1>Fiscal</h1>
                    </div>
                </div>
            )}
            <button className="toggle-button" onClick={toggleVisibility}>
                {isVisible ? '▲' : '▼'}
            </button>
        </div>
    );
}
