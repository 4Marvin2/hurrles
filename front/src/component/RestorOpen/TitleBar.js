import React from "react";
import Back from "../../imgs/titlebar/back.svg"
import LikeDisable from "../../imgs/titlebar/like_disable.svg"
import "../../css/RestorOpen/TitleBar.css"

export default class TitleBar extends React.Component {
  render() {
    return (
        <nav className="restor-open__restor-titlebar restor-titlebar">
          <div className="restor-titlebar__back">
            <button>
              <img src={Back} alt="back" />
            </button>
          </div>
          <div className="restor-titlebar__back"></div>
          <div className="restor-titlebar__img">
            <img src={require("../../imgs/restors/nagoya.webp")} alt="restor" />
          </div>
          <div className="restor-titlebar__like"></div>
          <div className="restor-titlebar__like">
            <button>
              <img src={LikeDisable} alt="like" />
            </button>
          </div>
        </nav>
    );
  }
}