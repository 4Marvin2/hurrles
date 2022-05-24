import React from "react";
import Home from './Home'
import Favor from './Favor'
import Cart from "./Cart/Cart";
import TapBar from './TapBar'
import AuthForm from "./AuthForm/AuthForm";
import Profile from "./Profile";
import '../css/App.css'

import { getUserRequest , logoutRequest} from '../requests/user';

import { Bars } from  'react-loader-spinner'
import Admin from "./Admin";
import RestorAdmin from "./RestorAdmin";

export default class App extends React.Component {
    constructor(props) {
        super(props);
        this.state = {promiseInProgress: true}
        getUserRequest()
        .then((data) => {
            this.setState(
                {
                    promiseInProgress: false,
                    loggedIn: true,
                    mainPage: 'home',
                    user: data,
                }
            )
        })
        .catch((error) => {
            this.setState(
                {
                    promiseInProgress: false,
                    loggedIn: false,
                }
            )
        })

        this.tapBarClick = this.tapBarClick.bind(this);
        this.logInHandler = this.logInHandler.bind(this);
        this.logoutClick = this.logoutClick.bind(this);
    }

    tapBarClick(payload) {
        this.setState({mainPage: payload});
    }

    logInHandler(payload) {
        console.log(payload.user)
        this.setState({loggedIn: payload.loggedIn, mainPage: 'home', user: payload.user, user: payload.user})
    }

    logoutClick() {
        logoutRequest()
        .then(() => {
            this.setState({loggedIn: false})
        })
        .catch(() => {
            // this.setState({loggedIn: false})
        })
    }

    render() {
        return (
            <div className={this.state.promiseInProgress ? "app__body app__body_center" : "app__body"} >
                { this.state.promiseInProgress &&
                    <Bars color="#FFFFFF" height={80} width={80} />
                }
                { !this.state.promiseInProgress && !this.state.loggedIn && 
                    <div className="app__body">
                        <AuthForm authType='login' logInHandler={this.logInHandler}/>
                    </div>
                }
                { !this.state.promiseInProgress && this.state.loggedIn && !this.state.user.isAdmin && !this.state.user.isRestaurant &&
                    <div className="app">
                        <div className="app__tapbar">
                            <TapBar reserveFlag={true} mainPage={this.state.mainPage} tapBarClick={this.tapBarClick} logoutClick={this.logoutClick}/>
                        </div>
                        <div className="app__body">
                            {this.state.mainPage === 'home' &&
                            <Home userId={this.state.user.id} />}
                            {this.state.mainPage === 'restorOpen' &&
                            <Favor userId={this.state.user.id} />}
                            {this.state.mainPage === 'cart' &&
                            <Cart />}
                            {this.state.mainPage === 'profile' &&
                            <Profile name={this.state.user.fullName} email={this.state.user.email} phoneNumber={this.state.number}/>}
                        </div>
                    </div>
                }
                { !this.state.promiseInProgress && this.state.loggedIn && this.state.user.isAdmin &&
                    <Admin logoutClick={this.logoutClick}/>
                }
                { !this.state.promiseInProgress && this.state.loggedIn && this.state.user.isRestaurant &&
                    <RestorAdmin />
                }
            </div>
        )
    }
}