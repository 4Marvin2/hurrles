import React from 'react'
import Canvas from '../Canvas';
import ReserveTitleBar from './ReserveTitleBar';

import '../../css/Reserve/Reserve.css'

export default class Reserve extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            places: [
                {
                    id: 1,
                    left: 250,
                    top: 250,
                    width: 25,
                    height:25,
                    nubmer: 1,
                    capacity: 1,
                    floor: 1,
                },
                {
                    id: 2,
                    left: 0,
                    top: 0,
                    width: 25,
                    height:25,
                    nubmer: 1,
                    capacity: 1,
                    floor: 2,
                }
            ],
            currentFloor: 1,
            isMouseDown: false,
            currentIndex: -1,
            offsetX: 0,
            offsetY: 0,
        }
        this.floorChange = this.floorChange.bind(this);
    }

    floorChange(payload) {
        this.setState({currentFloor: payload});
    }

    render(){
        return (
            <div className='reserve'>
                <ReserveTitleBar reserveClick={this.props.reserveClick} />
                <div className='reserve__floor-time'>
                <select name="pets" id="pet-select" onChange={(e) => this.floorChange(e.target.value)} value={this.state.currentFloor}>
                    <option value={1}>1</option>
                    <option value={2}>2</option>
                </select>
                <select name="pets" id="pet-select">
                    <option value={'10:00'}>10:00</option>
                    <option value={'11:00'}>11:00</option>
                </select>
                </div>
                <div className='reserve__canvas'>
                    <Canvas currentFloor={this.state.currentFloor} places={this.state.places} />
                </div>
            </div>
        );
    }
}
