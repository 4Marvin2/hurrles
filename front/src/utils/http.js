import { serverAddress } from '../constants/urls.js';

/** Class representing a http request. */
class Http {
    /** Базовый URL. */
    _baseURL = serverAddress;

    set baseURL(baseURL) {
        this._baseURL = baseURL;
    }

    /**
     * Делает запрос
     * @param {string} url - Запрос по url.
     * @param {string} method - Метод запроса
     * @param {Headers} headers - Заголовки запроса
     * @param {string} body - Тело запроса
     * @return {Promise<{status: number, data: any}>} - Возвращает Promise со статусом ответа и данными.
     */
    async _request({ url = '/', method = 'GET', headers = new Headers(), body = null }) {
        const csrfToken = sessionStorage.getItem('csrf');
        if (csrfToken) {
            const csrfHeader = new Headers();
            csrfHeader.set('x-csrf-Token', csrfToken);
            headers = csrfHeader;
        }

        const response = await fetch(this._baseURL + url, {
            method: method,
            headers: headers,
            body: body,
            credentials: 'include',
        });

        const responseData = await response.json();

        const csrfTokenResp = response.headers.get('csrf');
        if (csrfTokenResp) {
            sessionStorage.setItem('csrf', csrfTokenResp);
        }

        return responseData;
    }

    /**
     * Делает GET запрос
     * @param {string} url - Запрос по url.
     * @return {Promise<{status: number, data: any}>} - Возвращает Promise со статусом ответа и данными.
     */
    get(url) {
        return this._request({ url: url });
    }

    /**
     * Делает POST запрос
     * @param {string} url - Запрос по url.
     * @param {Object} body - Данные для отправки.
     * @return {Promise<{status: number, data: any}>} - Возвращает Promise со статусом ответа и данными.
     */
    post(url, body) {
        return this._request({
            url: url,
            method: 'POST',
            body: body,
        });
    }

    /**
     * Делает DELETE запрос
     * @param {string} url - Запрос по url.
     * @param {Object} body - Данные для отправки.
     * @return {Promise<{status: number, data: any}>} - Возвращает Promise со статусом ответа и данными.
     */
    delete(url, body = null) {
        return this._request({
            url: url,
            method: 'DELETE',
            body: body,
        });
    }

    /**
     * Делает PUT запрос
     * @param {string} url - Запрос по url.
     * @param {Object} body - Данные для отправки.
     * @return {Promise<{status: number, body: any}>} - Возвращает Promise со статусом ответа и данными.
     */
    put(url, body) {
        return this._request({
            url: url,
            method: 'PUT',
            body: body,
        });
    }
}

export default new Http();