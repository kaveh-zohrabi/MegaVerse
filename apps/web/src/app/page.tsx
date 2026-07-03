'use client';

import { useState, useEffect, useRef } from 'react';

interface Message {
  id: number;
  conversation_id: string;
  sender_id: string;
  sender_name: string;
  content: string;
  created_at: string;
}

interface User {
  id: string;
  username: string;
}

export default function Messenger() {
  const [user, setUser] = useState<User | null>(null);
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [isRegister, setIsRegister] = useState(false);
  const [users, setUsers] = useState<User[]>([]);
  const [messages, setMessages] = useState<Message[]>([]);
  const [input, setInput] = useState('');
  const [ws, setWs] = useState<WebSocket | null>(null);
  const [selectedUser, setSelectedUser] = useState<User | null>(null);
  const [currentConvID, setCurrentConvID] = useState('');
  const messagesEndRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    messagesEndRef.current?.scrollIntoView({ behavior: 'smooth' });
  }, [messages]);

  const connectWebSocket = (userId: string) => {
    const socket = new WebSocket(`ws://localhost:8080/ws?user_id=${userId}`);

    socket.onopen = () => {
      console.log('Connected to WebSocket');
    };

    socket.onmessage = (event) => {
      const data = JSON.parse(event.data);
      if (data.type === 'message' && data.message) {
        if (data.message.conversation_id === currentConvID) {
          setMessages(prev => [...prev, data.message]);
        }
      }
    };

    socket.onclose = () => {
      console.log('Disconnected from WebSocket');
    };

    setWs(socket);
    return socket;
  };

  const handleAuth = async () => {
    const endpoint = isRegister ? '/api/register' : '/api/login';
    const res = await fetch(`http://localhost:8080${endpoint}`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ username, password }),
    });

    if (res.ok) {
      const data = await res.json();
      setUser(data);
      connectWebSocket(data.id);
      loadUsers();
    } else {
      alert('خطا در ورود');
    }
  };

  const loadUsers = async () => {
    const res = await fetch('http://localhost:8080/api/users');
    if (res.ok) {
      const data = await res.json();
      setUsers(data.users || []);
    }
  };

  const startChat = (targetUser: User) => {
    setSelectedUser(targetUser);
    const convID = [user?.id, targetUser.id].sort().join('_');
    setCurrentConvID(`dm_${convID}`);
    loadMessages(`dm_${convID}`);
  };

  const loadMessages = async (convID: string) => {
    const res = await fetch(`http://localhost:8080/api/messages?conversation_id=${convID}`);
    if (res.ok) {
      const data = await res.json();
      setMessages(data.messages || []);
    }
  };

  const sendMessage = () => {
    if (!input.trim() || !ws || !selectedUser) return;

    ws.send(JSON.stringify({
      type: 'private_message',
      receiver_id: selectedUser.id,
      content: input,
    }));

    setInput('');
  };

  const handleKeyPress = (e: React.KeyboardEvent) => {
    if (e.key === 'Enter' && !e.shiftKey) {
      e.preventDefault();
      sendMessage();
    }
  };

  if (!user) {
    return (
      <div className="min-h-screen bg-gradient-to-br from-gray-900 via-purple-900 to-gray-900 flex items-center justify-center p-4">
        <div className="w-full max-w-md">
          <div className="text-center mb-8">
            <h1 className="text-4xl font-bold text-white mb-2">🚀 MegaVerse</h1>
            <p className="text-gray-400">پیامرسان مدرن</p>
          </div>

          <div className="bg-white/5 backdrop-blur-lg rounded-2xl p-8 border border-white/10">
            <h2 className="text-xl font-bold text-white mb-6">
              {isRegister ? 'ثبت نام' : 'ورود'}
            </h2>

            <input
              type="text"
              placeholder="نام کاربری"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              className="w-full px-4 py-3 bg-white/10 border border-white/20 rounded-xl text-white placeholder-gray-400 mb-4 focus:outline-none focus:border-purple-500"
            />

            <input
              type="password"
              placeholder="رمز عبور"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              className="w-full px-4 py-3 bg-white/10 border border-white/20 rounded-xl text-white placeholder-gray-400 mb-6 focus:outline-none focus:border-purple-500"
              onKeyPress={(e) => e.key === 'Enter' && handleAuth()}
            />

            <button
              onClick={handleAuth}
              className="w-full py-3 bg-gradient-to-r from-purple-500 to-pink-500 rounded-xl font-semibold text-white hover:opacity-90 transition-opacity"
            >
              {isRegister ? 'ثبت نام' : 'ورود'}
            </button>

            <p className="text-center text-gray-400 mt-4">
              {isRegister ? 'حساب دارید؟' : 'حساب ندارید؟'}
              <button
                onClick={() => setIsRegister(!isRegister)}
                className="text-purple-400 hover:text-purple-300 mr-2"
              >
                {isRegister ? 'ورود' : 'ثبت نام'}
              </button>
            </p>
          </div>
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-gradient-to-br from-gray-900 via-purple-900 to-gray-900 flex">
      {/* Sidebar */}
      <div className="w-80 bg-black/30 border-r border-white/10 flex flex-col">
        <div className="p-4 border-b border-white/10">
          <h1 className="text-xl font-bold text-white">🚀 MegaVerse</h1>
          <p className="text-sm text-gray-400">خوش آمدید {user.username}</p>
        </div>

        <div className="flex-1 overflow-y-auto">
          <div className="p-4">
            <h3 className="text-sm font-semibold text-gray-400 mb-3">کاربران آنلاین</h3>
            {users.filter(u => u.id !== user.id).map((u) => (
              <button
                key={u.id}
                onClick={() => startChat(u)}
                className={`w-full text-right p-3 rounded-xl mb-2 transition-colors ${
                  selectedUser?.id === u.id
                    ? 'bg-purple-500/30 text-white'
                    : 'hover:bg-white/5 text-gray-300'
                }`}
              >
                <div className="flex items-center gap-3">
                  <div className="w-10 h-10 rounded-full bg-gradient-to-r from-purple-500 to-pink-500 flex items-center justify-center text-white font-bold">
                    {u.username[0].toUpperCase()}
                  </div>
                  <div>
                    <div className="font-medium">{u.username}</div>
                    <div className="text-xs text-green-400">آنلاین</div>
                  </div>
                </div>
              </button>
            ))}
          </div>
        </div>
      </div>

      {/* Chat Area */}
      <div className="flex-1 flex flex-col">
        {selectedUser ? (
          <>
            {/* Chat Header */}
            <div className="p-4 border-b border-white/10 bg-black/20">
              <div className="flex items-center gap-3">
                <div className="w-10 h-10 rounded-full bg-gradient-to-r from-purple-500 to-pink-500 flex items-center justify-center text-white font-bold">
                  {selectedUser.username[0].toUpperCase()}
                </div>
                <div>
                  <div className="font-semibold text-white">{selectedUser.username}</div>
                  <div className="text-xs text-green-400">آنلاین</div>
                </div>
              </div>
            </div>

            {/* Messages */}
            <div className="flex-1 overflow-y-auto p-4 space-y-4">
              {messages.map((msg) => (
                <div
                  key={msg.id}
                  className={`flex ${msg.sender_id === user.id ? 'justify-end' : 'justify-start'}`}
                >
                  <div
                    className={`max-w-xs lg:max-w-md px-4 py-2 rounded-2xl ${
                      msg.sender_id === user.id
                        ? 'bg-gradient-to-r from-purple-500 to-pink-500 text-white'
                        : 'bg-white/10 text-white'
                    }`}
                  >
                    {msg.sender_id !== user.id && (
                      <div className="text-xs text-purple-300 mb-1">{msg.sender_name}</div>
                    )}
                    <div>{msg.content}</div>
                    <div className="text-xs opacity-60 mt-1">
                      {new Date(msg.created_at).toLocaleTimeString('fa-IR')}
                    </div>
                  </div>
                </div>
              ))}
              <div ref={messagesEndRef} />
            </div>

            {/* Input */}
            <div className="p-4 border-t border-white/10 bg-black/20">
              <div className="flex gap-3">
                <input
                  type="text"
                  value={input}
                  onChange={(e) => setInput(e.target.value)}
                  onKeyPress={handleKeyPress}
                  placeholder="پیام بنویسید..."
                  className="flex-1 px-4 py-3 bg-white/10 border border-white/20 rounded-xl text-white placeholder-gray-400 focus:outline-none focus:border-purple-500"
                />
                <button
                  onClick={sendMessage}
                  className="px-6 py-3 bg-gradient-to-r from-purple-500 to-pink-500 rounded-xl font-semibold text-white hover:opacity-90 transition-opacity"
                >
                  ارسال
                </button>
              </div>
            </div>
          </>
        ) : (
          <div className="flex-1 flex items-center justify-center">
            <div className="text-center text-gray-400">
              <div className="text-6xl mb-4">💬</div>
              <p className="text-xl">یک کاربر را انتخاب کنید</p>
              <p className="text-sm mt-2">برای شروع چت</p>
            </div>
          </div>
        )}
      </div>
    </div>
  );
}
