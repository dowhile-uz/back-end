package userModelFx

import (
	"context"
	"fmt"
	"time"
)

func (m *Model) CreateGithubTokens(ctx context.Context, userID int64, tokens *GithubAuthTokensResponse) (*GithubTokens, error) {
	key := fmt.Sprintf("user-%d-github-access-token", userID)

	now := time.Now()
	expiresIn := now.Add(time.Second * time.Duration(tokens.ExpiresIn))
	refreshTokenExpiresIn := now.Add(time.Second * time.Duration(tokens.RefreshTokenExpiresIn))

	githubTokens := GithubTokens{
		UserID:                &userID,
		AccessToken:           &tokens.AccessToken,
		ExpiresIn:             &expiresIn,
		RefreshToken:          &tokens.RefreshToken,
		RefreshTokenExpiresIn: &refreshTokenExpiresIn,
	}

	_, err := m.postgres.NamedExecContext(ctx, githubTokensInsertQuery, &githubTokens)
	if err != nil {
		return nil, err
	}

	redisExpireTime := githubTokens.ExpiresIn.Sub(now) - time.Minute

	cmd := m.redis.Set(ctx, key, *githubTokens.AccessToken, redisExpireTime)
	if cmd != nil && cmd.Err() != nil {
		fmt.Println("models.user.create-github-tokens: redis set error", cmd.Err())
	}

	return &githubTokens, nil
}

const githubTokensInsertQuery = `
INSERT INTO github_tokens
(
  user_id,
  access_token,
  expires_in,
  refresh_token,
  refresh_token_expires_in
) VALUES (
  :user_id,
  :access_token,
  :expires_in,
  :refresh_token,
  :refresh_token_expires_in
)`
