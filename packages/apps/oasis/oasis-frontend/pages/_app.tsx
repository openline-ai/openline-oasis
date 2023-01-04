import 'primereact/resources/themes/lara-light-blue/theme.css';
import 'primereact/resources/primereact.min.css';
import 'primeflex/primeflex.css';
import 'primeicons/primeicons.css';
import '@fortawesome/fontawesome-svg-core/styles.css'
import { config } from '@fortawesome/fontawesome-svg-core'

import '../styles/globals.css'
import '../styles/theme-override.css'
import '../styles/layout.css'
import axios from "axios";
import {SessionProvider} from "next-auth/react"

config.autoAddCss = false

axios.defaults.withCredentials = true

export default function App({
                                Component,
                                pageProps: {session, ...pageProps}
                            }: any) {

    return (
        <SessionProvider session={session}>
            <Component {...pageProps} />
        </SessionProvider>
    );

}