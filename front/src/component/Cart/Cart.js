import React from "react";
import CartRestors from "./CartRestors";

import '../../css/Cart/Cart.css'
import { getOrders } from "../../requests/orders";

export default class Cart extends React.Component {
    constructor(props) {
        super(props);
        this.state = {promiseInProgress: true}
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
        const cartRestors = [
            {id: 1},
            {id: 2},
            {id: 3},
        ]
        return (
            <div className='cart'>
                <div className="cart__title">
                    Ближайшие бронирования
                </div>
                <CartRestors cartRestors={cartRestors} />
            </div>
        );
    }
}