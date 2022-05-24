import React from "react";
import Button from "../Common/Button";
import DishLogo from "../../imgs/restors/nagoya/set_curry_rise.svg";
import Munis from '../../imgs/cart/minus.svg'
import Plus from '../../imgs/cart/plus.svg'
import '../../css/RestorOpen/RestorOpen.css'
import '../../css/RestorOpen/Menu.css'
import '../../css/RestorOpen/Dish.css'
import "../../css/App.css";
import '../../css/Cart/CartRestor.css'

export default class Dish extends React.Component {

  render() {
    let dishesNumber = 0
    if (this.props.order && this.props.order.dishesIds) {
      this.props.order.dishesIds.map((dishId, i) => {
        if (dishId === this.props.dish.id) {
          dishesNumber = this.props.order.dishesCounts[i]
        }
      })
    }
    let buttonTmpl
    if (this.props.order) {
      if (dishesNumber !== 0) {
        buttonTmpl = (
          <div className="dish__add-delete-buttons">
            <button className='cart-restor__row-comp'>
              <img src={Munis} alt="delete" onClick={() => { this.props.deleteDishButtonClick({ orderId: this.props.order.id, dishId: this.props.dish.id}); this.props.updateCallback(); }}/>
              </button>
            <div className='cart-restor__row-comp'>{dishesNumber}</div>
            <button className='cart-restor__row-comp'>
              <img src={Plus} alt="add" onClick={() => {this.props.addDishButtonClick({ orderId: this.props.order.id, dishId: this.props.dish.id}); this.props.updateCallback();}} />
            </button>
          </div>
        );
      }
    }
    if (this.props.order) {
      if (dishesNumber === 0) {
        buttonTmpl = (
          <div>
            <Button text='Добавить' theme='dark' width={82} height={24} onClick={() => { this.props.addDishButtonClick({ orderId: this.props.order.id, dishId: this.props.dish.id}); this.props.updateCallback(); }}/>
          </div>
        );
      }
      if (this.props.isAdmin) {
        buttonTmpl = (
          <div>
            <Button text='Изменить' theme='dark' width={82} height={24} onClick={() => { this.props.updateButtonClick(this.props.dish) }}/>
          </div>
        )
      }
    }


    return (
      <div className="menu__dish dish">
        {/* <div className="dish__img-column">
            <img src={DishLogo} alt="set_curry_rise" />
        </div> */}
        <div className="dish__info-column">
            <div className="dish__title">{this.props.dish.title}</div>
            <div className="dish__desc">{this.props.dish.desc}</div>
            <div className="dish__price">{this.props.dish.price} р.</div>
            <div className="dish__add-button">
              {buttonTmpl}
            </div>
        </div>
    </div>
    );
  }
}
