import type {NextPage} from 'next'
import {DataTable} from 'primereact/datatable';
import {useEffect, useState} from "react";
import {Column} from "primereact/column";
import {Button} from "primereact/button";
import {useRouter} from "next/router";
import {Toolbar} from "primereact/toolbar";
import {Fragment} from "preact";
import Layout from "../../components/layout/layout";
import useWebSocket from 'react-use-websocket';
import axios from "axios";
import {loggedInOrRedirectToLogin} from "../../utils/logged-in";
import {getSession} from "next-auth/react";


const FeedPage: NextPage = () => {
    const router = useRouter()
    const [feeds, setFeeds] = useState([] as any)

    const {lastMessage} = useWebSocket(`${process.env.NEXT_PUBLIC_WEBSOCKET_PATH}`, {
        onOpen: () => console.log('Websocket opened'),
        //Will attempt to reconnect on all close events, such as server shutting down
        shouldReconnect: (closeEvent) => true,
    });

    useEffect(() => {
        axios.get(`/server/feed`)
            .then(res => {
                setFeeds(res.data?.contact)
                console.log(JSON.stringify(res.data?.contact))
            })
    }, []);

    useEffect(() => {
        if (lastMessage && Object.keys(lastMessage).length !== 0) {
            handleWebsocketMessage(lastMessage);
        }

    }, [lastMessage, setFeeds]);

    const actionsColumn = (rowData: any) => {
        return <Button icon="pi pi-eye" className="p-button-info"
                       onClick={() => router.push(`/feed/${rowData.id}`)}/>;
    }

    const handleWebsocketMessage = function (msg: any) {
        console.log("Got a new feed!");
        axios.get(`${process.env.NEXT_PUBLIC_BE_PATH}/feed`)
            .then(res => {
                setFeeds(res.data?.contact);
            });
    }

    return (
        <>
                <DataTable value={feeds}>
                    <Column field="firstName" header="First Name"></Column>
                    <Column field="lastName" header="Last Name"></Column>
                    <Column field="email" header="E-Mail"></Column>
                    <Column field="phone" header="Phone"></Column>
                    <Column field="actions" header="Actions" align={'right'} body={actionsColumn}></Column>
                </DataTable>
        </>
    );
}

export async function getServerSideProps(context: any) {
    return loggedInOrRedirectToLogin(await getSession(context));
}

export default FeedPage
