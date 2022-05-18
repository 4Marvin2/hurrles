import React from 'react'
import Restors from './Restors'
import SearchBar from  './SearchBar'
import RestorOpen from './RestorOpen/RestorOpen'
import '../css/Favor.css'

import { getFavoriteRestors, getRestorMenu } from '../requests/restors';

var BreakException = {};

export default class Favor extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            currentIndex: 0,
            isFavorite: false,
            restors:[],
        }
        getFavoriteRestors().then((data) => {
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
                            rating: '4.8',
                            address: e.address,
                            workTime: `Понедельник-суббота ${e.openTime} - ${e.closeTime}`,
                            desc: e.description,
                            tag1: e.kitchen
                        },
                        dishes: dishes,
                    };
                    restors.push(restorElement);
                    this.setState({
                        currentIndex: 0,
                        restors: restors,
                    });
                    getFavoriteRestors().then((data) => {
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
                });
            });
        });

        this.restorClick = this.restorClick.bind(this);
    }

    restorClick(payload) {
        getFavoriteRestors().then((data) => {
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

    render(){
        return (
            <div className='favor'>
                <div className='favor__left'>
                    <SearchBar />
                    <Restors restors={this.state.restors} restorClick={this.restorClick} />
                </div>
                <div className='favor__right'>
                    {this.state.restors.length !== 0 && 
                        <div className='favor__restorOpen_in'>
                            <RestorOpen
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
