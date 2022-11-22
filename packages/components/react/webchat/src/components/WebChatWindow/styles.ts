import {CSSProperties} from "react";
import logoUrl from './teamLogo.png';

export const styles: { [key: string]: CSSProperties } = {
    supportWindow: {
        // Position
        position: 'fixed',
        bottom: '116px',
        right: '24px',
        // Size
        width: '420px',
        height: '530px',
        maxWidth: 'calc(100% - 48px)',
        maxHeight: 'calc(100% - 48px)',
        backgroundColor: 'white',
        // Border
        borderRadius: '12px',
        border: `2px solid #7a39e0`,
        overflow: 'hidden',
        // Shadow
        boxShadow: '0px 0px 16px 6px rgba(0, 0, 0, 0.33)',
    },
    emailFormWindow: {
        width: '100%',
        overflow: 'hidden',
        transition: "all 0.5s ease",
        WebkitTransition: "all 0.5s ease",
        MozTransition: "all 0.5s ease",
    },
    stripe: {
        position: 'relative',
        top: '-45px',
        width: '100%',
        height: '308px',
        backgroundColor: '#7a39e0',
        transform: 'skewY(-12deg)',
    },
    topText: {
        position: 'relative',
        width: '100%',
        top: '15%',
        color: 'white',
        fontSize: '24px',
        fontWeight: '600',
    },
    emailInput: {
        width: '66%',
        textAlign: 'center',
        outline: 'none',
        padding: '12px',
        borderRadius: '12px',
        border: '2px solid #7a39e0',
    },
    bottomText: {
        position: 'absolute',
        width: '100%',
        top: '60%',
        color: '#7a39e0',
        fontSize: '24px',
        fontWeight: '600'
    },
    loadingDiv: {
        position: 'absolute',
        height: '100%',
        width: '100%',
        textAlign: 'center',
        backgroundColor: 'white',
    },
    loadingIcon: {
        color: '#7a39e0',
        position: 'absolute',
        top: 'calc(50% - 51px)',
        left: 'calc(50% - 51px)',
        fontWeight: '600',
    },
    chatEngineWindow: {
        width: '100%',
        backgroundColor: '#fff',
    },
    divStyle: {
        position: 'absolute',
        bottom: '10px',
        width: '100%'
    },

}