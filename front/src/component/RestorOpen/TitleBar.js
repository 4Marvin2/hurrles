import React from "react";
import LikeDisable from "../../imgs/titlebar/like_disable.svg"
import LikeEnable from "../../imgs/titlebar/like_enable.svg"
import "../../css/RestorOpen/TitleBar.css"

export default class TitleBar extends React.Component {
  render() {
    return (
        <nav className="restor-open__restor-titlebar restor-titlebar">
          <div className="restor-titlebar__back">
          </div>
          <div className="restor-titlebar__back"></div>
          <div className="restor-titlebar__img">
            <img src={require("../../imgs/restors/nagoya.webp")} alt="restor" />
          </div>
          <div className="restor-titlebar__like"></div>
          <div className="restor-titlebar__like">
            <button>
              {this.props.isFavorite &&
                <img src={LikeEnable} alt="like" />
              }
              {!this.props.isFavorite &&
                <img src={LikeDisable} alt="like" />
              }
            </button>
          </div>
        </nav>
    );
  }
}