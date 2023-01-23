import type {NextPage} from 'next';
import {useEffect, useState} from "react";
import {useRouter} from "next/router";
import {Configuration, FrontendApi, Session} from "@ory/client";
import {edgeConfig} from "@ory/integrations/next";
import {getUserName} from "../utils/logged-in";
import {setClient} from "../utils/graphQLClient";
import axios from "axios";
import {GraphQLClient} from "graphql-request";

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
                    // User has a session!
                    setSession(data)
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

                    router.push('/feed');
                })
                .catch(() => {
                    // Redirect to login page
                    return router.push(edgeConfig.basePath + "/ui/login")
                })
    }, [router])

    if (!session) {
        // Still loading
        return null
    }

    return (
        <>
        </>
    )
}

export default Home
