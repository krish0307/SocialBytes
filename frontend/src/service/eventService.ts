import axios from "axios";
import { IFormInput } from "../pages/CreateEvent";

const baseUrl = 'http://localhost:9010/api';

export const eventService = {
    getAll,
    getById,
    create,
    update,
    searchOp,
    getUser,
    getAttendees,
    delete: _delete
};

function getAll() {
    return get(`${baseUrl}/getEvents`);
}

function getById(id: any):Promise<IFormInput> {
    return get(`${baseUrl}/getEvent?id=${id}`);
}

function getAttendees(id: any):Promise<Array<string>> {
    return get(`${baseUrl}/getAttendees?id=${id}`);
}

function searchOp(name:any){
    return get(`${baseUrl}/searchEvent?name=${name}`);
}

function create(params: any) {
    const requestOptions = {
        method: 'POST',
        // headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(params)
    };
    return axios.post(`${baseUrl}/createEvent`, JSON.stringify(params), {withCredentials: true}).then(response=>{
        if (response.status!= 200) {
            // const error = (data && data.message) || response.statusText;
            return Promise.reject(response.statusText);
        }
        return (response.status);
    });
}

function update(id: any, params: any) {
    const requestOptions = {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(params)
    };
    return fetch(`${baseUrl}/${id}`, requestOptions).then(handleResponse);
}

function _delete(id: any) {
    const requestOptions = {
        method: 'DELETE'
    };
    return fetch(`${baseUrl}/deleteEvent?id=${id}`, requestOptions).then(handleResponse);
}

function get(url: any) {
    const requestOptions = {
        method: 'GET'
    };
    return fetch(url, requestOptions).then(handleResponse);
}

function getUser(url: any) {
    const requestOptions = {
        method: 'GET',
        withCredentials: true
    };
    return axios.get(url, {withCredentials: true}).then(handleResponse);
}


function handleResponse(response:any) {
    return response.text().then((text: string) => {
        const data = text && JSON.parse(text);

        if (!response.ok) {
            const error = (data && data.message) || response.statusText;
            return Promise.reject(error);
        }

        return data;
    });
}