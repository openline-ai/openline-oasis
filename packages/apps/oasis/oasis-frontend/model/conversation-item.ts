import {Time} from "./time";

export type MessageId = {
  conversationEventId: string,
  conversationId: string,
}

export type ConversationItem = {
  messageId: MessageId,
  type: number,
  subtype: number,
  content: string,
  direction: number,
  time: Time,
  senderType: number,
  senderId: string,
  senderUsername: string,
}

export type FeedPostRequest = {
  username: string,
  message: string,
  channel: string,
  direction: string,
  destination: string[],
  replyTo?: string,
}