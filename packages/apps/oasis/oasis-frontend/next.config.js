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

        NEXTAUTH_URL: process.env.NEXTAUTH_URL,
        NEXTAUTH_OAUTH_CLIENT_ID: process.env.NEXTAUTH_OAUTH_CLIENT_ID,
        NEXTAUTH_OAUTH_CLIENT_SECRET: process.env.NEXTAUTH_OAUTH_CLIENT_SECRET,
        NEXTAUTH_OAUTH_TENANT_ID: process.env.NEXTAUTH_OAUTH_TENANT_ID,
        NEXTAUTH_OAUTH_SERVER_URL: process.env.NEXTAUTH_OAUTH_SERVER_URL,
        NEXTAUTH_SECRET: process.env.NEXTAUTH_SECRET,

        OASIS_API_KEY: process.env.OASIS_API_KEY,
        CUSTOMER_OS_API_KEY: process.env.CUSTOMER_OS_API_KEY
    },
    output: 'standalone'
}

module.exports = nextConfig
