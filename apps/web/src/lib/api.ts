import axios from 'axios';

const API_BASE = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';

const api = axios.create({
  baseURL: API_BASE,
  headers: {
    'Content-Type': 'application/json',
  },
});

api.interceptors.request.use((config) => {
  if (typeof window !== 'undefined') {
    const token = localStorage.getItem('token');
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
  }
  return config;
});

export interface User {
  id: string;
  email: string;
  name: string;
  role: string;
}

export interface Post {
  id: string;
  user_id: string;
  content: string;
  media_url?: string;
  likes_count: number;
  comments_count: number;
  created_at: string;
}

export interface AuthResponse {
  access_token: string;
  refresh_token: string;
  expires_at: number;
}

export const authApi = {
  login: (email: string, password: string) =>
    api.post<AuthResponse>('/auth/login', { email, password }),
  
  register: (email: string, name: string, password: string) =>
    api.post<User>('/auth/register', { email, name, password }),
};

export const userApi = {
  getProfile: (id: string) =>
    api.get(`/users/${id}`),
  
  updateProfile: (data: Partial<User>) =>
    api.put('/users/me', data),
};

export const postApi = {
  getFeed: () =>
    api.get<{ posts: Post[] }>('/feed'),
  
  createPost: (content: string) =>
    api.post<Post>('/posts', { content }),
  
  getPost: (id: string) =>
    api.get<Post>(`/posts/${id}`),
};

export default api;
