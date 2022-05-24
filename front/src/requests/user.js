import { userURL } from '../constants/urls.js';
import http from '../utils/http.js';

const signupRequest = (email, password) => {
    const body = JSON.stringify({
        email: email,
        password: password,
    });

    return http
        .post(userURL+'/signup', body)
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

const loginRequest = (email, password) => {
    const body = JSON.stringify({
        email: email,
        password: password,
    });

    return http
        .post(userURL+'/login', body)
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

const getUserRequest = () => {
    return http
        .get(userURL)
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

const logoutRequest = () => {
    return http
        .get(userURL+'/logout')
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

const editRequest = (name, email, phoneNumber) => {
    const body = JSON.stringify({
        fullName: name,
        email: email,
        number: phoneNumber,
    });

    return http
        .put(userURL+'/edit', body)
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

export { signupRequest, loginRequest, getUserRequest, logoutRequest, editRequest };