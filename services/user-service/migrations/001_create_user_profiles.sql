-- MegaVerse User Service Migration
-- 001_create_user_profiles.sql

CREATE TABLE IF NOT EXISTS user_profiles (
    user_id VARCHAR(36) PRIMARY KEY,
    bio TEXT,
    website VARCHAR(255),
    location VARCHAR(255),
    avatar_url VARCHAR(500),
    cover_url VARCHAR(500),
    followers_count INT DEFAULT 0,
    following_count INT DEFAULT 0,
    posts_count INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_profiles_location (location)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS user_follows (
    follower_id VARCHAR(36) NOT NULL,
    following_id VARCHAR(36) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (follower_id, following_id),
    INDEX idx_follows_following (following_id),
    FOREIGN KEY (follower_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (following_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
