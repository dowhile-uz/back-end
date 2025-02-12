package githubAuthServiceFx

import (
	"context"

	"github.com/google/go-github/v68/github"
)

func (s *Service) FetchUserData(ctx context.Context, accessToken string) (*github.User, error) {
	user, _, err := s.githubClient.WithAuthToken(accessToken).Users.Get(ctx, "")
	if err != nil {
		return nil, err
	}

	return user, nil
}
