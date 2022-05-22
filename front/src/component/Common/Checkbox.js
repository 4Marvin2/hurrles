import React from 'react'

import '../../css/Common/CheckboxButton.css'

export default class Checkbox extends React.Component {
    constructor(props) {
        super(props);

        if (props.name) {
            this.state = {name: props.name}
        } else {
            this.state = {name: 'default'}
        }
        if (props.id) {
            this.state['id'] = props.id
        } else {
            this.state['id'] = 'default'
        }
        if (props.value) {
            this.state['value'] = props.value
        } else {
            this.state['value'] = 'default'
        }
        if (props.text) {
            this.state['text'] = props.text
        } else {
            this.state['text'] = 'default'
        }
        if (props.checked) {
            this.state['checked'] = props.checked
        } else {
            this.state['checked'] = false
        }
        if (props.onChange) {
            this.state['onChange'] = props.onChange
        } else {
            // error
        }
    }

    shouldComponentUpdate = (newProps) => {
        if (this.state.checked !== newProps.checked && newProps.checked !== undefined) {
            this.setState({checked: newProps.checked})
            return true
        }
        return false
    }

    render(){
        return (
            <div className='checkbox'>
                { this.state.checked && 
                    <input
                        className='checkbox__input'
                        type='checkbox'
                        id={this.state.id}
                        name={this.state.name}
                        defaultValue={this.state.value}
                        onChange={this.state.onChange}
                        checked
                    />
                }
                { !this.state.checked &&
                    <input
                        className='checkbox__input'
                        type='checkbox'
                        id={this.state.id}
                        name={this.state.name}
                        defaultValue={this.state.value}
                        onChange={this.state.onChange}
                    />
                }
                <span className='checkbox__text'>{this.state.text}</span>
            </div>
        );
    }
}
