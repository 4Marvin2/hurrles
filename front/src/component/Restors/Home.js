import React from 'react'
import Restors from './Restors'
import SearchBar from  '../Restors/SearchBar'
import RestorOpen from '../RestorOpen/RestorOpen'
import '../../css/Home.css'

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

        const restor = {id: 1, srcImg: './imgs/restors/nagoya.webp', title: 'NAGOYA', rating: '4.1', 'metro': 'Бауманская'}

        return (
            <div className='home'>
                <div className='home__restors'>
                    <SearchBar />
                    <Restors restors={restors} />
                </div>
                <div className='home__restorOpen'>
                    <div className='home__restorOpen_in'>
                        <RestorOpen />
                    </div>
                </div>
            </div>
        );
    }
}
