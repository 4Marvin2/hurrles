import { restaurantsURL, restaurantURL } from '../constants/urls.js';
import http from '../utils/http.js';

const getRestors = () => {
    return http
        .get(restaurantsURL)
        .then((result) => {
            if (result.status !== 200) {
                throw 'Non 200 response';
            }
            return result.data
        })
        .catch((error) => {
            throw error
        })
};

const getFavoriteRestors = () => {
    return http
        .get(`${restaurantURL}/favorite`)
        .then((result) => {
            if (result.status !== 200) {
                throw 'Non 200 response';
            }
            return result.data
        })
        .catch((error) => {
            throw error
        })
};

const getRestorMenu = (id) => {
    return http
        .get(`${restaurantURL}/${id}/menu`)
        .then((result) => {
            if (result.status !== 200) {
                throw 'Non 200 response';
            }
            return result.data
        })
        .catch((error) => {
            throw error
        })
};

const addFavorite = (id) => {
    return http
        .post(`${restaurantURL}/favorite/${id}`, null)
        .then((result) => {
            if (result.status !== 200) {
                throw 'Non 200 response';
            }
            return result.data
        })
        .catch((error) => {
            throw error
        })
};

const deleteFavorite = (id) => {
    return http
        .delete(`${restaurantURL}/favorite/${id}`, null)
        .then((result) => {
            if (result.status !== 200) {
                throw 'Non 200 response';
            }
            return result.data
        })
        .catch((error) => {
            throw error
        })
};

const getPlaces = (id, time, floor) => {
    console.log(id, time, floor)
    const body = JSON.stringify({
        restaurantId: id,
        time: time,
        floor: floor,
    });

    return http
        .post(`${restaurantURL}/${id}/places`, body)
        .then((result) => {
            if (result.status !== 200) {
                throw 'Non 200 response';
            }
            return result.data
        })
        .catch((error) => {
            throw error
        })
};

export { getRestors, getRestorMenu, getFavoriteRestors, addFavorite, deleteFavorite, getPlaces };