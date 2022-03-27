import React from "react";
import Star from "../../imgs/restor/star.svg";
import Metro from "../../imgs/restor/metro.svg";
import MetroBlue from "../../imgs/restor/metro_blue.svg";
import "../../css/Restors/Restor.css";

export default class Restor extends React.Component {
  render() {
    return (
      <div className="restors__restor restor">
        <div className="restor__img-column">
          <img src={require("../../imgs/restors/nagoya.webp")} alt="restor" />
        </div>
        <div className="restor__info-column">
          <div className="restor__title">{this.props.restor.title}</div>
          <div className="restor__rating">
            <img src={Star} alt="star" />
            <p>{this.props.restor.rating}</p>
          </div>
          <div className="restor__metro">
            <img src={Metro} alt="metro" />
            <img src={MetroBlue} alt="mtero-color" />
            <div>{this.props.restor.metro}</div>
          </div>
        </div>
      </div>
    );
  }
}
