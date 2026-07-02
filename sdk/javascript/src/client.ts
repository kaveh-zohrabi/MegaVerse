import axios, { AxiosInstance } from 'axios';
import { MegaVerseConfig, User, Post, AuthResponse, PaginationParams } from './types';

export class MegaVerseClient {
  private client: AxiosInstance;
  private token: string | null = null;

  constructor(config: MegaVerseConfig) {
    this.client = axios.create({
      baseURL: config.baseUrl,
      headers: {
        'Content-Type': 'application/json',
        ...(config.apiKey && { 'X-API-Key': config.apiKey }),
      },
    });
  }

  setToken(token: string) {
    this.token = token;
    this.client.defaults.headers.common['Authorization'] = `Bearer ${token}`;
  }

  // Auth
  async register(email: string, name: string, password: string): Promise<User> {
    const { data } = await this.client.post<User>('/auth/register', { email, name, password });
    return data;
  }

  async login(email: string, password: string): Promise<AuthResponse> {
    const { data } = await this.client.post<AuthResponse>('/auth/login', { email, password });
    this.setToken(data.access_token);
    return data;
  }

  // Users
  async getUser(id: string): Promise<User> {
    const { data } = await this.client.get<User>(`/users/${id}`);
    return data;
  }

  async updateProfile(updates: Partial<User>): Promise<{ status: string }> {
    const { data } = await this.client.put<{ status: string }>('/users/me', updates);
    return data;
  }

  async follow(userId: string): Promise<{ status: string }> {
    const { data } = await this.client.post<{ status: string }>(`/users/${userId}/follow`);
    return data;
  }

  async unfollow(userId: string): Promise<{ status: string }> {
    const { data } = await this.client.post<{ status: string }>(`/users/${userId}/unfollow`);
    return data;
  }

  // Posts
  async createPost(content: string, mediaUrl?: string): Promise<Post> {
    const { data } = await this.client.post<Post>('/posts', { content, media_url: mediaUrl });
    return data;
  }

  async getPost(id: string): Promise<Post> {
    const { data } = await this.client.get<Post>(`/posts/${id}`);
    return data;
  }

  async deletePost(id: string): Promise<void> {
    await this.client.delete(`/posts/${id}`);
  }

  async getFeed(params?: PaginationParams): Promise<{ posts: Post[] }> {
    const { data } = await this.client.get<{ posts: Post[] }>('/feed', { params });
    return data;
  }

  async getUserPosts(userId: string, params?: PaginationParams): Promise<{ posts: Post[] }> {
    const { data } = await this.client.get<{ posts: Post[] }>(`/users/${userId}/posts`, { params });
    return data;
  }

  // Health
  async health(): Promise<{ status: string; service: string }> {
    const { data } = await this.client.get<{ status: string; service: string }>('/health');
    return data;
  }
}

export default MegaVerseClient;
