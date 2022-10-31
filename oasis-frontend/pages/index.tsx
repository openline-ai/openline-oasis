import type {NextPage} from 'next';
import {useEffect} from "react";
import {useRouter} from "next/router";
import {loadUserAccount} from "../lib/loadUserAccount";
import {useApi} from "../lib/useApi";


const Home: NextPage = () => {
    const router = useRouter();
    const axiosInstance = useApi();

    useEffect(() => {
        loadUserAccount(axiosInstance).then((userData) => {
            if(userData) {
                router.push('/case'); //todo switch to default user path ( depending on role )
            } else {
                router.push('/login');
            }
        });

        function checkUserData() {
            //storage changed. we need a way to notify the user to refresh or smth
        }

        window.addEventListener('storage', checkUserData)

        return () => {
            window.removeEventListener('storage', checkUserData)
        }
    }, []);


    return (
        <>
            <div>Loading ...</div>
        </>
    )
}

export default Home
