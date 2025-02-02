package usermodelfx

import "time"

type User struct {
	ID             *int64     `db:"id" json:"id"`
	GitHubID       *int64     `db:"github_id"`
	Username       *string    `db:"username" json:"username"`
	GitHubUsername *string    `db:"github_username" json:"github_username"`
	AvatarURL      *string    `db:"avatar_url" json:"avatar_url"`
	Email          *string    `db:"email" json:"email"`
	Name           *string    `db:"name" json:"name"`
	Timezone       *string    `db:"timezone" json:"timezone"`
	CreatedAt      *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt      *time.Time `db:"updated_at" json:"updated_at"`
}
