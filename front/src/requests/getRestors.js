import { restaurantsURL } from '../constants/urls.js';
import http from '../utils/http.js';

const getRestors = () => {
    return http
        .get(restaurantsURL)
        .then(res => res.json())
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

export { getRestors };