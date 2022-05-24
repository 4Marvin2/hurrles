import React from 'react'
import InputField from '../Common/InputField'
import Button from '../Common/Button';

import '../../css/AdminCommon/AddRestorForm.css'

import validator from 'validator'
import moment from 'moment'

import { addressErrorMsg, addRestorFormErrorMsg, closeTimeErrorMsg, descErrorMsg, kitchenErrorMsg, metroErrorMsg, openTimeErrorMsg, phoneNumberErrorMsg, titleErrorMsg } from '../../constants/errorMsg';
import { addRestorRequest, updateRestorRequest } from '../../requests/admin';
import { metroMap } from '../../utils/metroMap';

export default class AddRestorForm extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            titleIsValid: true,
            descIsValid: true,
            kitchenIsValid: true,
            addressIsValid: true,
            metroIsValid: true,
            phoneNumberIsValid: true,
            openTimeIsValid: true,
            closeTimeIsValid: true,
            formErrorIsActive: false,
            restor: props.restor
        }
        if (props.restor) {
            this.state['restor'] = props.restor
        } else {
            this.state['restor'] = {
                id: 0,
                title: '',
                desc: '',
                kitchen: '',
                address: '',
                metro: '',
                phoneNumber: '',
                openTime: '',
                closeTime: '',
            }
        }
        if (props.isUpdate) {
            this.state['isUpdate'] = props.isUpdate
        } else {
            this.state['isUpdate'] = false
        }
    }

    updateRestorsCallback(payload) {
        this.props.updateRestorsCallback(payload)
    }

    closeAllForms(payload) {
        this.props.closeAllForms(payload)
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

    kitchenOnInput = (payload) => {
        this.setState({kitchen: payload.target.value})
        if (payload.target.value.length !== 0) {
            this.setState({kitchenIsValid: true})
        } else {
            this.setState({kitchenIsValid: false})
        }
    }

    addressOnInput = (payload) => {
        this.setState({address: payload.target.value});
        if (payload.target.value.length !== 0) {
            this.setState({addressIsValid: true});
        } else {
            this.setState({addressIsValid: false});
        }
    }

    metroOnInput = (payload) => {
        this.setState({metro: payload.target.value.toLowerCase()});
        if (payload.target.value.length !== 0 && metroMap[payload.target.value.toLowerCase()]) {
            this.setState({metroIsValid: true});
        } else {
            this.setState({metroIsValid: false});
        }
    }

    phoneNumberOnInput = (payload) => {
        this.setState({phoneNumber: payload.target.value})
        if (validator.isMobilePhone(payload.target.value, 'ru-RU') && payload.target.value.length !== 0) {
            this.setState({phoneNumberIsValid: true})
        } else {
            this.setState({phoneNumberIsValid: false})
        }
    }

    openTimeOnInput = (payload) => {
        const fullDate = new Date();
        const d = moment(fullDate).format('L'); // d = "12/12/2017"
        const date = moment(d +' '+ payload.target.value).format();
        this.setState({openTime: date})
        console.log(payload.target.value)
        if (payload.target.value.length !== 0) {
            this.setState({openTimeIsValid: true})
        } else {
            this.setState({openTimeIsValid: false})
        }
    }

    closeTimeOnInput = (payload) => {
        const fullDate = new Date();
        const d = moment(fullDate).format('L'); // d = "12/12/2017"
        const date = moment(d +' '+ payload.target.value).format();
        this.setState({closeTime: date})
        if (payload.target.value.length !== 0) {
            this.setState({closeTimeIsValid: true})
        } else {
            this.setState({closeTimeIsValid: false})
        }
    }

    setDefaultRestorFields = () => {
        let state = true
        if (!this.state.title && this.state.restor.title !== '' && this.state.restor.title !== undefined) {
            this.state.title = this.state.restor.title
        } else if (!this.state.title && (this.state.restor.title === '' || this.state.restor.title === undefined)) {
            this.setState({titleIsValid: false})
            state = false
        }
        if (!this.state.desc && this.state.restor.desc !== '' && this.state.restor.desc !== undefined) {
            this.state.desc = this.state.restor.desc
        } else if (!this.state.desc && (this.state.restor.desc === '' || this.state.restor.desc === undefined)) {
            this.setState({descIsValid: false})
            state = false
        }
        if (!this.state.kitchen && this.state.restor.kitchen !== '' && this.state.restor.kitchen !== undefined) {
            this.state.kitchen = this.state.restor.kitchen
        } else if (!this.state.kitchen && (this.state.restor.kitchen === '' || this.state.restor.kitchen === undefined)) {
            this.setState({kitchenIsValid: false})
            state = false
        }
        if (!this.state.address && this.state.restor.address !== '' && this.state.restor.address !== undefined) {
            this.state.address = this.state.restor.address
        } else if (!this.state.address && (this.state.restor.address === '' || this.state.restor.address === undefined)) {
            this.setState({addressIsValid: false})
            state = false
        }
        if (!this.state.metro && this.state.restor.metro !== '' && this.state.restor.metro !== undefined) {
            this.state.metro = this.state.restor.metro
        } else if (!this.state.metro && (this.state.restor.metro === '' || this.state.restor.metro === undefined)) {
            this.setState({metroIsValid: false})
            state = false
        }
        if (!this.state.phoneNumber && this.state.restor.phoneNumber !== '' && this.state.restor.phoneNumber !== undefined) {
            this.state.phoneNumber = this.state.restor.phoneNumber
        } else if (!this.state.desc && (this.state.restor.phoneNumber === '' || this.state.restor.phoneNumber === undefined)) {
            this.setState({phoneNumberIsValid: false})
            state = false
        }
        if (!this.state.openTime && this.state.restor.openTime !== '' && this.state.restor.openTime !== undefined) {
            console.log(this.state.restor.openTime)
            this.state.openTime = this.state.restor.openTime
        } else if (!this.state.openTime && (this.state.restor.openTime === '' || this.state.restor.openTime === undefined)) {
            this.setState({openTimeIsValid: false})
            state = false
        }
        if (!this.state.closeTime && this.state.restor.closeTime !== '' && this.state.restor.closeTime !== undefined) {
            this.state.closeTime = this.state.restor.closeTime
        } else if (!this.state.closeTime && (this.state.restor.closeTime === '' || this.state.restor.closeTime === undefined)) {
            this.setState({closeTimeIsValid: false})
            state = false
        }
        return state
    }

    formButtonClick = () => {
        if (this.state.titleIsValid &&
            this.state.descIsValid &&
            this.state.kitchenIsValid &&
            this.state.addressIsValid &&
            this.state.metroIsValid &&
            this.state.phoneNumberIsValid &&
            this.state.openTimeIsValid &&
            this.state.closeTimeIsValid) {
                if (this.state.isUpdate) {
                    const ok = this.setDefaultRestorFields()
                    if (!ok) {
                        return
                    }
                    updateRestorRequest(
                        this.state.restor.id,
                        this.state.title,
                        this.state.desc,
                        this.state.kitchen,
                        this.state.address,
                        this.state.metro,
                        this.state.phoneNumber,
                        this.state.openTime,
                        this.state.closeTime
                    )
                    .then((data) => {
                        this.setState({formErrorIsActive: false})
                        this.closeAllForms()
                        this.updateRestorsCallback()
                    })
                    .catch((error) => {
                        this.setState({formErrorIsActive: true})
                        console.log(error)
                    })
                } else {
                    addRestorRequest(
                        this.state.title,
                        this.state.desc,
                        this.state.kitchen,
                        this.state.address,
                        this.state.metro,
                        this.state.phoneNumber,
                        this.state.openTime,
                        this.state.closeTime
                    )
                    .then((data) => {
                        this.setState({formErrorIsActive: false})
                        this.closeAllForms()
                        this.updateRestorsCallback()
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
            <div className='add-restor-form'>
                {/* exit button */}
                <div className='add-restor-form__fields'>
                        <InputField
                            type='text'
                            name='title'
                            label='Название'
                            value={this.state.restor.title}
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
                            value={this.state.restor.desc}
                            rows={10}
                            cols={100}
                            maxLength={1000}
                            prompt={descErrorMsg}
                            promptIsActive={!this.state.descIsValid}
                            onInput={(e) => {this.descOnInput(e)}}
                            theme='dark'
                        />
                        <InputField
                            type='text'
                            name='kitchen'
                            label='Кухня'
                            value={this.state.restor.kitchen}
                            prompt={kitchenErrorMsg}
                            promptIsActive={!this.state.kitchenIsValid}
                            onInput={(e) => {this.kitchenOnInput(e)}}
                            theme='dark'
                        />
                        <InputField
                            type='text'
                            name='address'
                            label='Адрес'
                            value={this.state.restor.address}
                            maxLength={255}
                            prompt={addressErrorMsg}
                            promptIsActive={!this.state.addressIsValid}
                            onInput={(e) => {this.addressOnInput(e)}}
                            theme='dark'
                        />
                        <InputField
                            type='text'
                            name='metro'
                            label='Ближайшее метро'
                            value={this.state.restor.metro}
                            maxLength={50}
                            prompt={metroErrorMsg}
                            promptIsActive={!this.state.metroIsValid}
                            onInput={(e) => {this.metroOnInput(e)}}
                            theme='dark'
                        />
                        <InputField
                            type='tel'
                            name='phoneNumber'
                            label='Номер телефона'
                            value={this.state.restor.phoneNumber}
                            maxLength={16}
                            prompt={phoneNumberErrorMsg}
                            promptIsActive={!this.state.phoneNumberIsValid}
                            onInput={(e) => {this.phoneNumberOnInput(e)}}
                            theme='dark'
                        />
                        <InputField
                            type='time'
                            name='openTime'
                            label='Время открытия'
                            value={this.state.restor.openTime}
                            prompt={openTimeErrorMsg}
                            promptIsActive={!this.state.openTimeIsValid}
                            onInput={(e) => {this.openTimeOnInput(e)}}
                            theme='dark'
                        />
                        <InputField
                            type='time'
                            name='closeTime'
                            label='Время закрытия'
                            value={this.state.restor.closeTime}
                            prompt={closeTimeErrorMsg}
                            promptIsActive={!this.state.closeTimeIsValid}
                            onInput={(e) => {this.closeTimeOnInput(e)}}
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
