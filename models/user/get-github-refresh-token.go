package userModelFx

import (
	"context"
)

func (m *Model) GetGithubRefreshToken(ctx context.Context, userID int64) (string, error) {
	var githubTokens GithubTokens

	err := m.postgres.GetContext(ctx, &githubTokens, githubRefreshTokenByUserIDSelectQuery, userID)
	if err != nil {
		return "", err
	}

	return *githubTokens.RefreshToken, nil
}

const githubRefreshTokenByUserIDSelectQuery = `
SELECT *
FROM github_tokens
WHERE user_id = $1 AND refresh_token_expires_in > CURRENT_TIMESTAMP
ORDER BY refresh_token_expires_in desc
LIMIT 1;`
