import React from "react";
import FilterLogo from '../imgs/navbar/filter.svg';
import SearchLogo from '../imgs/navbar/search.svg';
import '../css/NavBar.css';

export default class NavBar extends React.Component {
    render() {
        return (
            <nav class="home__navbar">
                <div class="home__filter">
                    <img src={FilterLogo} alt="test" />
                </div>
                <div class="home__search search">
                    <img class="search__img" src={SearchLogo} alt="search" />
                    <input class="search__input" type="text" placeholder='Search' />
                </div>
            </nav>
            // <div className="navBar">
            //     <button className="navBar-icons">
            //         <img src={FilterLogo} alt="test" />
            //     </button>
            //     <div className='search-container'>
            //         <img className='search-logo' src={SearchLogo} alt="test" />
            //         <input type="text" placeholder='Search' />
            //     </div>
            // </div>
        );
    }
}