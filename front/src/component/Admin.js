import React from 'react'
import Restors from './Restors'
import SearchBar from  './SearchBar'
import AddRestorForm from './AdminCommon/AddRestorForm'
import AdminRestorOpen from './AdminCommon/AdminRestorOpen'
import AdminPlaces from './AdminCommon/AdminPlaces' 
import '../css/Admin.css'

import { getRestors, getRestorMenu } from '../requests/restors';
import Button from './Common/Button'

import { dateToTimeString } from '../utils/date'

import Exit from '../imgs/tapbar/exit.svg';

var BreakException = {};

export default class Home extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            addFormActive: false,               // add restaurant
            updateRestorFormActive: false,      // update restaurant
            restorFormActive: false,            // open restaurant with pissible to add dish
            addDishFormActive: false,           // add dish
            currentIndex: 0,
            isFavorite: false,
            restors:[],
            map: false,
        }

        this.restorClick = this.restorClick.bind(this);

        getRestors().then((data) => {
            if (!data) {
                return
            }
            const restors = [];
            data.forEach(e => {
                getRestorMenu(e.id).then((data) => {
                    const dishes = [];
                    if (data !== null) {
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
                            title: e.title,
                            number: e.number,
                            address: e.address,
                            workTime: `Понедельник-суббота ${fullWorkTimeStr}`,
                            desc: e.description,
                            tag1: e.kitchen
                        },
                        dishes: dishes,
                        isFavorite: false,
                        floors: e.floors,
                    };
                    restors.push(restorElement);
                    this.setState({
                        currentIndex: 0,
                        restors: restors,
                    });
                });
            });
        });

        this.updateRestors = this.updateRestors.bind(this);
        this.openUpdateRestorClick = this.openUpdateRestorClick.bind(this);
        this.closeAllForms = this.closeAllForms.bind(this);
        this.mapClick = this.mapClick.bind(this);
        this.searchCallback = this.searchCallback.bind(this);
    }

    updateRestors(payload) {
        getRestors().then((data) => {
            if (!data) {
                return
            }
            const restors = [];
            data.forEach(e => {
                getRestorMenu(e.id).then((data) => {
                    const dishes = [];
                    if (data !== null) {
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
                            title: e.title,
                            number: e.number,
                            address: e.address,
                            workTime: `Понедельник-суббота ${fullWorkTimeStr}`,
                            desc: e.description,
                            tag1: e.kitchen,
                            openTime: e.openTime,
                            closeTime: e.closeTime,
                        },
                        dishes: dishes,
                        isFavorite: false,
                        floors: e.floors,
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

    openUpdateRestorClick(payload) {
        this.setState({updateRestorFormActive: true, addFormActive: false, restorFormActive: false})
    }

    restorClick(payload) {
        this.setState({
            currentIndex: payload,
            addFormActive: false,
            restorFormActive: true,
            updateRestorFormActive: false,
        });
    }

    addRestorButtonClick = () => {
        this.setState({addFormActive: !this.state.addFormActive, restorFormActive: false, updateRestorFormActive: false})
    }

    closeUpdateRestorForm(payload) {
        // ???
    }

    openAddDishForm = () => {
        this.setState({addFormActive: false, restorFormActive: false, updateRestorFormActive: false})
    }

    closeAllForms = () => {
        this.setState({
            addFormActive: false,
            restorFormActive: false,
            updateRestorFormActive: false,
        })
    }

    logoutClick() {
        this.props.logoutClick()
    }

    mapClick(payload) {
        this.setState({map: payload});
    }

    searchCallback(payload) {
        this.setState({
            currentIndex: 0,
            restors: payload.restors,
        })
    }

    render(){
        return (
            <div className='admin'>
                <div className='admin__left'>
                    <SearchBar searchCallback={this.searchCallback}/>
                    <div className='admin__add-button'>
                        <Button text='Добавить новый ресторан' theme='light' width={250} height={32} onClick={this.addRestorButtonClick}/>
                    </div>
                    <Restors restors={this.state.restors} restorClick={this.restorClick} />
                </div>
                <div className={ (!this.state.addFormActive && !this.state.restorFormActive && !this.state.updateRestorFormActive) ? 'admin__right admin__right_hidden' : 'admin__right' }>
                    { this.state.addFormActive && !this.state.restorFormActive && !this.state.updateRestorFormActive &&
                        <AddRestorForm
                            updateRestorsCallback={this.updateRestors}
                            buttonText='Добавить'
                            closeAllForms={this.closeAllForms}
                        />
                    }
                    { !this.state.addFormActive && !this.state.restorFormActive && this.state.updateRestorFormActive &&
                        <AddRestorForm
                            isUpdate={true}
                            updateRestorsCallback={this.updateRestors}
                            buttonText='Изменить'
                            restor={
                                {
                                    id: this.state.restors[this.state.currentIndex].id,
                                    title: this.state.restors[this.state.currentIndex].title,
                                    desc: this.state.restors[this.state.currentIndex].restorInfo.desc,
                                    kitchen: this.state.restors[this.state.currentIndex].restorInfo.tag1,
                                    address: this.state.restors[this.state.currentIndex].restorInfo.address,
                                    metro: this.state.restors[this.state.currentIndex].metro,
                                    phoneNumber: this.state.restors[this.state.currentIndex].restorInfo.number,
                                    openTime: this.state.restors[this.state.currentIndex].restorInfo.openTime,
                                    closeTime: this.state.restors[this.state.currentIndex].restorInfo.closeTime,
                                }
                            }
                            closeAllForms={this.closeAllForms}
                        />
                    }
                    { this.state.restorFormActive && !this.state.addFormActive && !this.state.updateRestorFormActive &&
                        <div className='admin__form'>
                            <AdminRestorOpen
                                admin={true}
                                dishes={this.state.restors[this.state.currentIndex].dishes}
                                restorId={this.state.restors[this.state.currentIndex].id}
                                restorInfo={this.state.restors[this.state.currentIndex].restorInfo}
                                updateRestorsCallback={this.updateRestors}
                                openUpdateRestorClick={this.openUpdateRestorClick}
                                closeAllForms={this.closeAllForms}
                                mapClick={this.mapClick}
                            />
                        </div>
                    }
                    <button className="admin__exit-button">
                        <img src={Exit} alt="Exit" onClick={() => {this.logoutClick()}} />
                    </button>
                    {this.state.map &&
                        <div className='admin__map'>
                            <AdminPlaces 
                                mapClick={this.mapClick}
                                id={this.state.restors[this.state.currentIndex].id}
                                userId={this.props.userId}
                                floors={this.state.restors[this.state.currentIndex].floors} />
                        </div>
                    } 
                </div>
            </div>
        );
    }
}
