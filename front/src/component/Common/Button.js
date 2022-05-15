import React from 'react'

import '../../css/Common/Button.css'

export default class Button extends React.Component {
    constructor(props) {
        super(props);

        if (props.text) {
            this.state = {buttonText: props.text}
        } else {
            this.state = {buttonText: 'default'}
        }
        if (props.onClick) {
            this.state['onClick'] = props.onClick
        } else {
            // error
        }
    }

    buttonClick = (payload) => {
        this.setState({login: true, signup: false})
        console.log('login button touched')
    }

    render(){
        return (
            <button className='button' onClick={this.state.onClick}>
                <span className='button__text'>{this.state.buttonText}</span>
            </button>
        );
    }
}
