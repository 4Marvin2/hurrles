import React from "react";
import HomeEnable from '../imgs/tapbar/homeEnable.svg';
import HomeDisable from '../imgs/tapbar/homeDisable.svg';
import LikedEnable from '../imgs/tapbar/likedEnable.svg';
import LikedDisable from '../imgs/tapbar/likedDisable.svg';
import CartEnable from '../imgs/tapbar/cartEnable.svg';
import CartDisable from '../imgs/tapbar/cartDisable.svg';
import ProfileEnable from '../imgs/tapbar/profileEnable.svg';
import ProfileDisable from '../imgs/tapbar/profileDisable.svg';
import Exit from '../imgs/tapbar/exit.svg';
import '../css/TapBar.css';

export default class TapBar extends React.Component {
    tapBarClick(payload) {
        this.props.tapBarClick(payload);
    }

    render() {
        return (
            <nav className="tapbar">
                <button className="tapbar__icon" onClick={() => this.tapBarClick('home')}>
                    {this.props.mainPage === 'home' &&
                        <img src={HomeEnable} alt="HomeEnable" />
                    }
                    {this.props.mainPage !== 'home' &&
                        <img src={HomeDisable} alt="HomeDisable" />
                    }
                </button>
                <button className="tapbar__icon" onClick={() => this.tapBarClick('restorOpen')}>
                    {this.props.mainPage === 'restorOpen' &&
                        <img src={LikedEnable} alt="LikedEnable" />
                    }
                    {this.props.mainPage !== 'restorOpen' &&
                        <img src={LikedDisable} alt="LikedDisable" />
                    }
                </button>
                <button className="tapbar__icon" onClick={() => this.tapBarClick('cart')}>
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
                <button className="tapbar__exit-icon">
                    <img src={Exit} alt="Exit" />
                </button>
            </nav>            
        )
    }
}