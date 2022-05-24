import React from "react";
import CancelButton from "../Common/CancelButton";
import Button from "../Common/Button";

import Metro from '../../imgs/restor/metro.svg'
import MetroBlue from '../../imgs/restor/metro_blue.svg'
import Munis from '../../imgs/cart/minus.svg'
import Plus from '../../imgs/cart/plus.svg'
import '../../css/Cart/CartRestor.css'
import { addDish, deleteDish, deleteOrder } from "../../requests/orders";

export default class CartRestor extends React.Component {
    updateCallback() {
        this.props.updateCallback();
    }

    addDishButtonClick(payload) {
        console.log(payload)
        addDish(payload.orderId, payload.dishId)
        .then((data) => {
            this.updateCallback()
        })
        .catch((error) => {
            console.log(error)
        })
    }

    deleteDishButtonClick(payload) {
        deleteDish(payload.orderId, payload.dishId)
        .then((data) => {
            this.updateCallback()
        })
        .catch((error) => {
            console.log(error)
        })
    }

    closeOrderButtonClick(payload) {
        deleteOrder(payload.orderId)
        .then((data) => {
            this.updateCallback()
        })
        .catch((error) => {
            console.log(error)
        })
    }

    render() {
        return (
            <div className="cart-restor">
                <div className="cart-restor__content">
                    <div className='cart-restor__title'>{this.props.data.restaurantTitle}</div>
                    <div className='cart-restor__metro'>
                        <div className='cart-restor__row-comp'><img src={Metro} alt="metro" /></div>
                        <div className='cart-restor__row-comp'><img src={MetroBlue} alt="metro_blue" /></div>
                        <div className='cart-restor__row-comp'>{this.props.data.restaurantMetro}</div>
                    </div>
                    <div className='cart-restor__column-comp'>{this.props.data.restaurantAddress}</div>
                    <div className='cart-restor__column-comp'>{this.props.data.startTime}</div>
                    <div className='cart-restor__column-comp'>Столик №{this.props.data.placeNumber} на {this.props.data.placeCapacity}</div>
                    { this.props.data.dishesIds &&
                        <ul className='cart-restor__menu'>
                            {
                                this.props.data.dishesIds.map((dishId, i) => {
                                    return (
                                        <li key={i}>
                                            <div className='cart-restor__dish'>
                                                <div className='cart-restor__dish-title'>{this.props.data.dishesTitles[i]}</div>
                                                <button className='cart-restor__row-comp'><img src={Munis} alt="delete" onClick={() => {this.deleteDishButtonClick({orderId: this.props.data.id, dishId: dishId})}} /></button>
                                                <div className='cart-restor__row-comp'>{this.props.data.dishesCounts[i]}</div>
                                                <button className='cart-restor__row-comp'><img src={Plus} alt="add" onClick={() => {this.addDishButtonClick({orderId: this.props.data.id, dishId: dishId})}} /></button>
                                                <div className='cart-restor__row-comp'>{this.props.data.dishesPrices[i]} р.</div>
                                            </div>
                                        </li>
                                    );
                                })
                            }
                            <li className='cart-restor__total'>
                                <div className='cart-restor__dish'>
                                    <div className='cart-restor__total-title'>Итого</div>
                                    <div className='cart-restor__row-comp'>{this.props.data.cost} р.</div>
                                </div>
                            </li>
                        </ul>
                    }
                    <div className='cart-restor__cancel-button'>
                        {/* <span className={ this.state.formErrorIsActive ? 'auth-form__form-error auth-form__form-error_active' : 'auth-form__form-error' }>{this.state.formError}</span> */}
                        <Button text='Отменить' theme='cancel' width={100} height={30} onClick={() => {this.closeOrderButtonClick({orderId: this.props.data.id})}}/>
                    </div>
                </div>
            </div>
        );
    }
}