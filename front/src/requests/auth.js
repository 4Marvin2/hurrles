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
            return result
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
            return result
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
            return result
        })
        .catch((error) => {
            throw error
        })
};

const logoutRequest = () => {
    return http
        .delete(userURL+'logout')
        .then((result) => {
            if (result.status !== 200) {
                throw 'Non 200 response';
            }
            return result
        })
        .catch((error) => {
            throw error
        })
};

// const editProfileRequest = (name, gender, prefer, date, description, photoPaths, tags) => {
//     const body = JSON.stringify({
//         name: name,
//         gender: gender,
//         prefer: prefer,
//         date: date,
//         description: description,
//         imgs: photoPaths,
//         tags: tags,
//     });

//     return http
//         .put(profileURL, body)
//         .then((response) => {
//             if (response.status !== HTTPSuccess) {
//                 throw 'server internal error';
//             }

//             return response.data;
//         })
//         .catch((err) => {
//             errorManager.pushAPIError();
//             throw err;
//         });
// };

export { signupRequest, loginRequest, getUserRequest, logoutRequest };