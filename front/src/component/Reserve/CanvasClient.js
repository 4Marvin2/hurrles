import React from "react";

export default class CanvasClient extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            places: this.props.places,
            currentFloor: 1,
            isMouseDown: false,
            currentIndex: -1,
            offsetX: 0,
            offsetY: 0,
        }
    }
    componentDidUpdate(){
        const comp = JSON.stringify(this.props.places) === JSON.stringify(this.state.places);
        if (this.props.currentFloor !== this.state.currentFloor) {
            this.setState({currentFloor: this.props.currentFloor});
            setTimeout(() => {
                this.updateCanvas();
            }, 100);
        }
    }
    componentDidMount() {
        this.updateCanvas();
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
                    currentIndex: curr,
                });
            }
        })

        canv.addEventListener('mouseup', (e) => {
            this.setState({
                isMouseDown: false,
            });
            clear(ctx);
            paintRects();
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
                    ctx.fillText(e.capacity, e.left + Math.floor(e.width/2), e.top + Math.floor(e.height/2))
                }
            });
        };

        paintRects();
    }
    render() {
        return (
            <canvas ref="canvas" width={300} height={300}/>
        );
    }
}