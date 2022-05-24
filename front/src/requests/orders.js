import { ordersURL, orderURL } from '../constants/urls.js';
import http from '../utils/http.js';

const getOrders = () => {
    return http
        .get(orderURL)
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

const getRestaurantOrders = () => {
    return http
        .get(`${orderURL}/restaurant`)
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

const createOrder = (userId, placeId, startTime) => {
    const str = startTime;
    const str1 = str.substring(0, 11);
    const str2 = str.substring(11, 13);
    const str3 = str.substring(13);
    const endHours = parseInt(str2, 10)+2
    const endTime = `${str1}${endHours}${str3}`
    const body = JSON.stringify({
        userId: userId,
        placeId: placeId,
        startTime: startTime,
        endTime: endTime,
        cost: 0
    });

    return http
        .post(`${ordersURL}`, body)
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

const addDish = (orderId, dishId) => {
    return http
        .post(`${orderURL}/${orderId}/dish/${dishId}`, null)
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

const deleteDish = (orderId, dishId) => {
    return http
        .delete(`${orderURL}/${orderId}/dish/${dishId}`, null)
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

const deleteOrder = (orderId) => {
    return http
        .delete(`${orderURL}/${orderId}`, null)
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

export { getOrders, addDish, deleteDish, deleteOrder, createOrder, getRestaurantOrders }
