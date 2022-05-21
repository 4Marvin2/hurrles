import React from 'react'
import { nameErorrMsg, phoneNumberErrorMsg, profileFormErrorMsg } from '../constants/errorMsg';
import InputField from './Common/InputField';

var BreakException = {};

export default class Profile extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            name: props.name,
            isFavorite: false,
            restors:[],
        }
    }

    render(){
        return (
            <div className='profile'>
                 {/* <InputField
                    type='text'
                    name='username'
                    label='Имя'
                    placeholder='Ваше имя'
                    value={this.state.name}
                    prompt={emailErrorMsg}
                    promptIsActive={!this.state.emailIsValid}
                    onInput={(e) => {this.emailOnInput(e)}}
                /> */}
            </div>
        );
    }
}