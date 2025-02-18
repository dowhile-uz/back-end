package userModelFx

import (
	"context"

	"github.com/google/go-github/v68/github"
)

var utc = "+05:00"

func (m *Model) CreateOrUpdateUser(ctx context.Context, githubUser *github.User) (*User, error) {
	var existingUser User

	err := m.postgres.GetContext(ctx, &existingUser, "SELECT * FROM users WHERE github_id = $1 LIMIT 1", githubUser.ID)
	if err != nil {
		newUser := User{
			GitHubID:       githubUser.ID,
			Username:       githubUser.Login,
			GitHubUsername: githubUser.Login, // TODO: if github username changes, we need a way to resolve conflict
			AvatarURL:      githubUser.AvatarURL,
			Email:          githubUser.Email,
			Name:           githubUser.Name,
			Timezone:       &utc,
		}

		_, err = m.postgres.NamedExecContext(ctx, userInsertQuery, &newUser)
		if err != nil {
			return nil, err
		}

		err = m.postgres.GetContext(ctx, &existingUser, "SELECT * FROM users WHERE github_id = $1 LIMIT 1", githubUser.ID)
		if err != nil {
			return nil, err
		}
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
