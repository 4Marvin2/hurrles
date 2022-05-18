import React from 'react'
import Restors from './Restors'
import SearchBar from  './SearchBar'
import RestorOpen from './RestorOpen/RestorOpen'
import '../css/Home.css'

import { getRestors } from '../requests/restors';

export default class Home extends React.Component {
    constructor(props) {
        super(props);
        let lol = [];
        getRestors().then((data) => {
            const restors = [];
            data.forEach(e => {
                const restorElement = {
                    id: e.id, 
                    srcImg: e.img ? e.img : '', 
                    title: e.title, 
                    rating: '4.1', 
                    metro: e.metro,
                    restorInfo: {
                        rating: '4.8',
                        address: e.address,
                        workTime: `Понедельник-суббота ${e.openTime} - ${e.closeTime}`,
                        desc: e.description,
                        tag1: e.kitchen
                    },
                    dishes: [
                        {id: 1, title: 'Сет Карри Райсу', desc: 'Мисо-суп, рис, овощное Карри, капуста', price: '340 p.'},
                        {id: 2, title: 'Сет Карри Райсу', desc: 'Мисо-суп, рис, овощное Карри, капуста', price: '340 p.'},
                        {id: 3, title: 'Сет Карри Райсу', desc: 'Мисо-суп, рис, овощное Карри, капуста', price: '340 p.'},
                    ]
                };
                restors.push(restorElement);
            });
            lol = restors;
            return restors
        }); 
        console.log(lol)
        this.state = {
            currentID: 0,
            restors: [ {id: 1, srcImg: '../imgs/restors/nagoya.webp', title: 'NAGOYA', rating: '4.1', 'metro': 'Бауманская',
                restorInfo: {rating: '4.8',
                address: 'м.Буманская, Старокирочный переулок 16/2 стр.1',
                workTime: 'Понедельник-суббота 11:00-22:00, Воскресенье 12:00-22:00',
                desc: 'Далеко-далеко за словесными горами в стране гласных и согласных живут рыбные тексты.Вдали от всех живут они в буквенных домах на берегу Семантика большого языкового океана.Маленький ручеек Даль журчит по всей стране и обеспечивает ее всеми необходимыми правилами.',
                tag1: 'Japan'},
                dishes: [
                    {id: 1, title: 'Сет Карри Райсу', desc: 'Мисо-суп, рис, овощное Карри, капуста', price: '340 p.'},
                    {id: 2, title: 'Сет Карри Райсу', desc: 'Мисо-суп, рис, овощное Карри, капуста', price: '340 p.'},
                    {id: 3, title: 'Сет Карри Райсу', desc: 'Мисо-суп, рис, овощное Карри, капуста', price: '340 p.'},
                ]},
                {id: 2, srcImg: '../imgs/restors/nagoya.webp', title: 'LOLCLUB', rating: '5', 'metro': 'Динамо',
                restorInfo: {rating: '4.8',
                address: 'м.Динамо, Старокирочный переулок 16/2 стр.1',
                workTime: 'Понедельник-суббота 11:00-22:00, Воскресенье 12:00-22:00',
                desc: 'Далеко-далеко за словесными горами в стране гласных и согласных живут рыбные тексты.Вдали от всех живут они в буквенных домах на берегу Семантика большого языкового океана.Маленький ручеек Даль журчит по всей стране и обеспечивает ее всеми необходимыми правилами.',
                tag1: 'Japan'},
                dishes: [
                    {id: 1, title: 'Сет Карри Райсу', desc: 'Мисо-суп, рис, овощное Карри, капуста', price: '340 p.'},
                ]},
                {id: 3, srcImg: '../imgs/restors/nagoya.webp', title: 'FUCKCLUB', rating: '2', 'metro': 'Котельники',
                restorInfo: {rating: '4.8',
                address: 'м.Котельники, Старокирочный переулок 16/2 стр.1',
                workTime: 'Понедельник-суббота 11:00-22:00, Воскресенье 12:00-22:00',
                desc: 'Далеко-далеко за словесными горами в стране гласных и согласных живут рыбные тексты.Вдали от всех живут они в буквенных домах на берегу Семантика большого языкового океана.Маленький ручеек Даль журчит по всей стране и обеспечивает ее всеми необходимыми правилами.',
                tag1: 'Russia'},
                dishes: []},
              ],
        };

        this.restorClick = this.restorClick.bind(this);
    }

    restorClick(payload) {
        this.setState({currentID: payload});
    }

    render(){
        return (
            <div className='home'>
                <div className='home__left'>
                    <SearchBar />
                    <Restors restors={this.state.restors} restorClick={this.restorClick} />
                </div>
                <div className='home__right'>
                    {this.state.restors.length !== 0 && 
                        <div className='home__restorOpen_in'>
                            <RestorOpen
                            restorInfo={this.state.restors[this.state.currentID].restorInfo}
                            dishes={this.state.restors[this.state.currentID].dishes} />
                        </div>
                    }
                </div>
            </div>
        );
    }
}
