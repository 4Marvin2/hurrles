import React from 'react'
import OrderButton from '../Common/OrderButton'
import ReserveTitleBar from '../Reserve/ReserveTitleBar';

import '../../css/AdminCommon/AdminPlaces.css'

import { getAllPlaces } from '../../requests/restors';
import { createPlace, updatePlaces, deletePlace } from '../../requests/adminPlace';
import Button from '../Common/Button'

var BreakException = {};

export default class AdminPlaces extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            places: [],
            currentFloor: 1,
            isMouseDown: false,
            currentIndex: -1,
            currentId: 0,
            offsetX: 0,
            offsetY: 0,
            addHeight: 25,
            addWidth: 25,
            addCapacity: 1,
        }

        this.floorChange = this.floorChange.bind(this);
        this.addPlace = this.addPlace.bind(this);
        this.changeHeight = this.changeHeight.bind(this);
        this.changeWidth = this.changeWidth.bind(this);
        this.changeCapacity = this.changeCapacity.bind(this);
        this.saveClick = this.saveClick.bind(this);
        this.back = this.back.bind(this);
    }

    componentDidMount(){
        getAllPlaces(this.props.id).then((data) => {
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
                            number: e.number,
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
                    addHeight: 25,
                    addWidth: 25,
                    addCapacity: 1,
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
            const rect = canv.getBoundingClientRect();
            const x = Math.floor((e.clientX - rect.left)/2);
            const y = Math.floor((e.clientY - rect.top)/2);
            const curr = currentPlace(e);
            if (curr !== -1) {
                this.setState({
                    isMouseDown: true,
                    currentIndex: curr,
                    offsetX: x - this.state.places[curr].left,
                    offsetY: y - this.state.places[curr].top,
                });
            } else {
                this.setState({
                    isMouseDown: false,
                    currentIndex: -1,
                });
            }
        })

        canv.addEventListener('mouseup', (e) => {
            this.setState({
                isMouseDown: false,
            });
            savePlaces(e);
            clear(ctx);
            paintRects();
        })

        canv.addEventListener('mousemove', (e) => {
            if (this.state.isMouseDown) {
                clear(ctx);
                changeRect(e);
            }
        })

        canv.addEventListener('dblclick', (e) => {
            const curr = currentPlace(e);
            if (curr !== -1) {                
                deletePlace(this.state.places[curr].id).then((data) => {
                    if (!data) {
                        return
                    }
                    const places = this.state.places;
                    places.splice(curr, 1);
                    this.setState({places: places})
                    setTimeout(() => {
                        paintRects(e);
                    }, 100);
                });
            }
        })


        const savePlaces = (event) => {
            const rect = canv.getBoundingClientRect();
            const x = Math.floor((event.clientX - rect.left)/2);
            const y = Math.floor((event.clientY - rect.top)/2);
            const idx = this.state.currentIndex;
            if (idx === -1) return;
            const places = this.state.places;
            const savedPlaces = [];
            places.forEach((e, index) => {
                if (index === idx) {
                    const savedPlace = {
                        id: e.id,
                        left: x - this.state.offsetX,
                        top: y - this.state.offsetY,
                        width: e.width,
                        height:e.height,
                        number: e.number,
                        capacity: e.capacity,
                        floor: e.floor,
                        isBooked: e.isBooked,
                    }
                    savedPlaces.push(savedPlace)
                } else {
                    savedPlaces.push(e)
                }
            });
            this.setState({places: savedPlaces})
        };

        const changeRect = (event) => {
            const idx = this.state.currentIndex;
            if (idx === -1) return;
            const places = this.state.places;
            const rect = canv.getBoundingClientRect();
            const x = Math.floor((event.clientX - rect.left)/2) - this.state.offsetX;
            const y = Math.floor((event.clientY - rect.top)/2) - this.state.offsetY;
            const offsetX = x <= 0 ? 0 : x;
            const offsetY = y <= 0 ? 0 : y;
            places.forEach((e, index) => {
                if (e.floor == this.state.currentFloor) {
                    if (index === idx) {
                        const placeX = (offsetX + e.width) >= 300 ? 300 - e.width : offsetX;
                        const placeY = (offsetY + e.height) >= 300 ? 300 - e.height : offsetY;
                        roundRect(
                            ctx,
                            placeX,
                            placeY,  
                            e.width, 
                            e.height,
                            10
                        )
                        ctx.fillText(
                            e.capacity,
                            placeX + Math.floor(e.width/2), 
                            placeY + Math.floor(e.height/2)
                        );
                    } else {
                        roundRect(
                            ctx,
                            e.left,
                            e.top, 
                            e.width, 
                            e.height,
                            10
                        )
                        ctx.fillText(e.capacity, e.left + Math.floor(e.width/2), e.top + Math.floor(e.height/2))
                    }
                }
            });
        };

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
                    ctx.fillText(e.capacity, e.left + Math.floor(e.width/2), e.top + Math.floor(e.height/2))
                }
            });
        };

        paintRects();
    }

    floorChange(payload) {
        this.setState({
            currentFloor: parseInt(payload, 10),
            isMouseDown: false,
            currentIndex: -1,
            currentID: 0,
        });
        setTimeout(() => {
            this.updateCanvas();
        }, 200);
    }

    back() {
        this.props.mapClick(false);
    }

    addPlace() {
        var currNumber = 1;
        const places = this.state.places;
        places.forEach(place => {
            if (place.number > currNumber) {
                currNumber = place.number;
            }
        });

        createPlace(
            this.props.id,
            this.state.addCapacity,
            currNumber + 1,
            Math.floor(Math.random() * 300),
            Math.floor(Math.random() * 300),
            this.state.addWidth,
            this.state.addHeight,
            this.state.currentFloor,
        ).then((data) => {
            if (!data) {
                return
            }
            const newPlace = {
                id: data.id,
                left: data.leftTop,
                top: data.rightBottom,
                width: data.width,
                height: data.height,
                number: data.number,
                capacity: data.capacity,
                floor: data.floor,
                isBooked: data.isBooked,
            }
    
            places.push(newPlace);
    
            this.setState({places: places})
            setTimeout(() => {
                this.updateCanvas();
            }, 200);
        });
    }

    changeHeight(payload) {
        this.setState({addHeight: parseInt(payload)});
    }

    changeWidth(payload) {
        this.setState({addWidth: parseInt(payload)});
    }

    changeCapacity(payload) {
        this.setState({addCapacity: parseInt(payload)});
    }

    saveClick() {
        const places = this.state.places;
        const updatedPlaces = [];

        places.forEach((e) => {
            const updatedPlace = {
                id: e.id,
                restaurantId: this.props.id,
                capacity: e.capacity,
                number: e.number,
                leftTop: e.left,
                rightBottom: e.top,
                width: e.width,
                height: e.height
            }

            updatedPlaces.push(updatedPlace);
        });

        updatePlaces(updatedPlaces).then((data) => {
            if (!data) {
                return
            }
            this.back();
        });
    }

    render(){
        const floorsArray = Array.from({length: this.props.floors}, (_, i) => i + 1)
        const floorsList = floorsArray.map((floor, index) =>
            <option key={index+1} value={floor}>{floor}</option>
        );

        return (
            <div className='reserve'>
                <ReserveTitleBar reserveClick={this.props.mapClick} />
                <div className='reserve__params'>
                    <div className='reserve__param'>
                        Этаж:
                        <select onChange={(e) => this.floorChange(e.target.value)} value={this.state.currentFloor} >
                            {floorsList}
                        </select>
                    </div>
                    <div className='reserve__param'>
                        Длина:
                        <select onChange={(e) => this.changeHeight(e.target.value)} value={this.state.addHeight}>
                            <option value={25}>25</option>
                            <option value={50}>50</option>
                            <option value={100}>100</option>
                            <option value={150}>150</option>
                        </select>
                    </div>
                    <div className='reserve__param'>
                        Ширина:
                        <select onChange={(e) => this.changeWidth(e.target.value)} value={this.state.addWidth}>
                            <option value={25}>25</option>
                            <option value={50}>50</option>
                            <option value={100}>100</option>
                            <option value={150}>150</option>
                        </select>
                    </div>
                    <div className='reserve__param'>
                        Вместимость:
                        <select onChange={(e) => this.changeCapacity(e.target.value)} value={this.state.addCapacity}>
                            <option value={1}>1</option>
                            <option value={2}>2</option>
                            <option value={4}>4</option>
                            <option value={6}>6</option>
                            <option value={8}>8</option>
                        </select>
                    </div>
                    <Button text='Добавить' theme='dark' width={100} height={30} onClick={() => this.addPlace()}/>
                </div>
                <div className='reserve__canvas'>
                    <canvas ref="canvas" width={300} height={300}/>
                </div>
                <div className='reserve__order-button'>
                    <Button text='Сохранить' theme='dark' width={200} height={50} onClick={() => this.saveClick()}/>
                </div>
            </div>
        );
    }
}