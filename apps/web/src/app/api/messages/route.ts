import { NextResponse } from 'next/server';
import { getMessages, addMessage, getConvId, getUserById } from '@/lib/store';

export async function GET(request: Request) {
  const { searchParams } = new URL(request.url);
  const convId = searchParams.get('conversation_id') || '';
  return NextResponse.json({ messages: getMessages(convId) });
}

export async function POST(request: Request) {
  try {
    const { conversation_id, sender_id, content } = await request.json();

    if (!conversation_id || !sender_id || !content) {
      return NextResponse.json({ error: 'missing fields' }, { status: 400 });
    }

    const user = getUserById(sender_id);
    const senderName = user ? user.username : 'Unknown';
    const msg = addMessage(conversation_id, sender_id, senderName, content);

    return NextResponse.json(msg);
  } catch {
    return NextResponse.json({ error: 'server error' }, { status: 500 });
  }
}
