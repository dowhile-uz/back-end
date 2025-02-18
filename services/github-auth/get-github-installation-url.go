package githubAuthServiceFx

import (
	"fmt"
)

func (s *Service) GetGithubInstallationURL() string {
	return fmt.Sprintf("https://github.com/apps/%s/installations/new", s.config.GithubAuth.AppName)
}
