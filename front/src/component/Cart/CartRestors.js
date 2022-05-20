import React from 'react'
import CartRestor from './CartRestor';

import '../../css/Cart/Cart.css'

export default class CartRestors extends React.Component {
    render(){
        const cartRestors = this.props.cartRestors.map((cartRestor) =>
            <CartRestor key={cartRestor.id} />
        );

        return (
            <div className="cart">
                <div className="cart__title">
                    Ближайшие бронирования
                </div>
                <div className="cart__cart-restors">
                    {cartRestors}
                </div>
            </div>
        );
    }
}