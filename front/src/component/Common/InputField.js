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
            this.state['placeholder'] = ''
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
        if (props.value) {
            this.state['value'] = props.value
        } else {
            this.state['value'] = ''
        }
        if (props.pattern) {
            this.state['pattern'] = props.pattern
        } else {
            this.state['pattern'] = ''
        }
        if (props.rows) {
            this.state['rows'] = props.rows
        } else {
            this.state['rows'] = 30
        }
        if (props.cols) {
            this.state['cols'] = props.cols
        } else {
            this.state['cols'] = 30
        }
        if (props.maxLength) {
            this.state['maxLength'] = props.maxLength
        } else {
            this.state['maxLength'] = 50
        }
        if (props.step) {
            this.state['step'] = props.step
        } else {
            this.state['step'] = 1
        }
        if (props.theme) {
            if (props.theme === 'light') {
                if (props.type === 'textarea') {
                    this.state['inputClassName'] = 'input-field__textarea input-field__textarea_light'   
                } else {
                    this.state['inputClassName'] = 'input-field__input input-field__input_light'
                }
                this.state['labelClassName'] = 'input-field__label input-field__label_light'
            } else if (props.theme === 'dark') {
                if (props.type === 'textarea') {
                    this.state['inputClassName'] = 'input-field__textarea input-field__textarea_dark'   
                } else {
                    this.state['inputClassName'] = 'input-field__input input-field__input_dark'
                }
                this.state['labelClassName'] = 'input-field__label input-field__label_dark'
            } else {
                if (props.type === 'textarea') {
                    this.state['inputClassName'] = 'input-field__textarea'   
                } else {
                    this.state['inputClassName'] = 'input-field__input'
                }
                this.state['labelClassName'] = 'input-field__label input-field__label'
            }
        } else {
            if (props.type === 'textarea') {
                this.state['inputClassName'] = 'input-field__textarea input-field__textarea_light'   
            } else {
                this.state['inputClassName'] = 'input-field__input input-field__input_light'
            }
            this.state['labelClassName'] = 'input-field__label input-field__label_light'
        }
        if (props.onInput) {
            this.state['onInput'] = props.onInput
        } else {
            // error
        }
        if (props.onFocus) {
            this.state['onFocus'] = props.onFocus
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
                <label htmlFor={this.state.name} className={this.state.labelClassName}>{this.state.label}</label>
                { this.state.disableCopy && this.state.disablePaste && this.state.type !== 'textarea' &&
                    <input
                        id={this.state.name}
                        className={ this.state.promptIsActive ? this.state.inputClassName + ' input-field__input_error' : this.state.inputClassName }
                        type={this.state.type}
                        placeholder={this.state.placeholder}
                        defaultValue={this.state.value}
                        pattern={this.state.pattern}
                        maxLength={this.state.maxLength}
                        step={this.state.step}
                        onInput={this.state.onInput}
                        onFocus={this.state.onFocus}
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
                { !this.state.disableCopy && this.state.disablePaste && this.state.type !== 'textarea' &&
                    <input
                        id={this.state.name}
                        className={ this.state.promptIsActive ? this.state.inputClassName + ' input-field__input_error' : this.state.inputClassName }
                        type={this.state.type}
                        placeholder={this.state.placeholder}
                        defaultValue={this.state.value}
                        pattern={this.state.pattern}
                        maxLength={this.state.maxLength}
                        step={this.state.step}
                        onInput={this.state.onInput}
                        onFocus={this.state.onFocus}
                        onPaste={(e) => {
                            e.preventDefault();
                            return false;
                        }}
                    />
                }
                { this.state.disableCopy && !this.state.disablePaste && this.state.type !== 'textarea' &&
                    <input
                        id={this.state.name}
                        className={ this.state.promptIsActive ? this.state.inputClassName + ' input-field__input_error' : this.state.inputClassName }
                        type={this.state.type}
                        placeholder={this.state.placeholder}
                        defaultValue={this.state.value}
                        pattern={this.state.pattern}
                        maxLength={this.state.maxLength}
                        step={this.state.step}
                        onInput={this.state.onInput}
                        onFocus={this.state.onFocus}
                        onCope={(e) => {
                            e.preventDefault();
                            return false;
                        }}
                    />
                }
                { !this.state.disableCopy && !this.state.disablePaste && this.state.type !== 'textarea' &&
                    <input
                        id={this.state.name}
                        className={ this.state.promptIsActive ? this.state.inputClassName + ' input-field__input_error' : this.state.inputClassName }
                        type={this.state.type}
                        placeholder={this.state.placeholder}
                        defaultValue={this.state.value}
                        pattern={this.state.pattern}
                        maxLength={this.state.maxLength}
                        step={this.state.step}
                        onInput={this.state.onInput}
                        onFocus={this.state.onFocus}
                    />
                }
                { !this.state.disableCopy && !this.state.disablePaste && this.state.type === 'textarea' &&
                    <textarea
                        id={this.state.name}
                        className={ this.state.promptIsActive ? this.state.inputClassName + ' input-field__textarea_error' : this.state.inputClassName }
                        placeholder={this.state.placeholder}
                        defaultValue={this.state.value}
                        pattern={this.state.pattern}
                        rows={this.state.rows}
                        cols={this.state.cols}
                        maxLength={this.state.maxLength}
                        onInput={this.state.onInput}
                        onFocus={this.state.onFocus}
                    />
                }
                <span className={ this.state.promptIsActive ? 'input-field__prompt input-field__prompt_active' : 'input-field__prompt' }>{this.state.prompt}</span>
            </div>
        );
    }
}
