package githubServiceFx

import (
	"context"

	"github.com/google/go-github/v68/github"
)

func (s *Service) ListInstallationRepos(ctx context.Context, accessToken string) ([]*github.Repository, error) {
	installations, err := s.ListInstallations(ctx, accessToken)
	if err != nil {
		return nil, err
	}

	var repositories []*github.Repository

	for _, installation := range installations {
		repos, _, err := s.githubClient.InstallationClient(*installation.ID).Apps.ListRepos(ctx, &github.ListOptions{
			Page:    1,
			PerPage: 100,
		})

		if err != nil {
			return nil, err
		}

		repositories = append(repositories, repos.Repositories...)
	}

	return repositories, nil
}
