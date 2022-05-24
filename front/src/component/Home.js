import React from 'react'
import Restors from './Restors'
import SearchBar from  './SearchBar'
import RestorOpen from './RestorOpen/RestorOpen'
import Reserve from './Reserve/Reserve'
import '../css/Home.css'

import { getRestors, getRestorMenu, getFavoriteRestors, addFavorite, deleteFavorite } from '../requests/restors';
import { dateToTimeString } from '../utils/date'

var BreakException = {};

export default class Home extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            reserve: false,
            currentIndex: 0,
            isFavorite: false,
            restors:[],
        }

        this.likeClick = this.likeClick.bind(this);
        this.restorClick = this.restorClick.bind(this);
        this.reserveClick = this.reserveClick.bind(this);
        this.searchCallback = this.searchCallback.bind(this);

        // getFavoriteRestors().then((data) => {
        //     if (!data) {
        //         return
        //     }
        //     try {
        //         if (this.state.restors.length === 0) {
        //             throw BreakException;
        //         }
        //         const id = this.state.restors[0].id
        //         data.forEach(e => {
        //             if (e.id === id) {
        //                 this.setState({isFavorite: true});
        //                 throw BreakException;
        //             };
        //         });
        //         this.setState({isFavorite: false});
        //       } catch (e) {
        //         if (e !== BreakException) throw e;
        //       }
        // });
        
        getRestors().then((data) => {
            if (!data) {
                return
            }
            const restors = [];
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

                    const startWorkTimeStr = dateToTimeString(e.openTime)
                    const finishWorkTimeStr = dateToTimeString(e.closeTime)
                    const fullWorkTimeStr = `${startWorkTimeStr} - ${finishWorkTimeStr}`
                    const restorElement = {
                        id: e.id, 
                        srcImg: e.img ? e.img : '', 
                        title: e.title, 
                        rating: '4.1', 
                        metro: e.metro,
                        restorInfo: {
                            number: e.number,
                            address: e.address,
                            workTime: `Понедельник-суббота ${fullWorkTimeStr}`,
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

    reserveClick(payload) {
        this.setState({reserve: payload});
    }

    shouldComponentUpdate(newProps, newStat) {
        console.log(this.props, newProps, this.state, newStat)
        if (this.state !== newStat) {
            this.state.searchUpdate = false
            return true
        }
        return false
    }

    searchCallback(payload) {
        this.setState({
            currentIndex: 0,
            restors: payload.restors,
        })
        setTimeout(() => {
            console.log(this.state.restors.length)
        }, 100);
        // getFavoriteRestors().then((data) => {
        //     if (!data) {
        //         return
        //     }
        //     try {
        //         if (this.state.restors.length === 0) {
        //             throw BreakException;
        //         }
        //         const id = this.state.restors[0].id
        //         data.forEach(e => {
        //             if (e.id === id) {
        //                 this.setState({isFavorite: true});
        //                 throw BreakException;
        //             };
        //         });
        //         this.setState({isFavorite: false});
        //     } catch (e) {
        //         if (e !== BreakException) throw e;
        //     }
        // });
        setTimeout(() => {
            console.log(this.state.restors)
        }, 100);
    }

    render(){
        return (
            <div className='home'>
                <div className='home__left'>
                    <SearchBar searchCallback={this.searchCallback}/>
                    <Restors restors={this.state.restors} restorClick={this.restorClick} />
                </div>
                <div className='home__right'>
                    {this.state.restors.length !== 0 && 
                        <div className='home__restorOpen_in'>
                            <RestorOpen
                            reserveClick={this.reserveClick}
                            likeClick={this.likeClick}
                            isFavorite={this.state.isFavorite}
                            dishes={this.state.restors[this.state.currentIndex].dishes}
                            restorInfo={this.state.restors[this.state.currentIndex].restorInfo} />
                        </div>
                    }
                    {this.state.reserve &&
                        <div className='home__reserve'>
                            <Reserve reserveClick={this.reserveClick} id={this.state.restors[this.state.currentIndex].id} userId={this.props.userId} />
                        </div>
                    }
                </div>
            </div>
        );
    }
}
