import React, { useDebugValue } from 'react'

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
        if (props.disableCopy) {
            this.state['disableCopy'] = props.disableCopy
        } else {
            this.state['disableCopy'] = false
        }
        if (props.disablePaste) {
            this.state['disablePaste'] = props.disablePaste
        } else {
            this.state['disablePaste'] = false
        }
        if (props.prompt) {
            this.state['prompt'] = props.prompt
        } else {
            this.state['prompt'] = 'default'
        }
        if (props.promptIsActive) {
            this.state['promptIsActive'] = props.promptIsActive
        } else {
            this.state['promptIsActive'] = false
        }
        if (props.onInput) {
            this.state['onInput'] = props.onInput
        } else {
            // error
        }
    }

    shouldComponentUpdate = (newProps) => {
        if (this.state.promptIsActive !== newProps.promptIsActive && newProps.promptIsActive !== undefined) {
            this.setState({promptIsActive: newProps.promptIsActive})
            return true
        }
        return false
    }

    render(){
        return (
            <div className='input-field'>
                <label htmlFor={this.state.name} className='input-field__label'>{this.state.label}</label>
                { this.state.disableCopy && this.state.disablePaste &&
                    <input
                        id={this.state.name}
                        className={ this.state.promptIsActive ? 'input-field__input input-field__input_error' : 'input-field__input' }
                        type={this.state.type}
                        placeholder={this.state.placeholder}
                        onInput={this.state.onInput}
                        onPaste={(e) => {
                            e.preventDefault();
                            return false;
                        }}
                        onCopy={(e) => {
                            e.preventDefault();
                            return false;
                        }}
                    />
                }
                { !this.state.disableCopy && this.state.disablePaste &&
                    <input
                        id={this.state.name}
                        className={ this.state.promptIsActive ? 'input-field__input input-field__input_error' : 'input-field__input' }
                        type={this.state.type}
                        placeholder={this.state.placeholder}
                        onInput={this.state.onInput}
                        onPaste={(e) => {
                            e.preventDefault();
                            return false;
                        }}
                    />
                }
                { this.state.disableCopy && !this.state.disablePaste &&
                    <input
                        id={this.state.name}
                        className={ this.state.promptIsActive ? 'input-field__input input-field__input_error' : 'input-field__input' }
                        type={this.state.type}
                        placeholder={this.state.placeholder}
                        onInput={this.state.onInput}
                        onCope={(e) => {
                            e.preventDefault();
                            return false;
                        }}
                    />
                }
                { !this.state.disableCopy && !this.state.disablePaste &&
                    <input
                        id={this.state.name}
                        className={ this.state.promptIsActive ? 'input-field__input input-field__input_error' : 'input-field__input' }
                        type={this.state.type}
                        placeholder={this.state.placeholder}
                        onInput={this.state.onInput}
                    />
                }
                <span className={ this.state.promptIsActive ? 'input-field__prompt input-field__prompt_active' : 'input-field__prompt' }>{this.state.prompt}</span>
            </div>
        );
    }
}
