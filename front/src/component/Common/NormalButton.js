import React from 'react'

import '../../css/Common/NormalButton.css'

export default class Button extends React.Component {
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

    shouldComponentUpdate = (newProps) => {
        if (this.state.text !== newProps.text && newProps.text !== undefined) {
            this.setState({text: newProps.text})
            return true
        }
        return false
    }

    render(){
        return (
            <button className='button' onClick={this.state.onClick}>
                <span className='button__text'>{this.state.text}</span>
            </button>
        );
    }
}