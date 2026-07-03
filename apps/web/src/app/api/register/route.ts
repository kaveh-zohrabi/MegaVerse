import { NextResponse } from 'next/server';
import { addUser, getUserByUsername } from '../../../lib/store';

export async function POST(request: Request) {
  try {
    const { username, password } = await request.json();

    if (!username || !password) {
      return NextResponse.json({ error: 'username and password required' }, { status: 400 });
    }

    if (getUserByUsername(username)) {
      return NextResponse.json({ error: 'username already exists' }, { status: 409 });
    }

    const id = `user_${Date.now()}`;
    const user = addUser(id, username, password);

    return NextResponse.json({ id: user.id, username: user.username });
  } catch {
    return NextResponse.json({ error: 'server error' }, { status: 500 });
  }
}
