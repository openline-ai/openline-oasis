import type {NextPage} from 'next';
import {useEffect} from "react";
import {useRouter} from "next/router";
import { useSession, signIn, signOut } from "next-auth/react"


const Home: NextPage = () => {
    const router = useRouter();

    useEffect(() => {
                router.push('/feed'); //todo switch to default user path ( depending on role )

    }, []);


    return (
        <>
            <div>Loading ...</div>
        </>
    )
}

export default Home
