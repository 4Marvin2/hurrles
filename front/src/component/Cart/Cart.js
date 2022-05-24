import React from "react";
import CartRestors from "./CartRestors";

import '../../css/Cart/Cart.css'
import { getOrders } from "../../requests/orders";
import { Bars } from  'react-loader-spinner'


export default class Cart extends React.Component {
    constructor(props) {
        super(props);
        this.state = {promiseInProgress: true, orders: []}
        getOrders()
        .then((data) => {
            console.log(data)
            this.setState(
                {
                    promiseInProgress: false,
                    orders: data,
                }
            )
        })
        .catch((error) => {
            console.log(error)
            this.setState(
                {
                    promiseInProgress: false,
                }
            )
        })

        this.updateCallback = this.updateCallback.bind(this);
    }

    updateCallback() {
        getOrders()
        .then((data) => {
            console.log(data)
            this.setState(
                {
                    promiseInProgress: false,
                    orders: data,
                }
            )
        })
        .catch((error) => {
            console.log(error)
            this.setState(
                {
                    promiseInProgress: false,
                }
            )
        })
    }

    render() {
        return (
            <div className='cart'>
                <div className="cart__title">
                    Ближайшие бронирования
                </div>
                { this.state.promiseInProgress &&
                    <Bars color="#FFFFFF" height={80} width={80} />
                }
                { !this.state.promiseInProgress && 
                    <CartRestors orders={this.state.orders} updateCallback={this.updateCallback}/>
                }
            </div>
        );
    }
}