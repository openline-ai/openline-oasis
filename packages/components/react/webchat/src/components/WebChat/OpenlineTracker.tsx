import React, {useEffect} from 'react';
import {enableActivityTracking, newTracker, trackPageView} from '@snowplow/browser-tracker';
import {enableLinkClickTracking, LinkClickTrackingPlugin} from '@snowplow/browser-plugin-link-click-tracking';

interface OpenlineTrackerProps {
    enabled: boolean
    trackerId: string
    appId: string
    collectorUrl: string
    bufferSize: string
    minimumVisitLength: string
    heartbeatDelay: string
}

export default function OpenlineTracker(props: OpenlineTrackerProps) {

    function propertyOrDefault<T>(propertyValue: T, defaultValue: T) {
        if (propertyValue !== undefined) {
            return propertyValue;
        }
        return defaultValue;
    }

    useEffect(() => {
        if (props.enabled) {
            newTracker(props.trackerId, props.collectorUrl, {
                appId: props.appId,
                discoverRootDomain: true,
                cookieSecure: true,
                cookieSameSite: "None",
                eventMethod: "post",
                postPath: "/ai.openline.sp/tp2",
                platform: "web",
                bufferSize: parseInt(propertyOrDefault<string>(props.bufferSize, "1")),
                contexts: {
                    webPage: true
                },
                plugins: [LinkClickTrackingPlugin()],
            });

            enableActivityTracking({
                minimumVisitLength: parseInt(propertyOrDefault(props.minimumVisitLength, "30")),
                heartbeatDelay: parseInt(propertyOrDefault(props.heartbeatDelay, "30"))
            });

            enableLinkClickTracking({
                pseudoClicks: true,
                //denylist: ['untracked'],
                trackContent: true
            });

            trackPageView({}, [props.trackerId]);
        }
    }, []);

    return (
        <></>
    )
}