import React from 'react'

import '../../css/Common/InputField.css'

export default class InputField extends React.Component {
    constructor(props) {
        super(props);

        if (props.name) {
            this.state = {name: props.name}
        } else {
            this.state = {name: 'default'}
        }
        if (props.label) {
            this.state['label'] = props.label
        } else {
            this.state['label'] = 'default'
        }
        if (props.placeholder) {
            this.state['placeholder'] = props.placeholder
        } else {
            this.state['placeholder'] = 'default'
        }
        if (props.type) {
            this.state['type'] = props.type
        } else {
            this.state['type'] = 'default'
        }
    }

    // TODO: handlers for focus and input in fields

    render(){
        return (
            <div className='input-field'>
                <label htmlFor={this.state.name} className='input-field__label'>{this.state.label}</label>
                <input id={this.state.name} className='input-field__input' type={this.state.type} placeholder={this.state.placeholder}></input>
            </div>
        );
    }
}
