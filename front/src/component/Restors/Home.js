import React from 'react'
// import Restors from './Restors'
// import NavBar from '../../NavBar'
// import '../../css/Restors/Home.css'
// import "../../css/App.css";

export default class Home extends React.Component {
    render(){
        const restors = [
            {id: 1, srcImg: './imgs/restors/nagoya.webp', title: 'NAGOYA', rating: '4.1', 'metro': 'Бауманская'},
            {id: 2, srcImg: './imgs/restors/nagoya.webp', title: 'RamenClub', rating: '4.8', 'metro': 'Бауманская'},
            {id: 3, srcImg: './imgs/restors/nagoya.webp', title: 'RamenClub', rating: '4.8', 'metro': 'Бауманская'},
            {id: 4, srcImg: './imgs/restors/nagoya.webp', title: 'RamenClub', rating: '4.8', 'metro': 'Бауманская'},
            {id: 5, srcImg: './imgs/restors/nagoya.webp', title: 'RamenClub', rating: '4.8', 'metro': 'Бауманская'},
            {id: 6, srcImg: './imgs/restors/nagoya.webp', title: 'RamenClub', rating: '4.8', 'metro': 'Бауманская'},
            {id: 7, srcImg: './imgs/restors/nagoya.webp', title: 'RamenClub', rating: '4.8', 'metro': 'Бауманская'},
            {id: 8, srcImg: './imgs/restors/nagoya.webp', title: 'RamenClub', rating: '4.8', 'metro': 'Бауманская'},
            {id: 9, srcImg: './imgs/restors/nagoya.webp', title: 'RamenClub', rating: '4.8', 'metro': 'Бауманская'},
            {id: 10, srcImg: './imgs/restors/nagoya.webp', title: 'RamenClub', rating: '4.8', 'metro': 'Бауманская'},
            {id: 11, srcImg: './imgs/restors/nagoya.webp', title: 'RamenClub', rating: '4.8', 'metro': 'Бауманская'},
            {id: 12, srcImg: './imgs/restors/nagoya.webp', title: 'RamenClub', rating: '4.8', 'metro': 'Бауманская'},
            {id: 13, srcImg: './imgs/restors/nagoya.webp', title: 'RamenClub', rating: '4.8', 'metro': 'Бауманская'},
            {id: 14, srcImg: './imgs/restors/nagoya.webp', title: 'RamenClub', rating: '4.8', 'metro': 'Бауманская'},
          ];

        return (
            // <div className='app__home home'>
            //     <NavBar />
            //     <Restors restors={restors} />
            // </div>
            <div>
                <img src='https://www.google.com/url?sa=i&url=https%3A%2F%2Fgithub.com%2Folehan%2Fkek&psig=AOvVaw3zhgdZ0-get2pdxJVkJFVJ&ust=1652626757054000&source=images&cd=vfe&ved=0CAwQjRxqFwoTCMj3r8mg3_cCFQAAAAAdAAAAABAD' alt='kek' />
            </div>
        );
    }
}
