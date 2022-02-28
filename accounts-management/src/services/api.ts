import axios, { AxiosRequestConfig, AxiosResponse} from 'axios';
import { signInURL, usersURL } from './apiPaths';

interface SignInData {
    username : string;
    password : string;
} 

interface createUserData {
    username: string,
    password: string,
    fullName: string,
    idNumber: string,
    role: string
}

interface updateUserData {
    username?: string,
    password?: string,
    fullName?: string,
    idNumber?: string,
    role?: string
}

const accessToken :string = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDYyMTc2OTMsImp0aSI6ImQwZjZkNTM0LWVlYTUtNGE4OS1iYzBmLTU0ZWRjMmNlZjRhMSIsImlhdCI6MTY0NTYxMjg5Mywic3ViIjoiNWU3ZjgyZWEtOTllYS00NDQwLWFmNTItYzQ4NTRjZGMwM2EwIn0.5J8V7d3MPxcXj0AcOEBh7K9wnVy4_4fVfHpLKuqQCUk"


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

export const getUsersAPI = async (id :string) => {
    const url = (id || '').length > 0 ? usersURL + `/${id}` : usersURL;
    console.log(url)
    const config : AxiosRequestConfig = {
        method: 'get',
        url: url,
        headers: {
            "Content-Type": "application/json",
            "Authorization": accessToken
        },
    };
    let response :AxiosResponse;
    try {
        response = await axios(config);
    } catch (error) {
        response = error.response;
    }
    return response;
}

export const createUsersAPI = async (data :createUserData) => {
    const config : AxiosRequestConfig = {
        method: 'post',
        url: usersURL,
        headers: {
            "Content-Type": "application/json",
            "Authorization": accessToken
        },
        data: [data]
    };
    let response :AxiosResponse;
    try {
        response = await axios(config);
    } catch (error) {
        response = error.response;
    }
    return response;
}

export const updateUserAPI = async (data :updateUserData, id :string) => {
    const config : AxiosRequestConfig = {
        method: 'patch',
        url: usersURL + `/${id}`,
        headers: {
            "Content-Type": "application/json",
            "Authorization": accessToken
        },
        data: data
    };
    let response :AxiosResponse;
    try {
        response = await axios(config);
    } catch (error) {
        response = error.response;
    }
    return response;
}

export const deleteUserAPI = async (id :string) => {
    const config : AxiosRequestConfig = {
        method: 'delete',
        url: usersURL + `/${id}`,
        headers: {
            "Content-Type": "application/json",
            "Authorization": accessToken
        },
    };
    let response :AxiosResponse;
    try {
        response = await axios(config);
    } catch (error) {
        response = error.response;
    }
    return response;
}
