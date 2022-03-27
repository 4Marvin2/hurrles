import React from "react";
import Back from "../imgs/navbar/back.svg"
import LikeDisable from "../imgs/navbar/like_disable.svg"
import "../css/RestorNavBar.css"

export default class RestorNavBar extends React.Component {
  render() {
    return (
        <nav className="restor-open__restor-navbar restor-navbar">
          <div className="restor-navbar__back">
            <button>
              <img src={Back} alt="back" />
            </button>
          </div>
          <div className="restor-navbar__back"></div>
          <div className="restor-navbar__img">
            <img src={require("../imgs/restors/nagoya.webp")} alt="restor" />
          </div>
          <div className="restor-navbar__like"></div>
          <div className="restor-navbar__like">
            <button>
              <img src={LikeDisable} alt="like" />
            </button>
          </div>
        </nav>
    );
  }
}
