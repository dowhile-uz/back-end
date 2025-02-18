package userModelFx

import (
	"context"
	"fmt"
	"time"
)

func (m *Model) GetGithubAccessToken(ctx context.Context, userID int64) (string, error) {
	key := fmt.Sprintf("user-%d-github-access-token", userID)

	status := m.redis.Get(ctx, key)
	if status != nil && status.Err() == nil {
		return status.Val(), nil
	}

	var githubTokens GithubTokens

	err := m.postgres.GetContext(ctx, &githubTokens, githubTokensByUserIDSelectQuery, userID)
	if err != nil {
		return "", err // TODO: custom error
	}

	now := time.Now()
	redisExpireTime := githubTokens.ExpiresIn.Sub(now) - time.Minute

	cmd := m.redis.Set(ctx, key, *githubTokens.AccessToken, redisExpireTime)
	if cmd != nil && cmd.Err() != nil {
		fmt.Println("models.user.get-github-access-token: redis set error", cmd.Err())
	}

	return *githubTokens.AccessToken, nil
}

const githubTokensByUserIDSelectQuery = `
SELECT *
FROM github_tokens
WHERE user_id = $1 AND expires_in > CURRENT_TIMESTAMP
ORDER BY expires_in desc
LIMIT 1;`
