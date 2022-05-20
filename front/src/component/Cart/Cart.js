import React from "react";
import CartRestors from "./CartRestors";

import '../../css/Cart/Cart.css'

export default class Cart extends React.Component {
    render() {
        const cartRestors = [
            {id: 1},
            {id: 2},
            {id: 3},
        ]
        return (
            <div className="cart">
                <CartRestors cartRestors={cartRestors} />
            </div>
        );
    }
}