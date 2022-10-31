import Header from "./header";
import LayoutMenu from "./menu";
import {useRouter} from "next/router";
import {useEffect, useState} from "react";
import {getUserAccount, loadUserAccount, User} from "../../lib/loadUserAccount";
import {useApi} from "../../lib/useApi";

export default function Layout({children}: any) {
    const router = useRouter();
    const [user, setUser] = useState<User | undefined>();
    const axiosInstance = useApi();

    useEffect(() => {
        loadUserAccount(axiosInstance).then(() => {
            if (!getUserAccount()) {
                router.push('/login'); //todo check if we redirect to login
            } else {
                setUser(getUserAccount);
            }
        })
    }, []);

    return (
        <>
            {!user &&
                <div>loading</div>
            }

            {user &&
                <>
                    <Header height={'70px'}/>

                    <div className="flex" style={{height: 'calc(100vh - 90px)'}}>

                        <div className="flex-grow-0 flex"
                             style={{width: '200px', height: '100%'}}>

                            <div style={{
                                width: '100%',
                                height: '100%',
                                padding: '10px 0px 10px 10px',
                                overflow: 'hidden'
                            }}>

                                <LayoutMenu/>

                            </div>

                        </div>

                        <div className="flex-grow-1 flex" style={{height: '100%'}}>
                            <div style={{
                                width: '100%',
                                height: '100%',
                                margin: '10px',
                                border: '1px solid #0b213f',
                                borderRadius: '6px'
                            }}>
                                {children}
                            </div>
                        </div>
                    </div>
                </>
            }

        </>
    )
}