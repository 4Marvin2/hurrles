import React from "react";
import SearchLogo from '../imgs/searchBar/search.svg';
import '../css/SearchBar.css';
import { searchRestorsRequest } from "../requests/restors";
import { getRestorMenu } from "../requests/restors";

export default class NavBar extends React.Component {
    searchRestors(payload) {
        searchRestorsRequest(payload.target.value)
        .then((data) => {
            console.log(data)
            const restors = [];
            if (!data) {
                this.props.searchCallback({restors: restors})
                return
            }
            data.forEach(e => {
                getRestorMenu(e.id).then((data) => {
                    const dishes = [];
                    if (data) {
                        data.forEach(e => {
                            const dishElement = {
                                id: e.id,
                                title: e.title,
                                desc: e.description,
                                price: e.price
                            };
                            dishes.push(dishElement);
                        });
                    }

                    const restorElement = {
                        id: e.id, 
                        srcImg: e.img ? e.img : '', 
                        title: e.title, 
                        rating: '4.1', 
                        metro: e.metro,
                        restorInfo: {
                            number: e.number,
                            address: e.address,
                            workTime: `Понедельник-суббота ${e.openTime} - ${e.closeTime}`,
                            desc: e.description,
                            tag1: e.kitchen
                        },
                        dishes: dishes,
                        isFavorite: false,
                    };
                    restors.push(restorElement);
                    this.setState({
                        currentIndex: 0,
                        restors: restors,
                    });
                    this.props.searchCallback({restors: restors})
                });
            });
        })
        .catch((error) => {
            console.log(error)
        })
    }


    render() {
        return (
            <div className="home__search search">
                    <img className="search__img" src={SearchLogo} alt="search" />
                    <input className="search__input" type="text" placeholder='Search' onInput={(e) => {this.searchRestors(e)}}/>
            </div>
        );
    }
}