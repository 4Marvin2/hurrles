import React from 'react'
import { profileFormErrorMsg } from '../constants/errorMsg';
import { editRequest } from '../requests/user';
import '../css/Profile.css'
import '../css/RestorAdmin.css'
import RestorAdminCart from './RestorAdminCart';
import Exit from '../imgs/tapbar/exit.svg';

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

    logoutClick() {
        this.props.logoutClick()
    }

    render(){
        return (
            <div className='restor-admin'>
                <RestorAdminCart />
                <button className="restor-admin__exit-button">
                    <img src={Exit} alt="Exit" onClick={() => {this.logoutClick()}} />
                </button> 
            </div>
        );
    }
}
