import React from 'react'

import '../css/ModalMenu.css'
import { getRestorMenu } from '../requests/restors';
import Dish from './RestorOpen/Dish';
import '../css/RestorOpen/Dish.css'
import { addDish, deleteDish } from '../requests/orders';
import Exit from '../imgs/cart/exit.svg'


export default class ModalMenu extends React.Component {
    constructor(props) {
        super(props);
        this.state = {promiseInProgress: true, dishes: []}
        getRestorMenu(this.props.restorId)
        .then((data) => {
            const dishes = [];
            console.log(data)
            if (data) {
                data.forEach(e => {
                    const dishElement = {
                        id: e.id,
                        title: e.title,
                        desc: e.description,
                        price: e.price
                    };
                    dishes.push(dishElement);
                });
            }
            this.setState(
                {
                    promiseInProgress: false,
                    dishes: dishes,
                }
            )
        })
        .catch((error) => {
            console.log(error)
            this.setState(
                {
                    promiseInProgress: false,
                    dishes: [],
                }
            )
        })      
    }

    addDishButtonClick(payload) {
        addDish(payload.orderId, payload.dishId)
        .then((data) => {
            console.log(data)
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

    menuCloseClick() {
        this.props.menuCloseClick();
    }

    render(){
        let dishesList = []
        console.log('rgerferfwf', this.state.dishes, this.props.order)
        if (this.state.dishes !== undefined || this.state.dishes.length !== 0) {
            dishesList = this.state.dishes.map((dish) =>
                <Dish key={dish.id} dish={dish} order={this.props.order} addDishButtonClick={this.addDishButtonClick} deleteDishButtonClick={this.deleteDishButtonClick} updateCallback={this.props.updateCallback}/>
            );
        }

        return (
            <div className="modal-menu">
                <div className="modal-menu__menu">
                    <button className='modal-menu__exit-button' onClick={() => {this.props.menuCloseClick()}}>
                        <img src={Exit} alt="exit" />
                    </button>
                    <div className="modal-menu__title">Меню</div>
                    {this.state.dishes !== undefined &&
                        <div className="modal-menu__dishes">
                            {dishesList}
                            <div className="menu__dish dish__hidden"></div>
                        </div>
                    }
                </div>
            </div>
        );
    }
}