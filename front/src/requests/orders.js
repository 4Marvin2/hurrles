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

export { getOrders }
