package githubServiceFx

import (
	"context"
)

/*
  TODO:
  for now just do get /repos/:user/:repo request
  in the future better to rewrite it using internal table with sync feature
  sync can be accomplished by the last updated datetime tracker
  https://stackoverflow.com/questions/18810743/github-api-how-to-check-if-repository-name-available
*/

func (s *Service) CommunityRepoAvailabilityCheck(ctx context.Context, repo string) (bool, error) {
	repoResult, _, err := s.githubClient.Bot.Repositories.Get(ctx, s.config.Github.CommunityOrgName, repo)
	if err != nil {
		// TODO: proper error handling
		return true, nil
	}

	return repoResult == nil, nil
}
