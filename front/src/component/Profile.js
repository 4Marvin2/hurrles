import React from 'react'
import { nameErorrMsg, emailErrorMsg, phoneNumberErrorMsg, profileFormErrorMsg } from '../constants/errorMsg';
import InputField from './Common/InputField';
import Button from './Common/Button';
import validator from 'validator'
import { editRequest, getUserRequest } from '../requests/user';
import '../css/Profile.css'

export default class Profile extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            name: props.name,
            nameIsValid: true,
            email: props.email,
            emailIsValid: true,
            phoneNumber: props.phoneNumber,
            phoneNumberIsValid: true,
            formError: profileFormErrorMsg,
            formErrorIsActive: false,
        }
    }

    nameOnInput = (payload) => {
        this.setState({name: payload.target.value})
        if (payload.target.value.length === 0) {
            this.setState({nameIsValid: false})
        } else {
            this.setState({nameIsValid: true})
        }
    }

    emailOnInput = (payload) => {
        this.setState({email: payload.target.value})
        if (validator.isEmail(payload.target.value)) {
            this.setState({emailIsValid: true})
        } else {
            this.setState({emailIsValid: false})
        }
    }

    phoneNumberOnInput = (payload) => {
        this.setState({phoneNumber: payload.target.value})
        if (validator.isMobilePhone(payload.target.value, 'ru-RU') || payload.target.value.length === 0) {
            this.setState({phoneNumberIsValid: true})
        } else {
            this.setState({phoneNumberIsValid: false})
        }
    }

    formButtonClick = (payload) => {
        if (this.state.nameIsValid && this.state.emailIsValid && this.state.phoneNumberIsValid) {
            editRequest(this.state.name, this.state.email, this.state.phoneNumber)
                .then((data) => {
                    this.setState({formErrorIsActive: false})
                    // TODO: set user data to store
                    console.log(data);
                })
                .catch((error) => {
                    console.log(error)
                    this.setState({formErrorIsActive: true})
                })
        } else {
            // error
            console.log('login error');
        }
    }

    render(){
        return (
            <div className='profile'>
                <InputField
                    type='text'
                    name='username'
                    label='Имя'
                    placeholder='Ваше имя'
                    value={this.state.name}
                    prompt={nameErorrMsg}
                    promptIsActive={!this.state.nameIsValid}
                    maxLength={128}
                    onInput={(e) => {this.nameOnInput(e)}}
                    theme='light'
                />
                <InputField
                    type='email'
                    name='email'
                    label='Почта'
                    placeholder='hurrles@example.com'
                    value={this.state.email}
                    prompt={emailErrorMsg}
                    promptIsActive={!this.state.emailIsValid}
                    maxLength={100}
                    onInput={(e) => {this.emailOnInput(e)}}
                    theme='light'
                />
                <InputField
                    type='tel'
                    name='phoneNumber'
                    label='Номер телефона'
                    placeholder='Ваш номер'
                    value={this.state.phoneNumber}
                    prompt={phoneNumberErrorMsg}
                    promptIsActive={!this.state.phoneNumberIsValid}
                    maxLength={16}
                    onInput={(e) => {this.phoneNumberOnInput(e)}}
                    theme='light'
                />
                <div className='profile__form-button'>
                    <span className={ this.state.formErrorIsActive ? 'profile__form-error profile__form-error_active' : 'profile__form-error' }>{this.state.formError}</span>
                    <Button text='Изменить' theme='light' width={100} height={32} onClick={this.formButtonClick}/>
                </div>
            </div>
        );
    }
}
