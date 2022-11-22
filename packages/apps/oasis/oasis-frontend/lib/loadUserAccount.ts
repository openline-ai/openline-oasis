import {AxiosInstance} from "axios";

export type User = {
    email: string;
    data: {},
    roles: string[];
}

export async function loadUserAccount(axiosInstance: AxiosInstance): Promise<User | undefined> {
    if (!localStorage.getItem('userData')) {
        return await axiosInstance.get(`${process.env.NEXT_PUBLIC_BE_PATH}/account`).then(r => {
            console.log('account fetch');
            console.log(r);
            if (r && r.status === 200) {
                localStorage.setItem('userData', JSON.stringify(r.data));
                return r.data;
            } else {
                console.log(r);
                localStorage.removeItem('userData');
                return undefined;
            }
        }).catch(reason => {
            //todo popup error
            return undefined;
        });
    } else {
        return getUserAccount();
    }
}

export function removeUserAccount(): void {
    localStorage.removeItem('userData');
}

export function reloadUserAccount(axiosInstance: AxiosInstance): Promise<User | undefined> {
    localStorage.removeItem('userData');
    return loadUserAccount(axiosInstance);
}

export function getUserAccount(): User | undefined {
    if (typeof window !== 'undefined' && localStorage.getItem('userData')) {
        return JSON.parse(localStorage.getItem('userData') as string);
    } else {
        return undefined;
    }
}
