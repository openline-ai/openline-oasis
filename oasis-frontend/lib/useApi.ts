import axios, {AxiosInstance} from "axios";
import {Router} from "next/router";
import {reloadUserAccount, removeUserAccount} from "./loadUserAccount";

let api = axios.create({
    baseURL: `${process.env.NEXT_PUBLIC_BE_PATH}`,
    timeout: 30000
});
api.interceptors.response.use((response) => response, (error) => {
    console.log(error);
    if(error.response.status === 401 && error.config.url !== `${process.env.NEXT_PUBLIC_BE_PATH}/account`) {
        //todo this needs to be changed when we introduce permissions in BE
        removeUserAccount();
        window.location.href = '/login';
    }
    return Promise.reject(error);
});

export function useApi(): AxiosInstance {
    return api;
}