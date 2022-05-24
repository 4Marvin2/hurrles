import { ordersURL } from '../constants/urls.js';
import http from '../utils/http.js';

const getOrders = () => {
    return http
        .get(ordersURL)
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
        .post(`${ordersURL}/${orderId}/dish/${dishId}`, null)
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
        .delete(`${ordersURL}/${orderId}/dish/${dishId}`, null)
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
        .delete(`${ordersURL}/${orderId}`, null)
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

export { getOrders, addDish, deleteDish, deleteOrder }
