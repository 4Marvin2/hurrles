import React from 'react'
import InputField from '../Common/InputField'
import Button from '../Common/Button';

import '../../css/AdminCommon/AddDishForm.css'

import { priceErrorMsg, addRestorFormErrorMsg, descErrorMsg, titleErrorMsg } from '../../constants/errorMsg';
import { addDishRequest, updateDishRequest } from '../../requests/admin';

export default class AddDishForm extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            restorId: props.restorId,
            formIsActive: props.formIsActive,
            titleIsValid: true,
            descIsValid: true,
            priceIsValid: true,
        }
        if (props.dish) {
            this.state['dish'] = props.dish
        } else {
            this.state['dish'] = {
                id: 0,
                title: '',
                desc: '',
                price: 0,
            }
        }
        if (props.isUpdate) {
            this.state['isUpdate'] = props.isUpdate
        } else {
            this.state['isUpdate'] = false
        }
    }

    closeAllForms(payload) {
        this.props.closeAllForms(payload);
    }

    addDishCallback(payload) {
        this.props.addDishCallback(payload);
    }

    titleOnInput = (payload) => {
        this.setState({title: payload.target.value});
        if (payload.target.value.length !== 0) {
            this.setState({titleIsValid: true});
        } else {
            this.setState({titleIsValid: false});
        }
    }

    descOnInput = (payload) => {
        this.setState({desc: payload.target.value});
        if (payload.target.value.length !== 0) {
            this.setState({descIsValid: true});
        } else {
            this.setState({descIsValid: false});
        }
    }

    priceOnInput = (payload) => {
        this.setState({price: Number(payload.target.value)})
        if (payload.target.value.length !== 0 && (payload.target.value * 10) % 10 === 0 && payload.target.value <= 50000) {
            this.setState({priceIsValid: true})
        } else {
            this.setState({priceIsValid: false})
        }
    }

    setDefaultRestorFields = () => {
        let state = true
        if (!this.state.title && this.state.dish.title !== '' && this.state.dish.title !== undefined) {
            this.state.title = this.state.dish.title
        } else if (!this.state.title && (this.state.dish.title === '' || this.state.dish.title === undefined)) {
            this.setState({titleIsValid: false})
            state = false
        }
        if (!this.state.desc && this.state.dish.desc !== '' && this.state.dish.desc !== undefined) {
            this.state.desc = this.state.dish.desc
        } else if (!this.state.desc && (this.state.dish.desc === '' || this.state.dish.desc === undefined)) {
            this.setState({descIsValid: false})
            state = false
        }
        if (!this.state.price && this.state.dish.price !== 0 && this.state.dish.price !== undefined) {
            this.state.price = this.state.dish.price
        } else if (!this.state.price && (this.state.dish.price === 0 || this.state.dish.price === undefined)) {
            this.setState({priceIsValid: false})
            state = false
        }
        return state
    }

    formButtonClick = () => {
        if (this.state.titleIsValid && this.state.descIsValid && this.state.priceIsValid) {
            if (this.state.isUpdate) {
                const ok = this.setDefaultRestorFields()
                if (!ok) {
                    return
                }
                updateDishRequest(
                    this.state.dish.id,
                    this.state.restorId,
                    this.state.title,
                    this.state.desc,
                    this.state.price,
                )
                .then((data) => {
                    this.setState({formErrorIsActive: false})
                    this.addDishCallback({formIsActive: false})
                    this.closeAllForms()
                })
                .catch((error) => {
                    this.setState({formErrorIsActive: true})
                    console.log(error)
                })
            } else {
                addDishRequest(
                    this.state.restorId,
                    this.state.title,
                    this.state.desc,
                    this.state.price,
                )
                .then((data) => {
                    this.setState({formErrorIsActive: false})
                    this.addDishCallback({formIsActive: false})
                    this.closeAllForms()
                })
                .catch((error) => {
                    this.setState({formErrorIsActive: true})
                    console.log(error)
                })
            }
        } else {
            // user see error msg under each fields
        }
    }

    render(){
        return (
            <div className='add-dish-form'>
                {/* exit button */}
                <div className='add-dish-form__fields'>
                        <InputField
                            type='text'
                            name='title'
                            label='Название'
                            value={this.state.dish.title}
                            maxLength={255}
                            prompt={titleErrorMsg}
                            promptIsActive={!this.state.titleIsValid}
                            onInput={(e) => {this.titleOnInput(e)}}
                            theme='dark'
                        />
                        <InputField
                            type='textarea'
                            name='desc'
                            label='Описание'
                            value={this.state.dish.desc}
                            rows={10}
                            cols={100}
                            maxLength={1000}
                            prompt={descErrorMsg}
                            promptIsActive={!this.state.descIsValid}
                            onInput={(e) => {this.descOnInput(e)}}
                            theme='dark'
                        />
                        <InputField
                            type='number'
                            name='price'
                            label='Цена'
                            value={this.state.dish.price}
                            prompt={priceErrorMsg}
                            promptIsActive={!this.state.priceIsValid}
                            onInput={(e) => {this.priceOnInput(e)}}
                            theme='dark'
                        />
                        {/* image */}
                </div>
                <div className='add-restor-form__button'>
                    <span className={ this.state.formErrorIsActive ? 'add-restor-form__form-error add-restor-form__form-error_active' : 'add-restor-form__form-error' }>{addRestorFormErrorMsg}</span>
                    <Button text={this.props.buttonText} theme='dark' width={300} height={72} onClick={this.formButtonClick}/>
                </div>
            </div>
        );
    }
}
