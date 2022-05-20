import React from "react";
import '../../css/RestorOpen/RestorOpen.css'
import '../../css/RestorOpen/Menu.css'
import '../../css/RestorOpen/Info.css'
import '../../css/RestorOpen/Dish.css'
import "../../css/App.css";

export default class Info extends React.Component {
  render() {
    return (
        <div className="restor-open__info">
          <div className="restor-open__address-schedule">
            Номер ресторана:  {this.props.restorInfo.number}
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
            <div className="restor-open__tag1">
              {this.props.restorInfo.tag1}
            </div>
            {this.props.restorInfo.tag2 && 
              <div className="restor-open__tag2">
              {this.props.restorInfo.tag2}
            </div>}
          </div>
        </div>
    );
  }
}
