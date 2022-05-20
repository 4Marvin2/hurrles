import React from "react";
import DishLogo from "../../imgs/restors/nagoya/set_curry_rise.svg";
import '../../css/RestorOpen/RestorOpen.css'
import '../../css/RestorOpen/Menu.css'
import '../../css/RestorOpen/Info.css'
import '../../css/RestorOpen/Dish.css'
import "../../css/App.css";

export default class Dish extends React.Component {
  render() {
    return (
      <div className="menu__dish dish">
        <div className="dish__img-column">
            <img src={DishLogo} alt="set_curry_rise" />
        </div>
        <div className="dish__info-column">
            <div className="dish__title">{this.props.dish.title}</div>
            <div className="dish__desc">{this.props.dish.desc}</div>
            <div className="dish__price">{this.props.dish.price} р.</div>
        </div>
    </div>
    );
  }
}
