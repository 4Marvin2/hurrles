import React from "react";

import Back from '../../imgs/titlebar/back.svg'

import "../../css/RestorOpen/TitleBar.css"

export default class ReserveTitleBar extends React.Component {
  reserveClick() {
    this.props.reserveClick(false);
  }

  render() {
    return (
        <nav className="restor-open__restor-titlebar restor-titlebar">
          <div className="restor-titlebar__back">
            <button onClick={() => this.reserveClick()}>
                <img src={Back} alt="like" />
            </button>
          </div>
          <div className="restor-titlebar__back"></div>
          <div className="restor-titlebar__img">
            <img src={require("../../imgs/restors/nagoya.webp")} alt="restor" />
          </div>
          <div className="restor-titlebar__like"></div>
          <div className="restor-titlebar__like"></div>
        </nav>
    );
  }
}