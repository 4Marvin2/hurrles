import React from "react";
import CancelButton from "../Common/CancelButton";

import Metro from '../../imgs/restor/metro.svg'
import MetroBlue from '../../imgs/restor/metro_blue.svg'
import Munis from '../../imgs/cart/minus.svg'
import Plus from '../../imgs/cart/plus.svg'
import '../../css/Cart/CartRestor.css'

export default class CartRestor extends React.Component {
    render() {
        return (
            <div className="cart-restor">
                <div className="cart-restor__content">
                    <div className='cart-restor__title'>NAGOYA</div>
                    <div className='cart-restor__metro'>
                        <div className='cart-restor__row-comp'><img src={Metro} alt="metro" /></div>
                        <div className='cart-restor__row-comp'><img src={MetroBlue} alt="metro_blue" /></div>
                        <div className='cart-restor__row-comp'>Бауманская</div>
                    </div>
                    <div className='cart-restor__column-comp'>Старокирочный переулок 16/2 стр.1</div>
                    <div className='cart-restor__column-comp'>7 февраля 16:30</div>
                    <div className='cart-restor__column-comp'>Столик №5 на двоих</div>
                    <ul className='cart-restor__menu'>
                        <li>
                            <div className='cart-restor__dish'>
                                <div className='cart-restor__dish-title'>Сет Биф А-Ля Русс</div>
                                <button className='cart-restor__row-comp'><img src={Munis} alt="metro_blue" /></button>
                                <div className='cart-restor__row-comp'>2</div>
                                <button className='cart-restor__row-comp'><img src={Plus} alt="metro_blue" /></button>
                                <div className='cart-restor__row-comp'>660 р.</div>
                            </div>
                        </li>
                        <li>
                            <div className='cart-restor__dish'>
                                <div className='cart-restor__dish-title'>Кока-Кола</div>
                                <button className='cart-restor__row-comp'><img src={Munis} alt="metro_blue" /></button>
                                <div className='cart-restor__row-comp'>2</div>
                                <button className='cart-restor__row-comp'><img src={Plus} alt="metro_blue" /></button>
                                <div className='cart-restor__row-comp'>660 р.</div>
                            </div>
                        </li>
                        <li>
                            <div className='cart-restor__dish'>
                                <div className='cart-restor__dish-title'>Карточшка фри</div>
                                <button className='cart-restor__row-comp'><img src={Munis} alt="metro_blue" /></button>
                                <div className='cart-restor__row-comp'>2</div>
                                <button className='cart-restor__row-comp'><img src={Plus} alt="metro_blue" /></button>
                                <div className='cart-restor__row-comp'>660 р.</div>
                            </div>
                        </li>
                        <li>
                            <div className='cart-restor__dish'>
                                <div className='cart-restor__dish-title'>Боксмастер</div>
                                <button className='cart-restor__row-comp'><img src={Munis} alt="metro_blue" /></button>
                                <div className='cart-restor__row-comp'>2</div>
                                <button className='cart-restor__row-comp'><img src={Plus} alt="metro_blue" /></button>
                                <div className='cart-restor__row-comp'>660 р.</div>
                            </div>
                        </li>
                        <li className='cart-restor__total'>
                            <div className='cart-restor__dish'>
                                <div className='cart-restor__total-title'>Итого</div>
                                <div className='cart-restor__row-comp'>660 р.</div>
                            </div>
                        </li>
                    </ul>
                    <div className='cart-restor__cancel-button' >
                        <CancelButton text='Отменить' />
                    </div>
                </div>
            </div>
        );
    }
}