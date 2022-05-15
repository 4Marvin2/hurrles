import React from 'react'

import Logo from '../../imgs/logo.svg';

import '../../css/AuthForm.css'

import InputField from '../Common/InputField';
import Button from '../Common/Button';

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

    loginSwitchClick = (payload) => {
        this.setState({login: true, signup: false})
        console.log('login button touched')
    }

    signupSwitchClick = (payload) => {
        this.setState({login: false, signup: true})
        console.log('signup button touched')
    }

    formButtonClick = (payload) => {
        // TODO: check valid
        if (this.state.login) {
            // login request
            console.log('login req')
        } else if (this.state.signup) {
            // signup request
            console.log('signup req')
        } else {
            // error
        }
    }

    render(){
        return (
            <div className='auth-form'>
                <div className='auth-form__card'>
                    <div className='auth-form__header'>
                        <div className='auth-form__logo'>
                            <img src={Logo} className='auth-form__logo-img'/>
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
                        <InputField type='email' name='email' label='Почта или номер телефона' placeholder='hurrles@example.com'/>
                        <InputField type='password' name='password' label='Пароль' placeholder='password'/>
                        { this.state.signup &&
                            <InputField type='password' name='repeatPassword' label='Повторите пароль' placeholder='repeat password'/>
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
