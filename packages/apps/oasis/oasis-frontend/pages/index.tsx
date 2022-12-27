import type {NextPage} from 'next';
import {useEffect} from "react";
import {useRouter} from "next/router";
import {getSession} from "next-auth/react"
import {loggedInOrRedirectToLogin} from "../utils/logged-in";


const Home: NextPage = () => {
    const router = useRouter();

    useEffect(() => {
        router.push('/feed');
    });

    return (
        <>
        </>
    )
}

export async function getServerSideProps(context: any) {
    return loggedInOrRedirectToLogin(await getSession(context));
}

export default Home
