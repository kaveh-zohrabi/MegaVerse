import { NextResponse } from 'next/server';
import { getUserByUsername } from '../../../lib/store';

export async function POST(request: Request) {
  try {
    const { username, password } = await request.json();

    if (!username || !password) {
      return NextResponse.json({ error: 'username and password required' }, { status: 400 });
    }

    const user = getUserByUsername(username);
    if (!user || user.password !== password) {
      return NextResponse.json({ error: 'invalid credentials' }, { status: 401 });
    }

    user.online = true;
    return NextResponse.json({ id: user.id, username: user.username });
  } catch {
    return NextResponse.json({ error: 'server error' }, { status: 500 });
  }
}
