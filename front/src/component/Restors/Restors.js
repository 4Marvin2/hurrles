import React from 'react'
import Restor from './Restor'
import '../../css/Restors.css'

export default class Restors extends React.Component {
    render(){
        const restorsList = this.props.restors.map((restor) =>
            <Restor key={restor.id} restor={restor} />
        );
        
        return (
            <div className='home__restors restors'>
                {restorsList}
            </div>
        );
    }
}