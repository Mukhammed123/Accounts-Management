import axios, { AxiosRequestConfig, AxiosResponse} from 'axios';
import { signInURL, usersURL } from './apiPaths';
import {useStore} from '@/store/useStore';

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


export const signInAPI = async (obj : SignInData) => {
  const config : AxiosRequestConfig = {
    method: 'post',
    url: signInURL,
    data: obj,
    headers: {
      'Content-Type': 'application/json',
    },
  };
  let response :AxiosResponse;
  try {
    response = await axios(config);
  } catch(error) {
    response = error.response;
  }
  return response;
};

export const getUsersAPI = async (id :string) => {
  const store = useStore();
  const url = (id || '').length > 0 ? usersURL + `/${id}` : usersURL;
  const config : AxiosRequestConfig = {
    method: 'get',
    url: url,
    headers: {
      'Content-Type': 'application/json',
      'Authorization': 'Bearer ' + store.accessToken
    },
  };
  let response :AxiosResponse;
  try {
    response = await axios(config);
  } catch (error) {
    response = error.response;
  }
  return response;
};

export const createUsersAPI = async (data :createUserData) => {
  const store = useStore();
  const config : AxiosRequestConfig = {
    method: 'post',
    url: usersURL,
    headers: {
      'Content-Type': 'application/json',
      'Authorization': 'Bearer ' + store.accessToken
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
};

export const updateUserAPI = async (data :updateUserData, id :string) => {
  const store = useStore();
  const config : AxiosRequestConfig = {
    method: 'patch',
    url: usersURL + `/${id}`,
    headers: {
      'Content-Type': 'application/json',
      'Authorization': 'Bearer ' + store.accessToken
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
};

export const deleteUserAPI = async (id :string) => {
  const store = useStore();
  const config : AxiosRequestConfig = {
    method: 'delete',
    url: usersURL + `/${id}`,
    headers: {
      'Content-Type': 'application/json',
      'Authorization': 'Bearer ' + store.accessToken
    },
  };
  let response :AxiosResponse;
  try {
    response = await axios(config);
  } catch (error) {
    response = error.response;
  }
  return response;
};
