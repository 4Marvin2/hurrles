import { adminPlaceURL, adminPlacesURL } from '../constants/urls.js';
import http from '../utils/http.js';

const createPlace = (restaurantId, capacity, number, left, top, width, height, floor) => {
    const body = JSON.stringify({
        restaurantId: restaurantId,
        capacity: capacity,
        number: number,
        leftTop: left,
        rightBottom: top,
        width: width,
        height: height,
        floor: floor
    });

    return http
        .post(`${adminPlaceURL}`, body)
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

const updatePlaces = (places) => {
    const body = JSON.stringify(places);

    return http
        .put(`${adminPlacesURL}`, body)
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

const deletePlace = (placeId) => {
    return http
        .delete(`${adminPlaceURL}/${placeId}`, null)
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

export { createPlace, deletePlace, updatePlaces };