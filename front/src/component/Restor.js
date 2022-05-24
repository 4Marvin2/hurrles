import React from "react";
import Star from "../imgs/restor/star.svg";
import Metro from "../imgs/restor/metro.svg";
// import Metro1 from "../imgs/restor/metro_1.svg";
// import Metro2 from "../imgs/restor/metro_2.svg";
// import Metro3 from "../imgs/restor/metro_3.svg";
// import Metro4 from "../imgs/restor/metro_4.svg";
// import Metro5 from "../imgs/restor/metro_5.svg";
// import Metro6 from "../imgs/restor/metro_6.svg";
// import Metro7 from "../imgs/restor/metro_7.svg";
// import Metro8 from "../imgs/restor/metro_8.svg";
// import Metro9 from "../imgs/restor/metro_9.svg";
// import Metro10 from "../imgs/restor/metro_10.svg";
// import Metro11 from "../imgs/restor/metro_11.svg";
// import Metro11k from "../imgs/restor/metro_12.svg";
// import Metro12 from "../imgs/restor/metro_13.svg";
import "../css/Restor.css";
import { metroMap } from "../utils/metroMap";

export default class Restor extends React.Component {
  restorClick(payload) {
    this.props.restorClick(payload);
  }

  render() {
    return (
      <button className="restor" onClick={() => this.restorClick(this.props.index)}>
        {/* <div className="restor__img-column"> */}
          {/* <img src={require("../imgs/restors/nagoya.webp")} alt="restor" /> */}
        {/* </div> */}
        <div className="restor__info-column">
          <div className="restor__title">{this.props.restor.title}</div>
          {/* <div className="restor__rating"> */}
            {/* <img src={Star} alt="star" /> */}
            {/* <p>{this.props.restor.rating}</p> */}
          {/* </div> */}
          <div className="restor__metro">
            <img src={Metro} alt="metro" />
            { metroMap[this.props.restor.metro.toLowerCase()] &&
              <img src={require(`../imgs/metro/${metroMap[this.props.restor.metro.toLowerCase()]}`)} alt="metro-color" />
            }
            <div>{this.props.restor.metro}</div>
          </div>
          <div className="restor__tag">
              <span className="restor__tag-text">{this.props.restor.restorInfo.tag1}</span>
          </div>
        </div>
      </button>
    );
  }
}
