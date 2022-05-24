import { ordersURL, orders } from '../constants/urls.js';
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
        .post(`${orders}`, body)
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

export { getOrders, createOrder }
