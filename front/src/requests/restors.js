import { restaurantsURL, restaurantURL } from '../constants/urls.js';
import http from '../utils/http.js';

const getRestors = () => {
    return http
        .get(restaurantsURL)
        .then(
            (result) => {
                console.log(result)
                return result
            },
            (error) => {
                console.log(error)
            }
        )
};

const getFavoriteRestors = () => {
    return http
        .get(`${restaurantURL}/favorite`)
        .then(
            (result) => {
                console.log(result)
                return result
            },
            (error) => {
                console.log(error)
            }
        )
};

const getRestorMenu = (id) => {
    return http
        .get(`${restaurantURL}/${id}/menu`)
        .then(
            (result) => {
                console.log(result)
                return result
            },
            (error) => {
                console.log(error)
            }
        )
};

export { getRestors, getRestorMenu, getFavoriteRestors };