import React from "react";
import FilterLogo from '../../imgs/navbar/filter.svg';
import SearchLogo from '../../imgs/navbar/search.svg';
import '../../css/Restors/NavBar.css';

export default class NavBar extends React.Component {
    render() {
        return (
            <nav className="home__navbar">
                <div className="home__filter">
                    <img src={FilterLogo} alt="test" />
                </div>
                <div className="home__search search">
                    <img className="search__img" src={SearchLogo} alt="search" />
                    <input className="search__input" type="text" placeholder='Search' />
                </div>
            </nav>
        );
    }
}