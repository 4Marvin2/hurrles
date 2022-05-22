import React from 'react'
import Dish from './Dish'
import Button from '../Common/Button'
import '../../css/RestorOpen/RestorOpen.css'
import '../../css/RestorOpen/Menu.css'
import '../../css/RestorOpen/Dish.css'
import "../../css/App.css";

export default class Menu extends React.Component {
    updateButtonClick(payload) {
        this.props.updateButtonClick(payload)
    }

    render(){
        const dishesList = this.props.dishes.map((dish) =>
            <Dish key={dish.id} dish={dish} isAdmin={this.props.isAdmin} updateButtonClick={this.props.updateButtonClick} />
        );

        return (
            <div className="restor-open__menu menu">
                <div className="menu__title">Меню</div>
                { this.props.isAdmin &&
                    <Button text='Добавить' id={this.props.restorId} theme='dark' width={140} height={32} onClick={this.props.addButtonClick}/>
                }
                <div className="menu__dishes">
                    {dishesList}
                    <div className="menu__dish dish__hidden"></div>
                </div>
            </div>
        );
    }
}