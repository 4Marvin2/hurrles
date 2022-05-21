import React from 'react'
import Restors from './Restors'
import SearchBar from  './SearchBar'
import RestorOpen from './RestorOpen/RestorOpen'
import '../css/Home.css'

import { getRestors, getRestorMenu, getFavoriteRestors, addFavorite, deleteFavorite } from '../requests/restors';

var BreakException = {};

export default class Home extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            currentIndex: 0,
            isFavorite: false,
            restors:[],
        }

        this.likeClick = this.likeClick.bind(this);
        this.restorClick = this.restorClick.bind(this);

        getFavoriteRestors().then((data) => {
            if (!data) {
                return
            }
            try {
                if (this.state.restors.length === 0) {
                    throw BreakException;
                }
                const id = this.state.restors[0].id
                data.forEach(e => {
                    if (e.id === id) {
                        this.setState({isFavorite: true});
                        throw BreakException;
                    };
                });
                this.setState({isFavorite: false});
              } catch (e) {
                if (e !== BreakException) throw e;
              }
        });

        getRestors().then((data) => {
            if (!data) {
                return
            }
            const restors = [];
            data.forEach(e => {
                getRestorMenu(e.id).then((data) => {
                    const dishes = [];
                    data.forEach(e => {
                        const dishElement = {
                            id: e.id,
                            title: e.title,
                            desc: e.description,
                            price: e.price
                        };
                        dishes.push(dishElement);
                    });

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
                });
            });
        });
    }

    restorClick(payload) {
        getFavoriteRestors().then((data) => {
            if (!data) {
                return
            }
            const id = this.state.restors[this.state.currentIndex].id
            try {
                data.forEach(e => {
                    if (e.id === id) {
                        this.setState({isFavorite: true});
                        throw BreakException;
                    };
                });
                this.setState({isFavorite: false});
              } catch (e) {
                if (e !== BreakException) throw e;
              }
        });
        this.setState({currentIndex: payload});
    }

    likeClick(payload) {
        const id  = this.state.restors[this.state.currentIndex].id
        if (!id) {
            return
        }
        if (!payload) {
            deleteFavorite(id).then((data) => {
                this.setState({isFavorite: payload});
            });
        } else {
            addFavorite(id).then((data) => {
                this.setState({isFavorite: payload});
            });
        }
    }

    render(){
        return (
            <div className='home'>
                <div className='home__left'>
                    <SearchBar />
                    <Restors restors={this.state.restors} restorClick={this.restorClick} />
                </div>
                <div className='home__right'>
                    {this.state.restors.length !== 0 && 
                        <div className='home__restorOpen_in'>
                            <RestorOpen
                            likeClick={this.likeClick}
                            isFavorite={this.state.isFavorite}
                            dishes={this.state.restors[this.state.currentIndex].dishes}
                            restorInfo={this.state.restors[this.state.currentIndex].restorInfo} />
                        </div>
                    }
                </div>
            </div>
        );
    }
}
