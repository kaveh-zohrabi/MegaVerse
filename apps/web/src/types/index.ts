import { User, Post } from '@/lib/api';

export interface AuthState {
  user: User | null;
  token: string | null;
  isAuthenticated: boolean;
}

export interface FeedState {
  posts: Post[];
  loading: boolean;
  error: string | null;
}
