// Service names
export const SERVICES = {
  API_GATEWAY: 'api-gateway',
  AUTH_SERVICE: 'auth-service',
  USER_SERVICE: 'user-service',
  SOCIAL_SERVICE: 'social-service',
  MESSAGING_SERVICE: 'messaging-service',
  AI_SERVICE: 'ai-service',
} as const;

// Default ports
export const PORTS = {
  API_GATEWAY: 8080,
  AUTH_SERVICE: 8081,
  USER_SERVICE: 8082,
  SOCIAL_SERVICE: 8083,
  MESSAGING_SERVICE: 8084,
  AI_SERVICE: 8085,
} as const;

// Database defaults
export const DATABASE = {
  HOST: 'localhost',
  PORT: 3306,
  NAME: 'megaverse',
  USER: 'megaverse',
  PASSWORD: 'password',
} as const;

// Redis defaults
export const REDIS = {
  HOST: 'localhost',
  PORT: 6379,
  DB: 0,
} as const;

// JWT defaults
export const JWT = {
  EXPIRATION: 86400,        // 24 hours
  REFRESH_EXPIRATION: 604800, // 7 days
} as const;

// Rate limiting defaults
export const RATE_LIMIT = {
  MAX_REQUESTS: 1000,
  WINDOW_SECONDS: 60,
} as const;

// User roles
export const ROLES = {
  ADMIN: 'admin',
  USER: 'user',
  MODERATOR: 'moderator',
} as const;

export type Role = typeof ROLES[keyof typeof ROLES];
