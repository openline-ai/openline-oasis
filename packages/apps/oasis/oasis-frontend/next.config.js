/** @type {import('next').NextConfig} */
const nextConfig = {
    reactStrictMode: true,
    swcMinify: true,
    env: {
        NEXT_PUBLIC_OASIS_API_PATH: process.env.NEXT_PUBLIC_OASIS_API_PATH,
        NEXT_PUBLIC_CUSTOMER_OS_API_PATH: process.env.NEXT_PUBLIC_CUSTOMER_OS_API_PATH,

        NEXT_PUBLIC_WEBRTC_WEBSOCKET_URL: process.env.NEXT_PUBLIC_WEBRTC_WEBSOCKET_URL,
        NEXT_PUBLIC_WEBSOCKET_PATH: process.env.NEXT_PUBLIC_WEBSOCKET_PATH,

        NEXT_PUBLIC_TURN_SERVER: process.env.NEXT_PUBLIC_TURN_SERVER,
        NEXT_PUBLIC_TURN_USER: process.env.NEXT_PUBLIC_TURN_USER,

        OASIS_API_KEY: process.env.OASIS_API_KEY,
        CUSTOMER_OS_API_KEY: process.env.CUSTOMER_OS_API_KEY
    },
    output: 'standalone'
}

module.exports = nextConfig
