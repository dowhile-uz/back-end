package userModelFx

import (
	"context"
	"errors"
	"time"

	"github.com/google/go-github/v68/github"
)

var utc = "+05:00"

func (m *Model) CreateOrUpdateWithTokens(ctx context.Context, user *github.User, tokens *GithubAuthTokensResponse) (*User, error) {
	if user == nil || tokens == nil {
		return nil, errors.New("user or tokens are not provided")
	}

	var existingUser User

	err := m.postgres.GetContext(ctx, &existingUser, "SELECT * FROM users WHERE github_id = $1 LIMIT 1", user.ID)
	if err != nil {
		newUser := User{
			GitHubID:       user.ID,
			Username:       user.Login,
			GitHubUsername: user.Login, // TODO: if github username changes, we need a way to resolve conflict
			AvatarURL:      user.AvatarURL,
			Email:          user.Email,
			Name:           user.Name,
			Timezone:       &utc,
		}

		_, err = m.postgres.NamedExecContext(ctx, userInsertQuery, &newUser)
		if err != nil {
			return nil, err
		}

		err = m.postgres.GetContext(ctx, &existingUser, "SELECT * FROM users WHERE github_id = $1 LIMIT 1", user.ID)
		if err != nil {
			return nil, err
		}
	}

	now := time.Now()
	expiresIn := now.Add(time.Second * time.Duration(tokens.ExpiresIn))
	refreshTokenExpiresIn := now.Add(time.Second * time.Duration(tokens.RefreshTokenExpiresIn))

	githubTokens := GithubTokens{
		UserID:                existingUser.ID,
		AccessToken:           &tokens.AccessToken,
		ExpiresIn:             &expiresIn,
		RefreshToken:          &tokens.RefreshToken,
		RefreshTokenExpiresIn: &refreshTokenExpiresIn,
	}

	_, err = m.postgres.NamedExecContext(ctx, githubTokensInsertQuery, &githubTokens)
	if err != nil {
		return nil, err
	}

	return &existingUser, nil
}

const userInsertQuery = `
INSERT INTO users
(
	github_id,
	username,
	github_username,
	avatar_url,
	email,
	name,
	timezone
) VALUES (
	:github_id,
	:username,
	:github_username,
	:avatar_url,
	:email,
	:name,
	:timezone
)`

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
