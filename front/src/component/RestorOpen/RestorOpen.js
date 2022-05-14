import React from 'react'
// import Menu from './Menu'
// import RestorNavBar from '../NavBar'
// import Info from '../RestorOpen/Info'
// import '../../css/RestorOpen/RestorOpen.css'
// import '../../css/RestorOpen/Menu.css'
// import '../../css/RestorOpen/Info.css'
// import '../../css/RestorOpen/Dish.css'
// import "../../css/App.css";

export default class Home extends React.Component {
    render(){
        const restorInfo = {rating: '4.8',
        address: 'м.Буманская, Старокирочный переулок 16/2 стр.1',
        workTime: 'Понедельник-суббота 11:00-22:00, Воскресенье 12:00-22:00',
        desc: 'Далеко-далеко за словесными горами в стране гласных и согласных живут рыбные тексты.Вдали от всех живут они в буквенных домах на берегу Семантика большого языкового океана.Маленький ручеек Даль журчит по всей стране и обеспечивает ее всеми необходимыми правилами.',
        tag1: 'Japan'}

        const dishes = [
            {id: 1, title: 'Сет Карри Райсу', desc: 'Мисо-суп, рис, овощное Карри, капуста', price: '340 p.'},
            {id: 2, title: 'Сет Карри Райсу', desc: 'Мисо-суп, рис, овощное Карри, капуста', price: '340 p.'},
            {id: 3, title: 'Сет Карри Райсу', desc: 'Мисо-суп, рис, овощное Карри, капуста', price: '340 p.'},
            {id: 4, title: 'Сет Карри Райсу', desc: 'Мисо-суп, рис, овощное Карри, капуста', price: '340 p.'},
            {id: 5, title: 'Сет Карри Райсу', desc: 'Мисо-суп, рис, овощное Карри, капуста', price: '340 p.'},
            {id: 6, title: 'Сет Карри Райсу', desc: 'Мисо-суп, рис, овощное Карри, капуста', price: '340 p.'},
            {id: 7, title: 'Сет Карри Райсу', desc: 'Мисо-суп, рис, овощное Карри, капуста', price: '340 p.'},
            {id: 8, title: 'Сет Карри Райсу', desc: 'Мисо-суп, рис, овощное Карри, капуста', price: '340 p.'},
            {id: 9, title: 'Сет Карри Райсу', desc: 'Мисо-суп, рис, овощное Карри, капуста', price: '340 p.'},
            {id: 10, title: 'Сет Карри Райсу', desc: 'Мисо-суп, рис, овощное Карри, капуста', price: '340 p.'},
            {id: 11, title: 'Сет Карри Райсу', desc: 'Мисо-суп, рис, овощное Карри, капуста', price: '340 p.'},
            {id: 12, title: 'Сет Карри Райсу', desc: 'Мисо-суп, рис, овощное Карри, капуста', price: '340 p.'},
            {id: 13, title: 'Сет Карри Райсу', desc: 'Мисо-суп, рис, овощное Карри, капуста', price: '340 p.'},
            {id: 14, title: 'Сет Карри Райсу', desc: 'Мисо-суп, рис, овощное Карри, капуста', price: '340 p.'},
          ];

        return (
            // <div className="app_restor-open restor-open">
            //     <RestorNavBar />
            //     <div className="restor-open__main">
            //         <Info  restorInfo={restorInfo}/>
            //         <Menu dishes={dishes} />
            //     </div>
            // </div>
            <div>
                <img src='https://www.google.com/url?sa=i&url=https%3A%2F%2Fwww.desicomments.com%2Fdesi%2Flol%2F&psig=AOvVaw1iYizGcKZiZYUgDkq7Jtzz&ust=1652626719109000&source=images&cd=vfe&ved=0CAwQjRxqFwoTCND-ybag3_cCFQAAAAAdAAAAABAD' alt='lol' />
            </div>
        );
    }
}
