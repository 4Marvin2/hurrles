import { adminURL } from '../constants/urls.js';
import http from '../utils/http.js';

const addRestorRequest = (title, desc, kitchen, address, metro, phoneNumber, openTime, closeTime) => {
    const body = JSON.stringify({
        title: title,
        description: desc,
        kitchen: kitchen,
        address: address,
        metro: metro,
        number: phoneNumber,
        openTime: openTime,
        closeTime: closeTime,
    });

    return http
        .post(adminURL+'/restaurant', body)
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

const updateRestorRequest = (id, title, desc, kitchen, address, metro, phoneNumber, floors, openTime, closeTime) => {
    const body = JSON.stringify({
        title: title,
        description: desc,
        kitchen: kitchen,
        address: address,
        metro: metro,
        number: phoneNumber,
        floors: floors,
        openTime: openTime,
        closeTime: closeTime,
    });

    return http
        .put(adminURL+`/restaurant/${id}`, body)
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

const addDishRequest = (restorId, title, desc, price) => {
    const body = JSON.stringify({
        restaurantId: restorId,
        title: title,
        description: desc,
        price: price,
    });

    return http
        .post(adminURL+'/dish', body)
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

const updateDishRequest = (dishId, restorId, title, desc, price) => {
    const body = JSON.stringify({
        restaurantId: restorId,
        title: title,
        description: desc,
        price: price,
    });

    return http
        .put(adminURL+`/dish/${dishId}`, body)
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

export { addRestorRequest, updateRestorRequest, addDishRequest, updateDishRequest };