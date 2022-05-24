import React from 'react'
import CartRestor from './CartRestor';

import '../../css/Cart/Cart.css'
import { dateToDateString } from '../../utils/date';

export default class CartRestors extends React.Component {
    constructor(props) {
        super(props);
    }

    render(){
        let orders
        if (this.props.orders !== undefined) {
            orders = this.props.orders.map((order) => {
                order['startTimeString'] = dateToDateString(order.startTime)
                return (
                    <CartRestor isRestor={this.props.isRestor} key={order.id} data={order} updateCallback={this.props.updateCallback} menuOpenClick={this.props.menuOpenClick}/>
                );
            });
        }

        return (
            <div className="cart__cart-restors">
                { this.props.orders !== undefined &&
                    <div className='cart__cart-restors_scroll'>
                        {orders}
                    </div>
                }
            </div>
        );
    }
}