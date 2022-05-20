import React from 'react'
import Restor from './Restor'
import '../css/Restors.css'

export default class Restors extends React.Component {
    render(){
        const restorsList = this.props.restors.map((restor, index) =>
            <Restor key={restor.id} restor={restor} restorClick={this.props.restorClick} index={index} />
        );
        
        return (
            <div className='home__restors restors'>
                {restorsList}
            </div>
        );
    }
}