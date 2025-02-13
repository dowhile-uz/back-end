package editorServiceFx

import (
	"context"

	"github.com/google/go-github/v68/github"
)

var publicRepoVisibility = "public"

func (s *Service) CreateCommunityRepo(ctx context.Context, repo string) (*github.Repository, error) {
	repository := &github.Repository{
		Name:       &repo,
		Visibility: &publicRepoVisibility,
	}

	repoResult, _, err := s.githubClient.Bot.Repositories.Create(ctx, s.config.Github.CommunityOrgName, repository)
	if err != nil {
		return nil, err
	}

	return repoResult, nil
}
