import React from "react";

export default class Canvas extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            places: [
                {
                    left: 100,
                    top: 100,
                    width: 100,
                    height:50,
                },
                {
                    left: 0,
                    top: 0,
                    width: 100,
                    height:50,
                }
            ],
            isMouseDown: false,
            currentIndex: -1,
            offsetX: 0,
            offsetY: 0,
        }
    }
    componentDidMount() {
        this.updateCanvas();
    }
    updateCanvas() {
        const canv = this.refs.canvas;
        const ctx = canv.getContext('2d');
        
        const clear = (ctx) => {
            ctx.fillStyle = 'white';
            ctx.fillRect(0, 0, 500, 500);
            ctx.beginPath();
            ctx.fillStyle = 'black';
        };

        const currentPlace = (event) => {
            const places = this.state.places;
            const x = Math.floor(event.clientX/2);
            const y = Math.floor(event.clientY/2);
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
            if (curr !== -1) {
                this.setState({
                    isMouseDown: true,
                    currentIndex: curr,
                    offsetX: Math.floor(e.clientX/2) - this.state.places[curr].left,
                    offsetY: Math.floor(e.clientY/2) - this.state.places[curr].top,
                });
            } else {
                this.setState({
                    isMouseDown: false,
                    currentIndex: curr,
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
        

        const savePlaces = (event) => {
            const idx = this.state.currentIndex;
            if (idx === -1) return;
            const places = this.state.places;
            const newPlaces = [];
            places.forEach((e, index) => {
                if (index === idx) {
                    const newPlace = {
                        left: Math.floor(event.clientX/2) - this.state.offsetX,
                        top: Math.floor(event.clientY/2) - this.state.offsetY,
                        width: e.width,
                        height:e.height,
                    }
                    newPlaces.push(newPlace)
                } else {
                    newPlaces.push(e)
                }
            });
            this.setState({places: newPlaces})
        };

        const changeRect = (event) => {
            const idx = this.state.currentIndex;
            if (idx === -1) return;
            const places = this.state.places;
            places.forEach((e, index) => {
                if (index === idx) {
                    ctx.strokeRect(
                        Math.floor(event.clientX/2) - this.state.offsetX,
                        Math.floor(event.clientY/2) - this.state.offsetY,  
                        e.width, 
                        e.height
                    );
                } else {
                    ctx.strokeRect(
                        e.left,
                        e.top, 
                        e.width, 
                        e.height
                    );
                }
            });
        };

        const paintRects = () => {
            const places = this.state.places;
            places.forEach(e => {
                ctx.strokeRect(
                    e.left,
                    e.top, 
                    e.width, 
                    e.height
                );
            });
        };

        paintRects();
    }
    render() {
        return (
            <canvas ref="canvas" width={500} height={500}/>
        );
    }
}