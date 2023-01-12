export type ConversationItem = {
    id:             string,
    conversationId: string,
    type:           number,
    subtype:        number,
    content:        string,
    direction:      number,
    time:           Time,
    senderType:     number,
    senderId:       string,
}

export type Time = {
    seconds: number,
}
