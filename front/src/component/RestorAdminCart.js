import React from "react";
import CartRestors from "./Cart/CartRestors";

import '../css/Cart/Cart.css'
import { getRestaurantOrders } from "../requests/orders";
import { Bars } from  'react-loader-spinner'
import ModalMenu from "./ModalMenu";


export default class RestorAdminCart extends React.Component {
    constructor(props) {
        super(props);
        this.state = {promiseInProgress: true, orders: [], menuOpened: false}
        getRestaurantOrders()
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
        this.menuOpenClick = this.menuOpenClick.bind(this);
        this.menuCloseClick = this.menuCloseClick.bind(this);
    }

    updateCallback() {
        getRestaurantOrders()
        .then((data) => {
            console.log(data)
            const orders = []
            let curOrder
            if (data) {
                data.map((order) => {
                    orders.push(order)
                    if (this.state.curOrder && order.id === this.state.curOrder.id) {
                        console.log(100)
                        curOrder = order
                    }
                })
            }
            this.setState(
                {
                    promiseInProgress: false,
                    orders: orders,
                    curOrder: curOrder
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

    menuOpenClick(payload) {
        let curOrder
        this.state.orders.map((order) => {
            if (order.id === payload.orderId) {
                curOrder = order
            }
        })
        this.setState({menuOpened: true, restorMenuId: payload.restorId, orderIdOpenedMenu: payload.orderId, curOrder: curOrder})
    }

    menuCloseClick() {
        this.setState({menuOpened: false})
    }

    render() {
        return (
            <div className='cart'>
                { this.state.menuOpened &&
                    <ModalMenu restorId={this.state.restorMenuId} updateCallback={this.updateCallback} order={this.state.curOrder} menuCloseClick={this.menuCloseClick}/>
                }
                <div className="cart__title">
                    Ближайшие бронирования
                </div>
                { this.state.promiseInProgress &&
                    <Bars color="#FFFFFF" height={80} width={80} />
                }
                { !this.state.promiseInProgress && 
                    <CartRestors isRestor={true} orders={this.state.orders} updateCallback={this.updateCallback} menuOpenClick={this.menuOpenClick}/>
                }
            </div>
        );
    }
}