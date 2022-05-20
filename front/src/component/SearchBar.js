import React from "react";
import SearchLogo from '../imgs/searchBar/search.svg';
import '../css/SearchBar.css';

export default class NavBar extends React.Component {
    render() {
        return (
            <div className="home__search search">
                    <img className="search__img" src={SearchLogo} alt="search" />
                    <input className="search__input" type="text" placeholder='Search' />
            </div>
        );
    }
}