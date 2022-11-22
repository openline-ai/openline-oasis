import React, {useState} from "react";
import {styles} from './styles'

interface IAvatarProps {
    onClick: () => void,
    style: {}
}

export default function Avatar(props: IAvatarProps) {

    const [hovered, isHovered] = useState(false)

    return (
        <div style={props.style}>
            <div
                className='transition-3'
                style={{
                    ...styles.avatarHello,
                    ...{opacity: hovered ? '1' : '0'}
                }}
            >
                {" Send us a message"}
            </div>
            <div
                onMouseEnter={() => isHovered(true)}
                onMouseLeave={() => isHovered(false)}
                onClick={() => props.onClick && props.onClick()}
                className='transition-3'
                style={{
                    ...styles.chatWithOpenlineButton,
                    ...{border: hovered ? '1px solid #f9f0ff' : '4px solid #7a39e0'}
                }}
            />
        </div>
    )
}
