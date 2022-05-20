import React from 'react'
import Dish from './Dish'
import '../../css/RestorOpen/RestorOpen.css'
import '../../css/RestorOpen/Menu.css'
import '../../css/RestorOpen/Dish.css'
import "../../css/App.css";

export default class Menu extends React.Component {
    render(){
        const dishesList = this.props.dishes.map((dish) =>
            <Dish key={dish.id} dish={dish} />
        );

        return (
            <div className="restor-open__menu menu">
                <div className="menu__title">Меню</div>
                <div className="menu__dishes">
                    {dishesList}
                    <div className="menu__dish dish__hidden"></div>
                </div>
            </div>
        );
    }
}