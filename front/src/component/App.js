import React from "react";
import Home from './Restors/Home'
import RestorOpen from './RestorOpen/RestorOpen'
import TapBar from './TapBar'
import '../css/RestorOpen/RestorOpen.css'
import '../css/RestorOpen/Menu.css'
import '../css/RestorOpen/Info.css'
import '../css/RestorOpen/Dish.css'
import "../css/App.css";

export default class App extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            mainPage: 'home'
        };

        this.tapBarClick = this.tapBarClick.bind(this);
    }
    
    tapBarClick(payload) {
        this.setState({mainPage: payload});
    }

    render() {
        return (
            <div className="app">
                {this.state.mainPage === 'home' &&
                <Home />}
                {this.state.mainPage === 'restorOpen' &&
                <RestorOpen />}
                <TapBar reserveFlag={true} mainPage={this.state.mainPage} tapBarClick={this.tapBarClick} />
            </div>
        )
    }
}