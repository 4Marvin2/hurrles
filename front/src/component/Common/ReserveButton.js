import React from 'react'

import '../../css/Common/ReserveButton.css'

export default class AddButton extends React.Component {
    constructor(props) {
        super(props);

        if (props.text) {
            this.state = {text: props.text}
        } else {
            this.state = {text: 'default'}
        }
        if (props.onClick) {
            this.state['onClick'] = props.onClick
        } else {
            // error
        }
    }

    reserveClick(payload) {
        this.props.reserveClick(payload);
    }

    shouldComponentUpdate = (newProps) => {
        if (this.state.text !== newProps.text && newProps.text !== undefined) {
            this.setState({text: newProps.text})
            return true
        }
        return false
    }

    render(){
        return (
            <button className='reserve-button' onClick={() => this.reserveClick(true)}>
                <span className='reserve-button__text'><h3>{this.state.text}</h3></span>
            </button>
        );
    }
}