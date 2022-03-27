import React from "react";
import Star from "../../imgs/restor/star.svg";
import Review from "../../imgs/restor/review.svg";
import '../../css/RestorOpen/RestorOpen.css'
import '../../css/RestorOpen/Menu.css'
import '../../css/RestorOpen/Info.css'
import '../../css/RestorOpen/Dish.css'
import "../../css/App.css";

export default class Info extends React.Component {
  render() {
    return (
        <div className="restor-open__info">
          <div className="restor-open__rating-reviews">
            <img src={Star} alt="star" />
            {this.props.restorInfo.rating}
            <img src={Review} alt="review" />
            Отзывы
          </div>
          <div className="restor-open__address-schedule">
            <p>{this.props.restorInfo.address}</p>
            <p>{this.props.restorInfo.workTime}</p>
          </div>
          <div className="restor-open__description">
            <div className="restor-open__title">
              Описание
            </div>
            <div className="restor-open__subtitle">
              {this.props.restorInfo.desc}
            </div>
          </div>
          <div className="restor-open__tags">
            <div className="restor-open__tag">
              {this.props.restorInfo.tag1}
            </div>
            {this.props.restorInfo.tag2 && 
              <div className="restor-open__tag" style="background: rgba(251, 177, 60, 0.7);">
              {this.props.restorInfo.tag2}
            </div>}
          </div>
        </div>
    );
  }
}
