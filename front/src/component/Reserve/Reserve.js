import React from 'react'
import CanvasClient from './CanvasClient';
import OrderButton from '../Common/OrderButton'
import ReserveTitleBar from './ReserveTitleBar';

import '../../css/Reserve/Reserve.css'

import { getPlaces } from '../../requests/restors';
import { createOrder } from '../../requests/orders';

var BreakException = {};

export default class Reserve extends React.Component {
    constructor(props) {
        super(props);
        const today = new Date();
        const month = today.getMonth()+1;
        const day = today.getDate();
        const date = today.getFullYear()+'-'+(this.formate(month))+'-'+this.formate(day);
        this.state = {
            places: [],
            currentFloor: 1,
            currentTime: 'T10:00:00.00Z',
            currentDate: date,
            isMouseDown: false,
            currentIndex: -1,
            currentId: 0,
            offsetX: 0,
            offsetY: 0,
        }

        this.floorChange = this.floorChange.bind(this);
        this.timeChange = this.timeChange.bind(this);
        this.dateChange = this.dateChange.bind(this);
        this.orderClick = this.orderClick.bind(this);
    }


    formate(num) {
        return num > 9 ? "" + num: "0" + num;
    }

    componentDidMount(){
        const today = new Date();
        const month = today.getMonth()+1;
        const day = today.getDate();
        const date = today.getFullYear()+'-'+(this.formate(month))+'-'+this.formate(day);
        const dateTime = `${date}T10:00:00.00Z`
        getPlaces(this.props.id, dateTime, 1).then((data) => {
            if (!data) {
                return
            }
            try {
                const places = []
                data.forEach(e => {
                    if (data) {
                        const placeElement = {
                            id: e.id,
                            left: e.leftTop,
                            top: e.rightBottom,
                            width: e.width,
                            height:e.height,
                            nubmer: e.number,
                            capacity: e.capacity,
                            floor: e.floor,
                            isBooked: e.isBooked,
                        };
                        places.push(placeElement);
                    }
                });
                this.setState({
                    places: places,
                    currentFloor: 1,
                    isMouseDown: false,
                    currentIndex: -1,
                    currentTime: 'T10:00:00.00Z',
                    currentDate: date,
                });
              } catch (e) {
                if (e !== BreakException) throw e;
              }
              this.updateCanvas();
        });
    }

    updateCanvas() {
        const canv = this.refs.canvas;
        const ctx = canv.getContext('2d');
        canv.style.width  = '600px';
        canv.style.height = '600px';
        
        const  roundRect = (ctx, x, y, width, height, radius, fill, stroke) => {
            if (typeof stroke === 'undefined') {
              stroke = true;
            }
            if (typeof radius === 'undefined') {
              radius = 5;
            }
            if (typeof radius === 'number') {
              radius = {tl: radius, tr: radius, br: radius, bl: radius};
            } else {
              var defaultRadius = {tl: 0, tr: 0, br: 0, bl: 0};
              for (var side in defaultRadius) {
                radius[side] = radius[side] || defaultRadius[side];
              }
            }
            ctx.beginPath();
            ctx.moveTo(x + radius.tl, y);
            ctx.lineTo(x + width - radius.tr, y);
            ctx.quadraticCurveTo(x + width, y, x + width, y + radius.tr);
            ctx.lineTo(x + width, y + height - radius.br);
            ctx.quadraticCurveTo(x + width, y + height, x + width - radius.br, y + height);
            ctx.lineTo(x + radius.bl, y + height);
            ctx.quadraticCurveTo(x, y + height, x, y + height - radius.bl);
            ctx.lineTo(x, y + radius.tl);
            ctx.quadraticCurveTo(x, y, x + radius.tl, y);
            ctx.closePath();
            if (fill) {
              ctx.fill();
            }
            if (stroke) {
              ctx.stroke();
            }
          }

        const clear = (ctx) => {
            ctx.fillStyle = 'white';
            ctx.fillRect(0, 0, 500, 500);
            ctx.beginPath();
            ctx.fillStyle = 'black';
        };

        const currentPlace = (event) => {
            const places = this.state.places;
            const rect = canv.getBoundingClientRect();
            const x = Math.floor((event.clientX - rect.left)/2);
            const y = Math.floor((event.clientY - rect.top)/2);
            for (const [index, e] of places.entries()) {
                if (
                    x > e.left &&
                    y > e.top &&
                    x < e.left + e.width &&
                    y < e.top + e.height
                ) {
                    return index;
                }
            }

            return -1;
        };

        canv.addEventListener('mousedown', (e) => {
            const curr = currentPlace(e);
            if (curr !== -1 && !this.state.places[curr].isBooked) {
                const id = this.state.places[curr].id;
                this.setState({
                    isMouseDown: true,
                    currentIndex: curr,
                    currentId: id,
                });
                clear(ctx);
                paintRects();
            } else {
                this.setState({
                    isMouseDown: false,
                    currentIndex: -1,
                });
            }
        })

        const paintRects = () => {
            clear(ctx);
            const places = this.state.places;
            places.forEach(e => {
                if (e.floor == this.state.currentFloor){
                    roundRect(
                        ctx,
                        e.left,
                        e.top, 
                        e.width, 
                        e.height,
                        10
                    )
                    if (e.isBooked) {
                        ctx.fillStyle = 'gray';
                        ctx.fill();
                        ctx.fillStyle = 'black';
                    }
                    if (!e.isBooked && e.id === this.state.currentId) {
                        ctx.fillStyle = 'green';
                        ctx.fill();
                        ctx.fillStyle = 'black';
                    }
                    ctx.fillText(e.capacity, e.left + Math.floor(e.width/2), e.top + Math.floor(e.height/2))
                }
            });
        };

        paintRects();
    }

    floorChange(payload) {
        const floor = parseInt(payload, 10)
        const date = this.state.currentDate;
        const time = this.state.currentTime;
        const dateTime = `${date}${time}`
        getPlaces(this.props.id, dateTime, floor).then((data) => {
            if (!data) {
                return
            }
            try {
                const places = []
                data.forEach(e => {
                    if (data) {
                        const placeElement = {
                            id: e.id,
                            left: e.leftTop,
                            top: e.rightBottom,
                            width: e.width,
                            height:e.height,
                            nubmer: e.number,
                            capacity: e.capacity,
                            floor: e.floor,
                            isBooked: e.isBooked,
                        };
                        places.push(placeElement);
                    }
                });
                this.setState({
                    places: places,
                    currentFloor: payload,
                    isMouseDown: false,
                    currentIndex: -1,
                    currentID: 0,
                });
              } catch (e) {
                if (e !== BreakException) throw e;
              }
              this.updateCanvas();
        });
    }

    timeChange(payload) {
        const floor = parseInt(this.state.currentFloor, 10)
        const date = this.state.currentDate;
        const dateTime = `${date}${payload}`
        getPlaces(this.props.id, dateTime, floor).then((data) => {
            if (!data) {
                return
            }
            try {
                const places = []
                data.forEach(e => {
                    if (data) {
                        const placeElement = {
                            id: e.id,
                            left: e.leftTop,
                            top: e.rightBottom,
                            width: e.width,
                            height:e.height,
                            nubmer: e.number,
                            capacity: e.capacity,
                            floor: e.floor,
                            isBooked: e.isBooked,
                        };
                        places.push(placeElement);
                    }
                });
                this.setState({
                    places: places,
                    currentTime: payload,
                    currentDate: date,
                    isMouseDown: false,
                    currentIndex: -1,
                    currentID: 0,
                });
              } catch (e) {
                if (e !== BreakException) throw e;
              }
              this.updateCanvas();
        });
    }

    dateChange(payload) {
        const floor = parseInt(this.state.currentFloor, 10)
        const time = this.state.currentTime;
        const dateTime = `${payload}${time}`
        getPlaces(this.props.id, dateTime, floor).then((data) => {
            if (!data) {
                return
            }
            try {
                const places = []
                data.forEach(e => {
                    if (data) {
                        const placeElement = {
                            id: e.id,
                            left: e.leftTop,
                            top: e.rightBottom,
                            width: e.width,
                            height:e.height,
                            nubmer: e.number,
                            capacity: e.capacity,
                            floor: e.floor,
                            isBooked: e.isBooked,
                        };
                        places.push(placeElement);
                    }
                });
                this.setState({
                    places: places,
                    currentTime: time,
                    currentDate: payload,
                    isMouseDown: false,
                    currentIndex: -1,
                    currentID: 0,
                });
              } catch (e) {
                if (e !== BreakException) throw e;
              }
              this.updateCanvas();
        });
    }


    back(payload) {
        this.props.reserveClick(payload);
    }

    orderClick() {
        if (this.state.currentId == 0) {
            return
        }
        const time = this.state.currentTime;
        const date = this.state.currentDate;
        const startTime = `${date}${time}`
        createOrder(this.props.userId, this.state.currentId, startTime).then((data) => {
            this.back(false);
        });
    }

    render(){
        const today = new Date();
        const month = today.getMonth()+1;
        const day = today.getDate();
        const date = today.getFullYear()+'-'+(this.formate(month))+'-'+this.formate(day);

        const nextDay = new Date(new Date().getTime() + 24 * 60 * 60 * 1000);
        const month2 = today.getMonth()+1;
        const day2 = today.getDate();
        const nextDate = nextDay.getFullYear()+'-'+(this.formate(month2))+'-'+this.formate(day2);

        const hours = new Date().getHours();
        var start =  10;
        if (hours > 10) {
        start = hours;
        }

        const times = [];
        while (start <= 22) {
            const value = `T${start}:00:00.00Z`
            const el = `${start}:00`
            times.push({
                value: value,
                el: el,
            })
            start = start + 2;
        }

        const timesList = times.map((time, index) =>
            <option key={index+1} value={time.value}>{time.el}</option>
        );

        return (
            <div className='reserve'>
                <ReserveTitleBar reserveClick={this.props.reserveClick} />
                <div className='reserve__floor-time'>
                <select onChange={(e) => this.floorChange(e.target.value)} value={this.state.currentFloor}>
                    <option value={1}>1</option>
                    <option value={2}>2</option>
                </select>
                <select onChange={(e) => this.dateChange(e.target.value)}>
                    <option value={date}>Сегодня</option>
                    <option value={nextDate}>Завтра</option>
                </select>
                <select onChange={(e) => this.timeChange(e.target.value)}>
                    {timesList}
                </select>
                </div>
                <div className='reserve__canvas'>
                    <canvas ref="canvas" width={300} height={300}/>
                </div>
                <div className='reserve__order-button'>
                    <OrderButton orderClick={this.orderClick} text={'Забронировать'} />
                </div>
            </div>
        );
    }
}