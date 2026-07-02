export interface Config {
  nodeEnv: string;
  appName: string;
  appVersion: string;
  logLevel: string;
  database: DatabaseConfig;
  redis: RedisConfig;
  jwt: JWTConfig;
  gateway: GatewayConfig;
}

export interface DatabaseConfig {
  host: string;
  port: number;
  name: string;
  user: string;
  password: string;
}

export interface RedisConfig {
  host: string;
  port: number;
  password: string;
  db: number;
}

export interface JWTConfig {
  secret: string;
  expiration: number;
  refreshExpiration: number;
}

export interface GatewayConfig {
  port: number;
  rateLimit: number;
  rateWindow: number;
}

function env(key: string, defaultValue: string = ''): string {
  return process.env[key] || defaultValue;
}

function envInt(key: string, defaultValue: number): number {
  const value = process.env[key];
  return value ? parseInt(value, 10) : defaultValue;
}

export function loadConfig(): Config {
  return {
    nodeEnv: env('NODE_ENV', 'development'),
    appName: env('APP_NAME', 'megaverse'),
    appVersion: env('APP_VERSION', '0.1.0'),
    logLevel: env('LOG_LEVEL', 'info'),
    database: {
      host: env('DATABASE_HOST', 'localhost'),
      port: envInt('DATABASE_PORT', 3306),
      name: env('DATABASE_NAME', 'megaverse'),
      user: env('DATABASE_USER', 'megaverse'),
      password: env('DATABASE_PASSWORD', 'password'),
    },
    redis: {
      host: env('REDIS_HOST', 'localhost'),
      port: envInt('REDIS_PORT', 6379),
      password: env('REDIS_PASSWORD', ''),
      db: envInt('REDIS_DB', 0),
    },
    jwt: {
      secret: env('JWT_SECRET', 'dev-secret-change-in-production'),
      expiration: envInt('JWT_EXPIRATION', 86400),
      refreshExpiration: envInt('JWT_REFRESH_EXPIRATION', 604800),
    },
    gateway: {
      port: envInt('GATEWAY_PORT', 8080),
      rateLimit: envInt('GATEWAY_RATE_LIMIT', 1000),
      rateWindow: envInt('GATEWAY_RATE_WINDOW', 60),
    },
  };
}
