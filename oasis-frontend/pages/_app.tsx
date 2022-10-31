import 'primereact/resources/themes/lara-light-blue/theme.css';
import 'primereact/resources/primereact.min.css';
import 'primeflex/primeflex.css';
import 'primeicons/primeicons.css';

import '../styles/globals.css'
import "../styles/login.css";
import axios from "axios";

axios.defaults.withCredentials = true

export default function App({
                                Component,
                                pageProps
                            }: any) {

    return <Component {...pageProps} />;

}