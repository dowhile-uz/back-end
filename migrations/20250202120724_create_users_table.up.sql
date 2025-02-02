CREATE TABLE IF NOT EXISTS users(
  id BIGSERIAL PRIMARY KEY,
  github_id BIGINT UNIQUE,
  username VARCHAR(128) UNIQUE,
  github_username VARCHAR(128),
  avatar_url VARCHAR,
  email VARCHAR(256) UNIQUE,
  name VARCHAR(256),
  timezone VARCHAR(128) DEFAULT '+05:00',
  created_at TIMESTAMP DEFAULT current_timestamp,
  updated_at TIMESTAMP DEFAULT current_timestamp
);
