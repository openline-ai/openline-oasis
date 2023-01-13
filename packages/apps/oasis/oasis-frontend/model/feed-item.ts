import {Time} from "./time";

export type FeedItem = {
    id: string;
    initiatorFirstName: string;
    initiatorLastName: string;
    initiatorUsername: string;
    initiatorType: string;
    lastSenderFirstName: string;
    lastSenderLastName: string;
    lastContentPreview: string;
    lastTimestamp: Time;
}
