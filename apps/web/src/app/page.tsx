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
  const [selectedUser, setSelectedUser] = useState<User | null>(null);
  const [currentConvID, setCurrentConvID] = useState('');
  const [loaded, setLoaded] = useState(false);
  const [showSidebar, setShowSidebar] = useState(true);
  const messagesEndRef = useRef<HTMLDivElement>(null);
  const inputRef = useRef<HTMLInputElement>(null);

  useEffect(() => {
    setLoaded(true);
    if (user) loadUsers();
  }, [user]);

  useEffect(() => {
    messagesEndRef.current?.scrollIntoView({ behavior: 'smooth' });
  }, [messages]);

  useEffect(() => {
    if (currentConvID) loadMessages();
  }, [currentConvID]);

  const loadUsers = async () => {
    const res = await fetch('/api/users');
    if (res.ok) {
      const data = await res.json();
      setUsers(data.users || []);
    }
  };

  const loadMessages = async () => {
    const res = await fetch(`/api/messages?conversation_id=${currentConvID}`);
    if (res.ok) {
      const data = await res.json();
      setMessages(data.messages || []);
    }
  };

  const handleAuth = async () => {
    const endpoint = isRegister ? '/api/register' : '/api/login';
    const res = await fetch(endpoint, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ username, password }),
    });

    if (res.ok) {
      const data = await res.json();
      setUser(data);
    } else {
      const err = await res.json();
      alert(err.error || 'خطا');
    }
  };

  const startChat = (targetUser: User) => {
    setSelectedUser(targetUser);
    const sorted = [user?.id, targetUser.id].sort();
    setCurrentConvID(`dm_${sorted[0]}_${sorted[1]}`);
    setTimeout(() => inputRef.current?.focus(), 100);
  };

  const sendMessage = async () => {
    if (!input.trim() || !selectedUser || !user) return;

    const sorted = [user.id, selectedUser.id].sort();
    const convID = `dm_${sorted[0]}_${sorted[1]}`;

    const res = await fetch('/api/messages', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ conversation_id: convID, sender_id: user.id, content: input }),
    });

    if (res.ok) {
      const msg = await res.json();
      setMessages(prev => [...prev, msg]);
      setInput('');
    }
  };

  const handleKeyPress = (e: React.KeyboardEvent) => {
    if (e.key === 'Enter' && !e.shiftKey) {
      e.preventDefault();
      sendMessage();
    }
  };

  const formatTime = (dateStr: string) => {
    return new Date(dateStr).toLocaleTimeString('fa-IR', { hour: '2-digit', minute: '2-digit' });
  };

  // Login Page
  if (!user) {
    return (
      <div className="min-h-screen bg-mesh bg-grid flex items-center justify-center p-4 relative overflow-hidden">
        {/* Animated background orbs */}
        <div className="absolute inset-0 pointer-events-none">
          <div className="absolute top-20 right-20 w-72 h-72 bg-purple-500/20 rounded-full blur-3xl animate-float-slow" />
          <div className="absolute bottom-20 left-20 w-96 h-96 bg-pink-500/15 rounded-full blur-3xl animate-float-slow" style={{ animationDelay: '2s' }} />
          <div className="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 w-[600px] h-[600px] bg-cyan-500/10 rounded-full blur-3xl animate-pulse-glow" />
        </div>

        <div className={`relative z-10 w-full max-w-md transition-all duration-1000 ${loaded ? 'opacity-100 translate-y-0' : 'opacity-0 translate-y-10'}`}>
          {/* Logo */}
          <div className="text-center mb-10 animate-slide-up">
            <div className="inline-block mb-6">
              <div className="w-24 h-24 rounded-3xl bg-gradient-to-br from-purple-500 via-pink-500 to-cyan-500 flex items-center justify-center text-5xl glow-purple animate-float">
                🚀
              </div>
            </div>
            <h1 className="text-5xl font-black mb-3 gradient-text-animated">MegaVerse</h1>
            <p className="text-gray-400 text-lg">پیامرسان مدرن و سریع</p>
          </div>

          {/* Form */}
          <div className="glass-strong rounded-3xl p-8 animate-slide-up stagger-2">
            <h2 className="text-2xl font-bold text-white mb-8 text-center">
              {isRegister ? '✨ ساخت حساب جدید' : '👋 خوش آمدید'}
            </h2>

            <div className="space-y-5">
              <div className="relative group">
                <input
                  type="text"
                  placeholder="نام کاربری"
                  value={username}
                  onChange={(e) => setUsername(e.target.value)}
                  className="w-full px-5 py-4 bg-white/5 border border-white/10 rounded-2xl text-white placeholder-gray-500 focus:outline-none focus:border-purple-500 focus:bg-white/8 transition-all duration-300 input-focus"
                />
                <div className="absolute left-4 top-1/2 -translate-y-1/2 text-gray-500 group-focus-within:text-purple-400 transition-colors">
                  👤
                </div>
              </div>

              <div className="relative group">
                <input
                  type="password"
                  placeholder="رمز عبور"
                  value={password}
                  onChange={(e) => setPassword(e.target.value)}
                  onKeyPress={(e) => e.key === 'Enter' && handleAuth()}
                  className="w-full px-5 py-4 bg-white/5 border border-white/10 rounded-2xl text-white placeholder-gray-500 focus:outline-none focus:border-purple-500 focus:bg-white/8 transition-all duration-300 input-focus"
                />
                <div className="absolute left-4 top-1/2 -translate-y-1/2 text-gray-500 group-focus-within:text-purple-400 transition-colors">
                  🔒
                </div>
              </div>

              <button
                onClick={handleAuth}
                className="w-full py-4 bg-gradient-to-r from-purple-500 via-pink-500 to-cyan-500 rounded-2xl font-bold text-lg text-white hover:opacity-90 transition-all duration-300 btn-glow glow-purple"
              >
                {isRegister ? '🚀 ثبت نام' : '🎯 ورود'}
              </button>
            </div>

            <div className="mt-6 text-center">
              <button
                onClick={() => setIsRegister(!isRegister)}
                className="text-purple-400 hover:text-purple-300 transition-colors text-sm"
              >
                {isRegister ? 'حساب دارید؟ ورود کنید' : 'حساب ندارید؟ ثبت نام کنید'}
              </button>
            </div>

            {/* Demo hint */}
            {!isRegister && (
              <div className="mt-6 p-4 bg-purple-500/10 rounded-xl border border-purple-500/20">
                <p className="text-xs text-purple-300 text-center">
                  💡 کاربران تستی: ali, reza, sara, mina, hassan
                </p>
                <p className="text-xs text-purple-400 text-center mt-1">
                  رمز: 123456
                </p>
              </div>
            )}
          </div>
        </div>
      </div>
    );
  }

  // Messenger
  return (
    <div className="h-screen flex bg-mesh bg-grid relative overflow-hidden">
      {/* Background effects */}
      <div className="absolute inset-0 pointer-events-none">
        <div className="absolute top-0 right-1/4 w-96 h-96 bg-purple-500/10 rounded-full blur-3xl animate-pulse-glow" />
        <div className="absolute bottom-0 left-1/4 w-96 h-96 bg-pink-500/10 rounded-full blur-3xl animate-pulse-glow" style={{ animationDelay: '3s' }} />
      </div>

      {/* Sidebar */}
      <div className={`relative z-10 ${showSidebar ? 'w-80' : 'w-0'} transition-all duration-500 glass border-l border-white/10 flex flex-col`}>
        {/* Header */}
        <div className="p-5 border-b border-white/10">
          <div className="flex items-center justify-between">
            <div className="flex items-center gap-3">
              <div className="w-12 h-12 rounded-2xl bg-gradient-to-br from-purple-500 via-pink-500 to-cyan-500 flex items-center justify-center text-2xl glow-purple animate-float">
                🚀
              </div>
              <div>
                <h1 className="text-lg font-black gradient-text-animated">MegaVerse</h1>
                <p className="text-xs text-gray-500">{user.username}</p>
              </div>
            </div>
            <div className="online-dot" />
          </div>
        </div>

        {/* Search */}
        <div className="p-4">
          <div className="relative">
            <input
              type="text"
              placeholder="جستجو..."
              className="w-full px-4 py-3 bg-white/5 border border-white/10 rounded-xl text-white placeholder-gray-500 text-sm focus:outline-none focus:border-purple-500 transition-all duration-300 input-focus pr-10"
            />
            <span className="absolute right-3 top-1/2 -translate-y-1/2 text-gray-500">🔍</span>
          </div>
        </div>

        {/* Users List */}
        <div className="flex-1 overflow-y-auto px-3">
          <p className="text-xs text-gray-500 px-2 mb-3 font-medium">کاربران آنلاین</p>
          {users.filter(u => u.id !== user.id).map((u, i) => (
            <button
              key={u.id}
              onClick={() => startChat(u)}
              className={`w-full text-right p-3 rounded-2xl mb-2 transition-all duration-300 animate-slide-in-right ${
                selectedUser?.id === u.id
                  ? 'bg-gradient-to-l from-purple-500/30 to-pink-500/30 text-white glow-purple'
                  : 'hover:bg-white/5 text-gray-300'
              }`}
              style={{ animationDelay: `${i * 0.05}s` }}
            >
              <div className="flex items-center gap-3">
                <div className="relative">
                  <div className="w-12 h-12 rounded-2xl bg-gradient-to-br from-purple-500 via-pink-500 to-cyan-500 flex items-center justify-center text-white font-bold text-lg">
                    {u.username[0].toUpperCase()}
                  </div>
                  <div className="absolute -bottom-1 -left-1 online-dot" />
                </div>
                <div className="flex-1">
                  <div className="font-bold">{u.username}</div>
                  <div className="text-xs text-green-400 flex items-center gap-1">
                    <span className="w-1.5 h-1.5 bg-green-400 rounded-full animate-pulse-glow" />
                    آنلاین
                  </div>
                </div>
              </div>
            </button>
          ))}
          {users.filter(u => u.id !== user.id).length === 0 && (
            <div className="text-center py-8 text-gray-500 animate-float-slow">
              <div className="text-4xl mb-3">👥</div>
              <p className="text-sm">هنوز کاربری ثبت نام نکرده</p>
            </div>
          )}
        </div>

        {/* Footer */}
        <div className="p-4 border-t border-white/10">
          <button
            onClick={() => { setUser(null); setMessages([]); setSelectedUser(null); }}
            className="w-full py-3 text-gray-400 hover:text-red-400 hover:bg-red-500/10 rounded-xl transition-all duration-300 text-sm font-medium"
          >
            🚪 خروج
          </button>
        </div>
      </div>

      {/* Chat Area */}
      <div className="flex-1 flex flex-col relative z-10">
        {selectedUser ? (
          <>
            {/* Chat Header */}
            <div className="glass border-b border-white/10 p-4 animate-slide-up">
              <div className="flex items-center gap-4">
                <button
                  onClick={() => setShowSidebar(!showSidebar)}
                  className="lg:hidden p-2 hover:bg-white/10 rounded-xl transition-colors"
                >
                  ☰
                </button>
                <div className="relative">
                  <div className="w-12 h-12 rounded-2xl bg-gradient-to-br from-purple-500 via-pink-500 to-cyan-500 flex items-center justify-center text-white font-bold text-lg">
                    {selectedUser.username[0].toUpperCase()}
                  </div>
                  <div className="absolute -bottom-1 -left-1 online-dot" />
                </div>
                <div className="flex-1">
                  <h3 className="font-bold text-lg">{selectedUser.username}</h3>
                  <p className="text-xs text-green-400 flex items-center gap-1">
                    <span className="w-1.5 h-1.5 bg-green-400 rounded-full animate-pulse" />
                    آنلاین • در حال حاضر
                  </p>
                </div>
                <div className="flex gap-2">
                  <button className="p-3 hover:bg-white/10 rounded-xl transition-all duration-300 text-lg">📞</button>
                  <button className="p-3 hover:bg-white/10 rounded-xl transition-all duration-300 text-lg">📹</button>
                  <button className="p-3 hover:bg-white/10 rounded-xl transition-all duration-300 text-lg">⋮</button>
                </div>
              </div>
            </div>

            {/* Messages */}
            <div className="flex-1 overflow-y-auto p-6 space-y-4">
              {messages.length === 0 && (
                <div className="flex flex-col items-center justify-center h-full text-gray-500 animate-scale-in">
                  <div className="text-6xl mb-4 animate-float">💬</div>
                  <p className="text-xl font-bold mb-2">شروع مکالمه</p>
                  <p className="text-sm">اولین پیام را برای {selectedUser.username} ارسال کنید</p>
                </div>
              )}
              {messages.map((msg, i) => (
                <div
                  key={msg.id}
                  className={`flex ${msg.sender_id === user.id ? 'justify-end' : 'justify-start'} animate-message-in`}
                  style={{ animationDelay: `${i * 0.05}s` }}
                >
                  <div
                    className={`max-w-md px-5 py-3 ${
                      msg.sender_id === user.id
                        ? 'message-sent text-white'
                        : 'message-received text-white'
                    }`}
                  >
                    {msg.sender_id !== user.id && (
                      <div className="text-xs text-purple-300 mb-1 font-bold">{msg.sender_name}</div>
                    )}
                    <div className="text-sm leading-relaxed">{msg.content}</div>
                    <div className="text-xs opacity-50 mt-1 text-left">{formatTime(msg.created_at)}</div>
                  </div>
                </div>
              ))}
              <div ref={messagesEndRef} />
            </div>

            {/* Input */}
            <div className="glass border-t border-white/10 p-4 animate-slide-up">
              <div className="flex gap-3 items-center">
                <button className="p-3 hover:bg-white/10 rounded-xl transition-all duration-300 text-lg">😊</button>
                <button className="p-3 hover:bg-white/10 rounded-xl transition-all duration-300 text-lg">📎</button>
                <input
                  ref={inputRef}
                  type="text"
                  value={input}
                  onChange={(e) => setInput(e.target.value)}
                  onKeyPress={handleKeyPress}
                  placeholder="پیام بنویسید..."
                  className="flex-1 px-5 py-3 bg-white/5 border border-white/10 rounded-2xl text-white placeholder-gray-500 focus:outline-none focus:border-purple-500 transition-all duration-300 input-focus"
                />
                <button
                  onClick={sendMessage}
                  disabled={!input.trim()}
                  className="p-3 bg-gradient-to-r from-purple-500 to-pink-500 rounded-2xl text-white hover:opacity-90 transition-all duration-300 btn-glow glow-purple disabled:opacity-30 disabled:cursor-not-allowed"
                >
                  <svg className="w-6 h-6 rotate-180" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 19l9 2-9-18-9 18 9-2zm0 0v-8" />
                  </svg>
                </button>
              </div>
            </div>
          </>
        ) : (
          <div className="flex-1 flex items-center justify-center">
            <div className="text-center animate-scale-in">
              <div className="text-8xl mb-6 animate-float">💬</div>
              <h2 className="text-3xl font-black gradient-text-animated mb-3">MegaVerse Messenger</h2>
              <p className="text-gray-400 text-lg">یک کاربر را انتخاب کنید تا شروع کنید</p>
              <div className="mt-6 flex justify-center gap-4">
                <div className="glass px-4 py-2 rounded-xl text-sm text-gray-400 animate-slide-in-right stagger-1">
                  ⚡ سریع و امن
                </div>
                <div className="glass px-4 py-2 rounded-xl text-sm text-gray-400 animate-slide-in-right stagger-2">
                  🔄 Real-time
                </div>
                <div className="glass px-4 py-2 rounded-xl text-sm text-gray-400 animate-slide-in-right stagger-3">
                  🎨 زیبا و مدرن
                </div>
              </div>
            </div>
          </div>
        )}
      </div>
    </div>
  );
}
