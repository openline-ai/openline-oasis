import type {NextPage} from 'next'
import * as React from "react";
import {useEffect, useState} from "react";
import {useRouter} from "next/router";
import axios from "axios";

import {GraphQLClient} from "graphql-request";
import {Configuration, FrontendApi, Session} from "@ory/client";
import {edgeConfig} from "@ory/integrations/next";
import {getUserName} from "../../utils/logged-in";
import {setClient} from "../../utils/graphQLClient";
import {Feed} from "./feed";

const ory = new FrontendApi(new Configuration(edgeConfig))

const FeedPage: NextPage = () => {
    const router = useRouter()
    const {id} = router.query;

    //region AUTH
    const [session, setSession] = useState<Session | undefined>()
    const [userEmail, setUserEmail] = useState<string | undefined>()
    const [logoutUrl, setLogoutUrl] = useState<string | undefined>()

    useEffect(() => {
        ory
                .toSession()
                .then(({data}) => {

                    let userName = getUserName(data.identity);
                    setUserEmail(userName);

                    let graphQLClient = new GraphQLClient(`/customer-os-api/query`, {
                        headers: {
                            'X-Openline-USERNAME': userName
                        }
                    });

                    setClient(graphQLClient);

                    axios.defaults.headers.common['X-Openline-USERNAME'] = userName;

                    // Create a logout url
                    ory.createBrowserLogoutFlow().then(({data}) => {
                        setLogoutUrl(data.logout_url)
                    })

                    // User has a session!
                    setSession(data)
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
    //endregion

    return (
            <>
                {
                        session && userEmail &&
                        <Feed feedId={id as string} logoutUrl={logoutUrl} userLoggedInEmail={userEmail}/>
                }
            </>
    );
}

export default FeedPage
