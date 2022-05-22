import React from 'react'

import '../../css/Common/Button.css'

export default class Button extends React.Component {
    constructor(props) {
        super(props);

        if (props.text) {
            this.state = {text: props.text}
        } else {
            this.state = {text: 'default'}
        }
        if (props.theme) {
            if (props.theme === 'light') {
                this.state['buttonClassName'] = 'button button_light'
                this.state['buttonTextClassName'] = 'button__text button__text_light'
            } else if (props.theme === 'dark') {
                this.state['buttonClassName'] = 'button button_dark'
                this.state['buttonTextClassName'] = 'button__text button__text_dark'
            } else {
                this.state['buttonClassName'] = 'button'
                this.state['buttonTextClassName'] = 'button__text'
            }
        } else {
            this.state['buttonClassName'] = 'button'
            this.state['buttonTextClassName'] = 'button__text'
        }
        if (props.height) {
            this.state['height'] = props.height
        } else {
            this.state['height'] = 100
        }
        if (props.width) {
            this.state['width'] = props.width
        } else {
            this.state['width'] = 100
        }
        if (props.id) {
            this.state['id'] = props.id
        } else {
            this.state['id'] = 'default'
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
            <button className={this.state.buttonClassName} id={this.state.id} onClick={this.state.onClick} style={{width: this.state.width, height: this.state.height}}>
                <span className={this.state.buttonTextClassName}>{this.state.text}</span>
            </button>
        );
    }
}
