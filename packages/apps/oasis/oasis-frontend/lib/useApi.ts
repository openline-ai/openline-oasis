import axios, {AxiosInstance} from "axios";

let api = axios.create({
    baseURL: `${process.env.NEXT_PUBLIC_BE_PATH}`,
    timeout: 30000
});
api.interceptors.response.use((response) => response, (error) => {
    console.log(error);
    if (error.response.status === 401 && error.config.url !== `${process.env.NEXT_PUBLIC_BE_PATH}/account`) {
        //todo this needs to be changed when we introduce permissions in BE

    }
    return Promise.reject(error);
});

export function useApi(): AxiosInstance {
    return api;
}