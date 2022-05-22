import React from 'react'
import { nameErorrMsg, emailErrorMsg, phoneNumberErrorMsg, profileFormErrorMsg } from '../constants/errorMsg';
import InputField from './Common/InputField';
import Button from './Common/Button';
import validator from 'validator'
import { editRequest } from '../requests/user';
import '../css/Profile.css'

export default class RestorAdmin extends React.Component {
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
            <div className='restor-admin'>
                RESTOR ADMIN
            </div>
        );
    }
}
