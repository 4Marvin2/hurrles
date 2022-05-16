import React from 'react'

import Logo from '../../imgs/logo.svg';

import '../../css/AuthForm.css'

import InputField from '../Common/InputField';
import Button from '../Common/Button';

import validator from 'validator'

import { getUserRequest, loginRequest, signupRequest } from '../../requests/requests';

export default class AuthForm extends React.Component {
    constructor(props) {
        super(props);
        // cookie request
        // if ok go to main page
        if (props.authType && props.authType === 'signup') {
            this.state = {login: false, signup: true}
        } else {
            this.state = {login: true, signup: false}
        }
    }

    componentDidMount = () => {
        getUserRequest().then((data) => {
            console.log(data)
        });
    }

    loginSwitchClick = (payload) => {
        this.setState({login: true, signup: false})
        console.log('login button touched')
    }

    signupSwitchClick = (payload) => {
        this.setState({login: false, signup: true})
        console.log('signup button touched')
    }

    emailOnChange = (payload) => {
        this.setState({email: payload.target.value})
        if (validator.isEmail(payload.target.value)) {
            this.setState({emailIsValid: true})
        } else {
            this.setState({emailIsValid: false})
        }
    }

    passwordOnChange = (payload) => {
        this.setState({password: payload.target.value})
        // if (validator.isStrongPassword(payload.target.value)) {
        if (payload.target.value) {
            this.setState({passwordIsValid: true})
        } else {
            this.setState({passwordIsValid: false})
        }
    }

    repeatPasswordOnChange = (payload) => {
        this.setState({repeatPassword: payload.target.value})
        if (this.state.passwordIsValid !== payload.target.value) {
            this.setState({repeatPasswordIsValid: true})
        } else {
            this.setState({repeatPasswordIsValid: false})
        }
    }

    formButtonClick = (payload) => {
        if (this.state.login) {
            // login request
            if (this.state.emailIsValid && this.state.passwordIsValid) {
                loginRequest(this.state.email, this.state.password).then((data) => {
                    console.log(data)
                })
            } else {
                // error
                console.log('login error')
            }
        } else if (this.state.signup) {
            // signup request
            if (this.state.emailIsValid && this.state.passwordIsValid && this.state.repeatPasswordIsValid) {
                signupRequest(this.state.email, this.state.password).then((data) => {
                    console.log(data)
                })
            } else {
                // error
                console.log('signup error')
            }
        } else {
            // error
            console.log('error')
        }
    }

    render(){
        return (
            <div className='auth-form'>
                <div className='auth-form__card'>
                    <div className='auth-form__header'>
                        <div className='auth-form__logo'>
                            <img src={Logo} className='auth-form__logo-img' alt='logo'/>
                            <span className='auth-form__logo-text'>hurrles</span>
                        </div>
                        <div className='auth-form__switches'>
                            <button className={ this.state.login ? 'auth-form__login-switch auth-form__login-switch_active' : 'auth-form__login-switch' } onClick={this.loginSwitchClick}>
                                <span className='auth-form__login-switch-text'>Вход</span>
                            </button>
                            <button className={ this.state.signup ? 'auth-form__signup-switch auth-form__signup-switch_active' : 'auth-form__signup-switch' } onClick={this.signupSwitchClick}>
                                <span className='auth-form__signup-switch-text'>Регистрация</span>
                            </button>
                        </div>
                    </div>
                    <div className='auth-form__fields'>
                        <InputField type='email' name='email' label='Почта или номер телефона' placeholder='hurrles@example.com' onChange={(e) => {this.emailOnChange(e)}}/>
                        <InputField type='password' name='password' label='Пароль' placeholder='password' onChange={(e) => {this.passwordOnChange(e)}}/>
                        { this.state.signup &&
                            <InputField type='password' name='repeatPassword' label='Повторите пароль' placeholder='repeat password' onChange={(e) => {this.repeatPasswordOnChange(e)}}/>
                        }
                    </div>
                    {/* forggot password link here */}
                    <div className='auth-form__form-button'>
                        <Button text='Войти' onClick={this.formButtonClick}/>
                    </div>
                </div>
            </div>
        );
    }
}
