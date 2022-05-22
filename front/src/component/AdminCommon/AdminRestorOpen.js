import React from 'react'
import Menu from '../RestorOpen/Menu'
import Button from '../Common/Button'
import TitleBar from '../RestorOpen/TitleBar'
import Info from '../RestorOpen/Info'
import AddDishForm from './AddDishForm'

import '../../css/AdminCommon/AdminRestorOpen.css'
import '../../css/RestorOpen/Menu.css'
import '../../css/RestorOpen/Dish.css'

export default class AdminRestorOpen extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            addFormActive: false,
            updateFormActive: false,
            openRestorId: props.restorId,
        };

        this.addDishCallback = this.addDishCallback.bind(this);
        this.addButtonClick = this.addButtonClick.bind(this);
        this.updateButtonClick = this.updateButtonClick.bind(this);
    }

    addButtonClick = () => {
        this.setState({addFormActive: true, updateFormActive: false});
    }

    updateButtonClick = (payload) => {
        this.setState({
            addFormActive: false,
            updateFormActive: true,
            openDishId: payload.id,
            openDish: payload,
        });
    }

    shouldComponentUpdate = (newProps, newState) => {
        if (this.state.addFormActive !== newState.addFormActive && this.state.addFormActive === false) {
            if (this.state.openRestorId !== newProps.restorId) {
                this.setState({openRestorId: newProps.restorId})
            }
            return true
        }
        if (this.state.updateFormActive !== newState.updateFormActive && this.state.updateFormActive === false) {
            if (this.state.openRestorId !== newProps.restorId) {
                this.setState({openRestorId: newProps.restorId})
            }
            return true
        }
        if (this.state.openRestorId !== newProps.restorId && newProps.restorId !== undefined) {
            this.setState({openRestorId: newProps.restorId, addFormActive: false, updateFormActive: false})
            return true
        }

        return false
    }

    addDishCallback(payload) {
        if (payload.formIsActive === false) {
            this.updateRestorsCallback()
        }

        this.setState({addFormActive: payload.formIsActive, updateFormActive: payload.formIsActive});
    }

    updateRestorsCallback(payload) {
        this.props.updateRestorsCallback(payload)
    }

    openUpdateRestorClick(payload) {
        this.props.openUpdateRestorClick();
    }

    render() {
        let tmpl
        if (!this.state.addFormActive && !this.state.updateFormActive) {
            tmpl = (
                <div className="admin-restor-open">
                    <TitleBar isFavorite={this.props.isFavorite} likeClick={this.props.likeClick} />
                    <div className="admin-restor-open__main">
                        <Info restorInfo={this.props.restorInfo}/>
                        <Menu dishes={this.props.dishes} isAdmin={true} restorId={this.state.openRestorId} addButtonClick={this.addButtonClick} updateButtonClick={this.updateButtonClick}/>
                    </div>
                    <div className="admin-restor-open__reverse-button">
                        <Button text='Изменить' theme='dark' width={300} height={72} onClick={(e) => this.openUpdateRestorClick(e)}/>
                    </div>
                </div>
            );
        } else {
            tmpl = (
                <div className='admin-restor-open admin-restor-open_form-open'>
                    { !this.state.addFormActive && this.state.updateFormActive &&
                        <AddDishForm
                            restorId={this.state.openRestorId}
                            addDishCallback={this.addDishCallback}
                            dish={
                                {
                                    id: this.state.openDish.id,
                                    title: this.state.openDish.title,
                                    desc: this.state.openDish.desc,
                                    price: this.state.openDish.price,
                                }
                            }
                            isUpdate={true}
                            closeAllForms={this.props.closeAllForms}
                        />
                    }
                    { this.state.addFormActive && !this.state.updateFormActive &&
                        <AddDishForm
                            restorId={this.state.openRestorId}
                            addDishCallback={this.addDishCallback}
                            closeAllForms={this.props.closeAllForms}
                        />
                    }
                </div>
            );
        }


        return tmpl;
    }
}
