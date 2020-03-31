import axios from 'axios'
import config from "../../config";

class Http {
    constructor() {
        this.backendDomain = config.backendDomain
        this.token = localStorage.getItem('token') ? localStorage.getItem('token').toString() : ''
        this.headers = {
            'Content-Type': 'application/json',
        }
    }

    get(url) {
        return axios({
            method: 'GET',
            headers: this.headers,
            url: this.backendDomain + url,
            params: {token: this.token}
        }).then(response => {
            return response.data.Data
        }).catch(error => {
            console.log(error)
        });
    }

    post(url, data) {
        data.token = this.token
        return axios.post(
            this.backendDomain + url,
            data,
            {headers: this.headers}
            ).then(response => {
                return response.data.Data
            }).catch(error => {
                console.log(error)
            });
    }
}

export default Http