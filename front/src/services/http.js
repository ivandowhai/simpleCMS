import axios from 'axios'

class Http {
    constructor() {
        this.backendDomain = 'http://127.0.0.1:8080/'
    }

    get(url) {
        return axios({
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
            },
            url: this.backendDomain + url
        }).then(response => {
            return response.data
        });
    }
}