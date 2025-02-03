package usermodelfx

import "time"

type User struct {
	ID             *int64     `db:"id" json:"id,omitempty" doc:"User unique identifier"`
	GitHubID       *int64     `db:"github_id" json:"-"`
	Username       *string    `db:"username" json:"username,omitempty" doc:"User username" example:"v4l1j0n"`
	GitHubUsername *string    `db:"github_username" json:"github_username,omitempty" doc:"User username in GitHub" example:"v4l1j0n"`
	AvatarURL      *string    `db:"avatar_url" json:"avatar_url,omitempty" doc:"User avatar URL" example:"https://avatars.githubusercontent.com/u/196029471?s=200&v=4"`
	Email          *string    `db:"email" json:"email,omitempty" doc:"User email" example:"valijon@mailservice.tld"`
	Name           *string    `db:"name" json:"name,omitempty" doc:"User name" example:"Valijon"`
	Timezone       *string    `db:"timezone" json:"timezone,omitempty" doc:"User timezone" example:"+05:00"`
	CreatedAt      *time.Time `db:"created_at" json:"created_at,omitempty"`
	UpdatedAt      *time.Time `db:"updated_at" json:"updated_at,omitempty"`
}
