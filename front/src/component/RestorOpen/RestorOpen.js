import React from 'react'
import Menu from './Menu'
import ReserveButton from '../Common/ReserveButton'
import TitleBar from '../RestorOpen/TitleBar'
import Info from '../RestorOpen/Info'
import '../../css/RestorOpen/RestorOpen.css'
import '../../css/RestorOpen/Menu.css'
import '../../css/RestorOpen/Dish.css'


export default class RestorOpen extends React.Component {
    render(){
        return (
            <div className="restor-open">
                <TitleBar isFavorite={this.props.isFavorite} likeClick={this.props.likeClick} />
                <div className="restor-open__main">
                    <Info  restorInfo={this.props.restorInfo}/>
                    <Menu dishes={this.props.dishes} />
                </div>
                <div className="restor-open__reverse-button">
                    <ReserveButton text={'Забронировать'} />
                </div>
            </div>
        );
    }
}
