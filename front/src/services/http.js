import axios from 'axios'
import config from "../../config";

class Http {
    constructor() {
        this.backendDomain = config.backendDomain
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

export default Http