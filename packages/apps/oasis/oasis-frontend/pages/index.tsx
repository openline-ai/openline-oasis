import type {NextPage} from 'next';
import {useEffect, useState} from "react";
import {useRouter} from "next/router";
import {Configuration, FrontendApi, Session} from "@ory/client";
import {edgeConfig} from "@ory/integrations/next";
import {getUserName} from "../utils/logged-in";
import {setClient} from "../utils/graphQLClient";
import axios from "axios";
import {GraphQLClient} from "graphql-request";
import {getReturnToUrl} from "./feed/index"

const ory = new FrontendApi(new Configuration(edgeConfig))

const Home: NextPage = () => {
    const router = useRouter();

    const [session, setSession] = useState<Session | undefined>()
    const [userEmail, setUserEmail] = useState<string | undefined>()
    const [logoutUrl, setLogoutUrl] = useState<string | undefined>()

    useEffect(() => {
        ory
                .toSession()
                .then(({data}) => {

                    console.log('HAVE SESSION')
                    console.log(data)

                    let userName = getUserName(data.identity);
                    setUserEmail(userName)

                    let graphQLClient = new GraphQLClient(`/customer-os-api/query`, {
                        headers: {
                            'X-Openline-USERNAME': userName
                        }
                    });

                    setClient(graphQLClient)
                    axios.defaults.headers.common['X-Openline-USERNAME'] = userName;

                    // Create a logout url
                    ory.createBrowserLogoutFlow().then(({data}) => {
                        setLogoutUrl(data.logout_url)
                    })

                    // User has a session!
                    setSession(data)

                })
                .catch((e) => {
                    // Redirect to login page
                    console.log('NO SESSION')
                    console.log(e)
                    return router.push(edgeConfig.basePath + "/ui/login" + getReturnToUrl())
                })
    }, [router])

    if (!session) {
        console.log('checking for session. no session')
        // Still loading
        return null
    } else {
        console.log('checking for session. have session')
        router.push('/feed');
    }

    console.log('printing empy html')
    return (
        <>
        </>
    )
}


export default Home
