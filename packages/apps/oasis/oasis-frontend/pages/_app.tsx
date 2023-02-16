import 'primereact/resources/themes/lara-light-blue/theme.css';
import 'primereact/resources/primereact.min.css';
import 'primeflex/primeflex.css';
import 'primeicons/primeicons.css';

import '../styles/globals.css'
import '../styles/theme-override.css'
import '../styles/layout.css'
import 'react-toastify/dist/ReactToastify.css';

import * as React from "react";
import {AppProps} from "next/app";

import axios from "axios";

axios.defaults.withCredentials = true

export default function App({
                              Component,
                              pageProps: {session, ...pageProps},
                            }: AppProps) {

  return (
    <Component {...pageProps} />
  )
}