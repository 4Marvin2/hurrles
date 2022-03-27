import React from "react";
import HomeEnable from '../imgs/tapbar/homeEnable.svg';
import HomeDisable from '../imgs/tapbar/homeDisable.svg';
import LikedEnable from '../imgs/tapbar/likedEnable.svg';
import LikedDisable from '../imgs/tapbar/likedDisable.svg';
import CartEnable from '../imgs/tapbar/cartEnable.svg';
import CartDisable from '../imgs/tapbar/cartDisable.svg';
import ProfileEnable from '../imgs/tapbar/profileEnable.svg';
import ProfileDisable from '../imgs/tapbar/profileDisable.svg';
import '../css/TapBar.css';
import '../css/App.css';

export default class TapBar extends React.Component {
    handleOnClick(payload) {
        this.props.tapBarClick(payload);
    }

    render() {
        return (
            <footer className="tapbar-reserve">
                {this.props.reserveFlag && 
                <button className="tapbar-reserve__reserve">
                Забронировать
                </button>}
                <div className="tapbar-reserve__tapbar tapbar">
                <button className="tapbar__icon" onClick={() => this.handleOnClick('home')}>
                    {this.props.mainPage === 'home' &&
                        <img src={HomeEnable} alt="HomeEnable" />
                    }
                    {this.props.mainPage !== 'home' &&
                        <img src={HomeDisable} alt="HomeDisable" />
                    }
                </button>
                <button className="tapbar__icon" onClick={() => this.handleOnClick('restorOpen')}>
                    {this.props.mainPage === 'restorOpen' &&
                        <img src={LikedEnable} alt="LikedEnable" />
                    }
                    {this.props.mainPage !== 'restorOpen' &&
                        <img src={LikedDisable} alt="LikedDisable" />
                    }
                </button>
                <button className="tapbar__icon">
                    {this.props.mainPage === 'cart' &&
                        <img src={CartEnable} alt="CartEnable" />
                    }
                    {this.props.mainPage !== 'cart' &&
                        <img src={CartDisable} alt="CartDisable" />
                    }
                </button>
                <button className="tapbar__icon">
                    {this.props.mainPage === 'profile' &&
                        <img src={ProfileEnable} alt="ProfileEnable" />
                    }
                    {this.props.mainPage !== 'profile' &&
                        <img src={ProfileDisable} alt="ProfileDisable" />
                    }
                </button>
                </div>
            </footer>            
        )
    }
}