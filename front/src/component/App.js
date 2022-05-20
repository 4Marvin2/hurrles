import React from "react";
import Home from './Home'
import Favor from './Favor'
import Cart from "./Cart/Cart";
import TapBar from './TapBar'
import '../css/App.css'

export default class App extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            mainPage: 'home',
        };

        this.tapBarClick = this.tapBarClick.bind(this);
    }

    tapBarClick(payload) {
        this.setState({mainPage: payload});
    }

    render() {
        return (
            <div className="app">
                <div className="app__tapbar">
                    <TapBar reserveFlag={true} mainPage={this.state.mainPage} tapBarClick={this.tapBarClick} />
                </div>
                <div className="app__body">
                    {this.state.mainPage === 'home' &&
                    <Home />}
                    {this.state.mainPage === 'restorOpen' &&
                    <Favor />}
                    {this.state.mainPage === 'cart' &&
                    <Cart />}
                </div>
            </div>
        )
    }
}