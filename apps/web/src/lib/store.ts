export interface User {
  id: string;
  username: string;
  password: string;
  online: boolean;
}

export interface Message {
  id: number;
  conversation_id: string;
  sender_id: string;
  sender_name: string;
  content: string;
  created_at: string;
}

export interface Conversation {
  id: string;
  name: string;
  is_group: boolean;
  updated_at: string;
}

export const userStore: Record<string, User> = {};
export const messageStore: Record<string, Message[]> = {};
export const conversationStore: Record<string, Conversation> = {};
let nextMsgId = 1;

export function addUser(id: string, username: string, password: string): User {
  userStore[id] = { id, username, password, online: true };
  return userStore[id];
}

export function getUserByUsername(username: string): User | undefined {
  return Object.values(userStore).find(u => u.username === username);
}

export function getUserById(id: string): User | undefined {
  return userStore[id];
}

export function getConvId(userId1: string, userId2: string): string {
  const sorted = [userId1, userId2].sort();
  return `dm_${sorted[0]}_${sorted[1]}`;
}

export function addMessage(convId: string, senderId: string, senderName: string, content: string): Message {
  if (!messageStore[convId]) {
    messageStore[convId] = [];
  }
  const msg: Message = {
    id: nextMsgId++,
    conversation_id: convId,
    sender_id: senderId,
    sender_name: senderName,
    content,
    created_at: new Date().toISOString(),
  };
  messageStore[convId].push(msg);

  if (!conversationStore[convId]) {
    conversationStore[convId] = {
      id: convId,
      name: '',
      is_group: false,
      updated_at: msg.created_at,
    };
  }
  conversationStore[convId].updated_at = msg.created_at;

  return msg;
}

export function getMessages(convId: string): Message[] {
  return messageStore[convId] || [];
}

export function getAllUsers(): Array<{ id: string; username: string; online: boolean }> {
  return Object.values(userStore).map(u => ({
    id: u.id,
    username: u.username,
    online: u.online,
  }));
}
