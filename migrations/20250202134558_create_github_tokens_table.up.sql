CREATE TABLE IF NOT EXISTS github_tokens(
  id BIGSERIAL PRIMARY KEY,
  user_id BIGINT REFERENCES users (id),
  access_token VARCHAR(256),
  expires_in TIMESTAMP,
  refresh_token VARCHAR(256),
  refresh_token_expires_in TIMESTAMP,
  created_at TIMESTAMP DEFAULT current_timestamp
);


