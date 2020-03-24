import axios from 'axios'
import config from "../../config";

class Http {
    constructor() {
        this.backendDomain = config.backendDomain
        this.token = localStorage.getItem('token') ? localStorage.getItem('token').toString() : ''
        this.headers = {
            'Content-Type': 'application/json',
            "Access-Control-Allow-Origin": this.backendDomain,
            "Access-Control-Allow-Headers": "Content-Type, Content-Type,access-control-allow-origin, access-control-allow-headers",
            'token': this.token
        }
    }

    get(url) {
        return axios({
            method: 'GET',
            headers: this.headers,
            url: this.backendDomain + url
        }).then(response => {
            return response.data
        });
    }

    post(url, data) {
        return axios({
            method: 'POST',
            headers: this.headers,
            url: this.backendDomain + url,
            data: data,
        }).then(response => {
            return response.data
        });
    }
}

export default Http