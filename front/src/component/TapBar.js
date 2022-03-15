import React from "react";
import HomeEnable from '../imgs/tapbar/homeEnable.svg';
import LikedDisable from '../imgs/tapbar/likedDisable.svg';
import CartDisable from '../imgs/tapbar/cartDisable.svg';
import ProfileDisable from '../imgs/tapbar/profileDisable.svg';
import '../css/TapBar.css';

export default class TapBar extends React.Component {
    render() {
        return (
            <div className="app__tapbar tapbar">
                <button className="tapbar__icon">
                    <img src={HomeEnable} alt="home" />
                </button>
                <button className="tapbar__icon">
                    <img src={LikedDisable} alt="liked" />
                </button>
                <button className="tapbar__icon">
                    <img src={CartDisable} alt="cart" />
                </button>
                <button className="tapbar__icon">
                    <img src={ProfileDisable} alt="profile" />
                </button>
            </div>
            // <div className="tapBar">
            //     <button className="tapBar-icons">
            //         <img src={HomeEnable} alt="test" />
            //     </button>
            //     <button className="tapBar-icons">
            //         <img src={LikedDisable} alt="test" />
            //     </button>
            //     <button className="tapBar-icons">
            //         <img src={CartDisable} alt="test" />
            //     </button>
            //     <button className="tapBar-icons">
            //         <img src={ProfileDisable} alt="test" />
            //     </button>
            // </div>
        )
    }
}