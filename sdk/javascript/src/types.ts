export interface MegaVerseConfig {
  apiKey?: string;
  baseUrl: string;
}

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

export interface APIResponse<T> {
  success: boolean;
  data?: T;
  error?: { code: string; message: string };
}

export interface PaginationParams {
  page?: number;
  limit?: number;
}
