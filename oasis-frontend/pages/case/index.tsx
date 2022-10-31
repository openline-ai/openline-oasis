import type {NextPage} from 'next'
import {DataTable} from 'primereact/datatable';
import {useEffect, useState} from "react";
import {Column} from "primereact/column";
import axios from "axios";
import {Button} from "primereact/button";
import {useRouter} from "next/router";
import {Toolbar} from "primereact/toolbar";
import {Fragment} from "preact";
import Layout from "../../components/layout/layout";
import {useStomp, configureStomp} from "./useStomp";
import {IFrame} from "@stomp/stompjs";

const Index: NextPage = () => {
    const router = useRouter();
    const [cases, setCases] = useState([] as any);
    const incomingMsg:IFrame = useStomp();

    useEffect(() => {
        axios.get(`${process.env.NEXT_PUBLIC_BE_PATH}/case`)
            .then(res => {
                setCases(res.data.content);
            })
        configureStomp(`${process.env.NEXT_PUBLIC_STOMP_WEBSOCKET_PATH}/websocket`, `/queue/cases`);

    }, []);

    useEffect(() => {
        if (incomingMsg && Object.keys(incomingMsg).length !== 0) {
            handleWebsocketMessage(incomingMsg);
        }
    }, [incomingMsg]);

    const actionsColumn = (rowData: any) => {
        return <Button icon="pi pi-eye" className="p-button-info"
                       onClick={() => router.push(`/case/${rowData.id}`)}/>;
    }

    const leftContents = (
        <Fragment>
        </Fragment>
    );

    const handleWebsocketMessage = function (msg: any) {
        console.log("Got a new case!");
        axios.get(`${process.env.NEXT_PUBLIC_BE_PATH}/case`)
            .then(res => {
                setCases(res.data.content);
            });
    }

    return (
        <>
            <Layout>
                <Toolbar left={leftContents}/>
                <DataTable value={cases}>
                    <Column field="userName" header="Name"></Column>
                    <Column field="state" header="State"></Column>
                    <Column field="actions" header="Actions" align={'right'} body={actionsColumn}></Column>
                </DataTable>
            </Layout>
        </>
    );
}

export default Index
