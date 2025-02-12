package userModelFx

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

func (m *Model) GetUser(ctx context.Context, userID int64) (*User, error) {
	var user User

	key := fmt.Sprintf("user-%d", userID)

	status := m.redis.Get(ctx, key)
	if status != nil && status.Err() == nil {
		err := json.Unmarshal([]byte(status.Val()), &user)
		if err == nil {
			return &user, nil
		}

		fmt.Println("models.User: redis get unmarshal error")
	}

	err := m.postgres.GetContext(ctx, &user, userByIDSelectQuery, userID)
	if err != nil {
		return nil, err
	}

	userRaw, err := json.Marshal(&user)
	if err != nil {
		return nil, errors.New("models.user: json marshal error")
	}

	cmd := m.redis.Set(ctx, key, string(userRaw), time.Duration(5)*time.Minute)
	if cmd != nil && cmd.Err() != nil {
		fmt.Println("models.User: redis set error", cmd.Err())
	}

	return &user, nil
}

const userByIDSelectQuery = `
SELECT *
FROM users
WHERE id = $1`
