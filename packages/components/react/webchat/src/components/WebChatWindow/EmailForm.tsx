import React, {useState} from "react"

import {styles} from "./styles"

import {LoadingOutlined} from '@ant-design/icons'

import Avatar from '../Avatar'

interface EmailPros {
    onSetUser: (user: string) => void
    visible: boolean
    apikey: string
    httpServerPath: string
}

export default function EmailForm(props: EmailPros) {
    const [email, setEmail] = useState('')
    const [loading, setLoading] = useState(false)

    function getOrCreateUser(callback: any) {
        fetch(`${props.httpServerPath}/webchat/?email=${email}`, {
            headers: {
                "Content-Type": "application/json",
                'WebChatApiKey': `${props.apikey}`
            },
        }).then((response) => {
            if (!response.ok) {
                throw Error(response.statusText);
            }
            return response.json();
        }).then(response => {
            let stringify = JSON.stringify(response);
            console.log(stringify)
            callback(response.username)
        }).catch(error => {
            console.log(error);
        });

    }

    function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
        event.preventDefault()
        setLoading(true)
        getOrCreateUser(
            (user: string) => {
                setLoading(false)
                props.onSetUser(user)
            }
        )
    }

    return (
        <div
            style={{
                ...styles.emailFormWindow,
                ...{
                    height: props.visible ? '100%' : '0px',
                    opacity: props.visible ? '1' : '0'
                }
            }}
        >
            <div style={{height: '0px'}}>
                <div style={styles.stripe}/>
            </div>

            <div
                className='transition-5'
                style={{
                    ...styles.loadingDiv,
                    ...{
                        zIndex: loading ? '10' : '-1',
                        opacity: loading ? '0.33' : '0',
                    }
                }}
            />
            <LoadingOutlined
                className='transition-5'
                style={{
                    ...styles.loadingIcon,
                    ...{
                        zIndex: loading ? '10' : '-1',
                        opacity: loading ? '1' : '0',
                        fontSize: '82px',
                        top: 'calc(50% - 41px)',
                        left: 'calc(50% - 41px)',
                    }
                }}
            />

            <div style={{position: 'absolute', height: '100%', width: '100%', textAlign: 'center'}}>
                <Avatar onClick={() => {
                }}
                        style={{
                            position: 'relative',
                            left: 'calc(50% - 44px)',
                            top: '10%',
                        }}
                />

                <div style={styles.topText}>
                    Welcome to my <br/> support 👋
                </div>

                <form
                    onSubmit={e => handleSubmit(e)}
                    style={{position: 'relative', width: '100%', top: '19.75%'}}
                >
                    <input
                        placeholder='Your Email'
                        onChange={e => setEmail(e.target.value)}
                        style={styles.emailInput}
                    />
                </form>

                <div style={styles.bottomText}>
                    Enter your email <br/> to get started.
                </div>
            </div>
        </div>
    )
}