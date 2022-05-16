import React from 'react'
import Menu from './Menu'
import TitleBar from '../RestorOpen/TitleBar'
import Info from '../RestorOpen/Info'
import '../../css/RestorOpen/RestorOpen.css'
import '../../css/RestorOpen/Menu.css'
import '../../css/RestorOpen/Info.css'
import '../../css/RestorOpen/Dish.css'

export default class Home extends React.Component {
    render(){
        return (
            <div className="app_restor-open restor-open">
                <TitleBar />
                <div className="restor-open__main">
                    <Info  restorInfo={this.props.restorInfo}/>
                    <Menu dishes={this.props.dishes} />
                </div>
            </div>
        );
    }
}
