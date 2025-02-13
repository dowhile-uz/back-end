package userModelFx

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

func (m *Model) GetGithubTokens(ctx context.Context, userID int64) (*GithubTokens, error) {
	var githubTokens GithubTokens

	key := fmt.Sprintf("user-%d-github-tokens", userID)

	status := m.redis.Get(ctx, key)
	if status != nil && status.Err() == nil {
		err := json.Unmarshal([]byte(status.Val()), &githubTokens)
		if err == nil {
			return &githubTokens, nil
		}

		fmt.Println("models.User: redis get unmarshal error")
	}

	err := m.postgres.GetContext(ctx, &githubTokens, githubTokensByUserIDSelectQuery, userID)
	if err != nil {
		return nil, err
	}

	githubTokensRaw, err := json.Marshal(&githubTokens)
	if err != nil {
		return nil, errors.New("models.user: json marshal error")
	}

	now := time.Now()

	cmd := m.redis.Set(ctx, key, string(githubTokensRaw), now.Sub(*githubTokens.ExpiresIn)-time.Minute)
	if cmd != nil && cmd.Err() != nil {
		fmt.Println("models.User: redis set error", cmd.Err())
	}

	return &githubTokens, nil
}

const githubTokensByUserIDSelectQuery = `
SELECT *
FROM github_tokens
WHERE user_id = $1 AND expires_in > CURRENT_TIMESTAMP
ORDER BY expires_in desc
LIMIT 1;`
