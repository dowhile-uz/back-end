package githubServiceFx

import (
	"context"

	"github.com/google/go-github/v68/github"
)

func (s *Service) ListInstallations(ctx context.Context, accessToken string) ([]*github.Installation, error) {
	installations, _, err := s.githubClient.WithAuthToken(accessToken).Apps.ListUserInstallations(ctx, &github.ListOptions{
		Page:    1,
		PerPage: 100,
	})
	if err != nil {
		return nil, err
	}

	return installations, nil
}
