import '../styles/globals.css'
import '@openline-ai/openline-web-chat/dist/esm/index.css'
import type {AppProps} from 'next/app'

export default function App({Component, pageProps}: AppProps) {
    return <Component {...pageProps} />
}
