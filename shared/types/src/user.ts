export interface User {
  id: string;
  email: string;
  name: string;
  avatar?: string;
  role: UserRole;
  active: boolean;
  created_at: Date;
  updated_at: Date;
}

export type UserRole = 'admin' | 'user' | 'moderator';

export interface CreateUserRequest {
  email: string;
  password: string;
  name: string;
}

export interface UpdateUserRequest {
  name?: string;
  avatar?: string;
}

export interface LoginRequest {
  email: string;
  password: string;
}

export interface AuthResponse {
  token: string;
  refresh_token: string;
  expires_at: number;
  user: User;
}
