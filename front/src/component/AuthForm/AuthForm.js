import React from 'react'

import Logo from '../../imgs/logo.svg';

import '../../css/AuthForm.css'

import InputField from '../Common/InputField';
import Button from '../Common/Button';

import validator from 'validator'

import { loginRequest, signupRequest } from '../../requests/auth';

import { emailErrorMsg, passwordErrorMsg, repeatPasswordErrorMsg, loginFormErrorMsg, signupFormErrorMsg } from '../../constants/errorMsg';

export default class AuthForm extends React.Component {
    constructor(props) {
        super(props);
        if (props.authType && props.authType === 'signup') {
            this.state = {
                login: false,
                signup: true,
                emailIsValid: true,
                passwordIsValid: true,
                repeatPasswordIsValid: true,
                formError: signupFormErrorMsg,
                formErrorIsActive: false,
                buttonText: 'Зарегистрироваться',
            };
        } else {
            this.state = {
                login: true,
                signup: false,
                emailIsValid: true,
                passwordIsValid: true,
                formErrorIsActive: false,
                formError: loginFormErrorMsg,
                buttonText: 'Войти',
            };
        }
        if (props.callback) {
            this.state['callback'] = props.callback
        } else {
            // error
        }
    }

    logInHandler(payload) {
        this.props.logInHandler(payload);
    }

    loginSwitchClick = (payload) => {
        this.setState({
            login: true,
            signup: false,
            formError: loginFormErrorMsg,
            formErrorIsActive: false,
            buttonText: 'Войти',
        });
    }

    signupSwitchClick = (payload) => {
        this.setState({
            login: false,
            signup: true,
            repeatPasswordIsValid: true,
            formError: signupFormErrorMsg,
            formErrorIsActive: false,
            buttonText: 'Зарегистрироваться',
        });
    }

    emailOnInput = (payload) => {
        this.setState({email: payload.target.value});
        if (validator.isEmail(payload.target.value)) {
            this.setState({emailIsValid: true});
        } else {
            this.setState({emailIsValid: false});
        }
    }

    passwordOnInput = (payload) => {
        this.setState({password: payload.target.value});
        if (validator.isStrongPassword(payload.target.value)) {
            this.setState({passwordIsValid: true});
        } else {
            this.setState({passwordIsValid: false});
        }
    }

    repeatPasswordOnInput = (payload) => {
        this.setState({repeatPassword: payload.target.value});
        if (this.state.password !== payload.target.value) {
            this.setState({repeatPasswordIsValid: false});
        } else {
            this.setState({repeatPasswordIsValid: true});
        }
    }

    formButtonClick = (payload) => {
        if (this.state.login) {
            // login request
            if (this.state.emailIsValid && this.state.passwordIsValid) {
                loginRequest(this.state.email, this.state.password)
                    .then((data) => {
                        this.setState({formErrorIsActive: false})
                        // TODO: set user data to store
                        console.log(data);
                        this.logInHandler({loggedIn: true})
                    })
                    .catch((error) => {
                        this.setState({formErrorIsActive: true})
                    })
            } else {
                // error
                console.log('login error');
            }
        } else if (this.state.signup) {
            // signup request
            if (this.state.emailIsValid && this.state.passwordIsValid && this.state.repeatPasswordIsValid) {
                signupRequest(this.state.email, this.state.password)
                .then((data) => {
                    this.setState({formErrorIsActive: false})
                    // TODO: set user data to store
                    console.log(data);
                    this.logInHandler({loggedIn: true})
                })
                .catch((error) => {
                    this.setState({formErrorIsActive: true})
                })
            } else {
                // error
                console.log('signup error');
            }
        } else {
            // error
            console.log('error');
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
                            <button className={ this.state.login ? 'auth-form__switch auth-form__switch_active' : 'auth-form__switch' } onClick={this.loginSwitchClick}>
                                <span className='auth-form__switch-text'>Вход</span>
                            </button>
                            <button className={ this.state.signup ? 'auth-form__switch auth-form__switch_active' : 'auth-form__switch' } onClick={this.signupSwitchClick}>
                                <span className='auth-form__switch-text'>Регистрация</span>
                            </button>
                        </div>
                    </div>
                    <div className='auth-form__fields'>
                        <InputField
                            type='email'
                            name='email'
                            label='Почта'
                            placeholder='hurrles@example.com'
                            prompt={emailErrorMsg}
                            promptIsActive={!this.state.emailIsValid}
                            onInput={(e) => {this.emailOnInput(e)}}
                        />
                        <InputField
                            type='password'
                            name='password'
                            label='Пароль'
                            placeholder='password'
                            prompt={passwordErrorMsg}
                            promptIsActive={!this.state.passwordIsValid}
                            onInput={(e) => {this.passwordOnInput(e)}}
                            disableCope={true}
                            disablePaste={true}
                        />

                        { this.state.signup &&
                            <InputField
                                type='password'
                                name='repeatPassword'
                                label='Повторите пароль'
                                placeholder='repeat password'
                                prompt={repeatPasswordErrorMsg}
                                promptIsActive={!this.state.repeatPasswordIsValid}
                                onInput={(e) => {this.repeatPasswordOnInput(e)}}
                                disableCope={true}
                                disablePaste={true}
                            />
                        }
                    </div>
                    {/* TODO: forggot password link here */}
                    <div className='auth-form__form-button'>
                        <span className={ this.state.formErrorIsActive ? 'auth-form__form-error auth-form__form-error_active' : 'auth-form__form-error' }>{this.state.formError}</span>
                        <Button text={this.state.buttonText} onClick={this.formButtonClick}/>
                    </div>
                </div>
            </div>
        );
    }
}