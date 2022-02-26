import axios, { AxiosRequestConfig, AxiosResponse} from 'axios';
import { signInURL, usersURL } from './apiPaths';

type SignInData = {
    username : string;
    password : string;
}

export const signInAPI = async (obj : SignInData) => {
    const config : AxiosRequestConfig = {
        method: 'post',
        url: signInURL,
        data: obj,
        headers: {
            "Content-Type": "application/json",
        },
    };
      const response :AxiosResponse = await axios(config);
      console.log(response);
}

