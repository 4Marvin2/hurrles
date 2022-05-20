import React from "react";
import LikeDisable from "../../imgs/titlebar/like_disable.svg"
import LikeEnable from "../../imgs/titlebar/like_enable.svg"
import "../../css/RestorOpen/TitleBar.css"

export default class TitleBar extends React.Component {
  likeClick(payload) {
    this.props.likeClick(payload);
  }

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
            <button onClick={() => this.likeClick(!this.props.isFavorite)}>
                <img src={this.props.isFavorite ? LikeEnable : LikeDisable} alt="like" />
            </button>
          </div>
        </nav>
    );
  }
}