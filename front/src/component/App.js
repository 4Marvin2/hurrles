import React from "react";
import Home from './Home'
import TapBar from './TapBar'
import '../css/App.css';

export default class App extends React.Component {
    render() {  
        return (
            <div className="app">
                <Home />
                <TapBar />
            </div>
        )
    }
}