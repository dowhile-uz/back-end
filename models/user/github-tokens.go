package userModelFx

import (
	"time"
)

type GithubTokens struct {
	ID                    *int64     `db:"id" json:"id"`
	UserID                *int64     `db:"user_id" json:"user_id"`
	AccessToken           *string    `db:"access_token" json:"access_token"`
	ExpiresIn             *time.Time `db:"expires_in" json:"expires_in"`
	RefreshToken          *string    `db:"refresh_token" json:"refresh_token"`
	RefreshTokenExpiresIn *time.Time `db:"refresh_token_expires_in" json:"refresh_token_expires_in"`
	CreatedAt             *time.Time `db:"created_at" json:"created_at"`
}
