import type {NextPage} from 'next';
import {useEffect, useState} from "react";
import {useRouter} from "next/router";
import axios from "axios";
import {InputText} from "primereact/inputtext";
import {Button} from "primereact/button";
import {loadUserAccount} from "../lib/loadUserAccount";
import {useApi} from "../lib/useApi";

const Login: NextPage = () => {
    const router = useRouter();
    const axiosInstance = useApi();

    const [showLoginForm, setShowLoginForm] = useState(false);

    useEffect(() => {
        loadUserAccount(axiosInstance).then((userData) => {
            if (userData) {
                router.push('/case'); //todo switch to default user path ( depending on role )
            } else {
                setShowLoginForm(true);
            }
        });
    }, []);

    const [state, setState] = useState({
        error: null,
        username: '',
        password: '',
    });

    const handleLogin = async () => {
        state.error = null;

        await axios.post(`${process.env.NEXT_PUBLIC_BE_PATH}/login`, {
            username: state.username,
            password: state.password
        }).then(response => {
            console.log(response);
            if (response.status === 200) {
                localStorage.setItem('userData', JSON.stringify(response.data));
                router.push('/case'); //todo switch to default user path ( depending on role )
            }
        }).catch(err => {
            //todo show error
            return {
                error: true,
                response: err.response
            }
        });
    }

    return (
        <>
            {!showLoginForm &&
                <div>Loading ...</div>
            }

            {showLoginForm &&

                <div className="flex align-items-center justify-content-center">
                    <div className="surface-card p-4 shadow-2 border-round w-full lg:w-6">
                        <div className="text-center mb-5">

                            <svg height="50" viewBox="0 0 147 35" fill="black"
                                 xmlns="http://www.w3.org/2000/svg">
                                <path
                                    d="M0 17.5C0 7.83502 7.83502 0 17.5 0C27.165 0 35 7.83502 35 17.5C35 27.165 27.165 35 17.5 35C7.83502 35 0 27.165 0 17.5Z"
                                    fill="#7626FA"/>
                                <mask id="mask0_1_46" style={{maskType: 'alpha'}} maskUnits="userSpaceOnUse" x="3" y="3"
                                      width="29" height="29">
                                    <circle cx="17.5" cy="17.5" r="13.6719" fill="#EAE7E4"/>
                                </mask>
                                <g mask="url(#mask0_1_46)">
                                    <path fillRule="evenodd" clipRule="evenodd"
                                          d="M12.4403 4.79485C12.9974 4.57281 13.5725 4.38635 14.1628 4.23827V30.7616C13.5725 30.6136 12.9974 30.4271 12.4403 30.2051V4.79485ZM31.1719 17.5C31.1719 25.0148 25.1089 31.1135 17.6077 31.1714V23.5852C19.1829 23.5574 20.6876 22.9195 21.8036 21.8036C22.945 20.6622 23.5862 19.1141 23.5862 17.5C23.5862 15.8858 22.945 14.3378 21.8036 13.1964C20.6876 12.0804 19.1829 11.4426 17.6077 11.4147V3.82849C25.1089 3.88638 31.1719 9.98512 31.1719 17.5ZM9.64138 6.31095C9.02827 6.74237 8.45221 7.2229 7.91894 7.74682V27.2531C8.45221 27.777 9.02827 28.2575 9.64138 28.6889V6.31095ZM3.82812 17.5C3.82812 14.9331 4.53548 12.5315 5.76589 10.4794V24.5205C4.53548 22.4684 3.82812 20.0668 3.82812 17.5Z"
                                          fill="#EAE7E4"/>
                                </g>
                                <path
                                    d="M42.9474 18.37C42.9474 22.27 46.0414 25.39 50.2274 25.39C54.3614 25.39 57.5074 22.27 57.5074 18.37C57.5074 14.47 54.3614 11.35 50.2274 11.35C46.0414 11.35 42.9474 14.47 42.9474 18.37ZM46.7174 18.37C46.7174 16.03 48.2774 14.47 50.2274 14.47C52.1774 14.47 53.7374 16.03 53.7374 18.37C53.7374 20.71 52.1774 22.27 50.2274 22.27C48.2774 22.27 46.7174 20.71 46.7174 18.37ZM58.8165 29.94H62.4565V23.31H62.5865C62.8725 23.7 63.2105 24.038 63.6005 24.35C64.3025 24.87 65.2905 25.39 66.7465 25.39C70.1265 25.39 72.8565 22.66 72.8565 18.37C72.8565 14.08 70.1265 11.35 66.7465 11.35C65.2905 11.35 64.3025 11.87 63.6005 12.39C63.2105 12.702 62.8725 13.04 62.5865 13.43H62.4565V11.74H58.8165V29.94ZM62.4565 18.37C62.4565 15.874 63.8605 14.47 65.8365 14.47C67.6825 14.47 69.0865 15.874 69.0865 18.37C69.0865 20.866 67.6825 22.27 65.8365 22.27C63.8605 22.27 62.4565 20.866 62.4565 18.37ZM73.5065 18.37C73.5065 22.27 76.6005 25.39 80.6565 25.39C83.5165 25.39 85.2065 24.22 86.2465 23.05C86.8445 22.348 87.2605 21.568 87.5465 20.71H83.6465C83.4905 20.996 83.2825 21.256 83.0225 21.49C82.5285 21.88 81.8265 22.27 80.6565 22.27C78.9665 22.27 77.5625 20.84 77.2765 19.15H87.6765V18.37C87.6765 14.496 84.8165 11.35 80.6565 11.35C76.6005 11.35 73.5065 14.47 73.5065 18.37ZM77.4065 16.94C77.7965 15.614 78.9665 14.47 80.6565 14.47C82.3465 14.47 83.5165 15.64 83.7765 16.94H77.4065ZM89.0964 25H92.7364V17.98C92.7364 15.9 94.1664 14.47 95.8564 14.47C97.5464 14.47 98.5864 15.484 98.5864 17.46V25H102.226V16.94C102.226 13.274 100.276 11.35 97.0264 11.35C95.4924 11.35 94.5304 11.922 93.8284 12.468C93.4384 12.78 93.1004 13.144 92.8664 13.56H92.7364V11.74H89.0964V25ZM104.186 25H110.816V21.88H107.826V6.8H104.186V25ZM111.474 8.1C111.474 9.14 112.384 10.05 113.554 10.05C114.724 10.05 115.634 9.14 115.634 8.1C115.634 7.06 114.724 6.15 113.554 6.15C112.384 6.15 111.474 7.06 111.474 8.1ZM111.734 25H115.374V11.74H111.734V25ZM117.454 25H121.094V17.98C121.094 15.9 122.524 14.47 124.214 14.47C125.904 14.47 126.944 15.484 126.944 17.46V25H130.584V16.94C130.584 13.274 128.634 11.35 125.384 11.35C123.85 11.35 122.888 11.922 122.186 12.468C121.796 12.78 121.458 13.144 121.224 13.56H121.094V11.74H117.454V25ZM131.763 18.37C131.763 22.27 134.857 25.39 138.913 25.39C141.773 25.39 143.463 24.22 144.503 23.05C145.101 22.348 145.517 21.568 145.803 20.71H141.903C141.747 20.996 141.539 21.256 141.279 21.49C140.785 21.88 140.083 22.27 138.913 22.27C137.223 22.27 135.819 20.84 135.533 19.15H145.933V18.37C145.933 14.496 143.073 11.35 138.913 11.35C134.857 11.35 131.763 14.47 131.763 18.37ZM135.663 16.94C136.053 15.614 137.223 14.47 138.913 14.47C140.603 14.47 141.773 15.64 142.033 16.94H135.663Z"
                                    fill="#EAE7E4"/>
                            </svg>
                        </div>

                        <div>
                            <label htmlFor="email" className="block text-900 font-medium mb-2">Email</label>
                            <InputText id="email" type="text" className="w-full mb-3" value={state.username}
                                       onChange={(e: any) => {
                                           setState((prevState) => {
                                               return {...prevState, ...{username: e.target.value}};
                                           });
                                       }}/>

                            <label htmlFor="password" className="block text-900 font-medium mb-2">Password</label>
                            <InputText id="password" type="password" className="w-full mb-5" value={state.password}
                                       onChange={(e: any) => {
                                           setState((prevState) => {
                                               return {...prevState, ...{password: e.target.value}};
                                           });
                                       }}/>

                            <div className="flex align-items-center justify-content-between mb-6">
                                <a className="font-medium no-underline ml-2 text-blue-500 cursor-pointer">Create your account!</a>
                                <a className="font-medium no-underline ml-2 text-blue-500 text-right cursor-pointer">Forgot your password?</a>
                            </div>

                            <Button label="Sign In" icon="pi pi-user" className="w-full" onClick={() => handleLogin()}/>
                        </div>
                    </div>
                </div>

            }
        </>
    );

}

export default Login
