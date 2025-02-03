package githubauthservicefx

import (
	"fmt"
	"net/url"
	"strings"

	"dowhile.uz/back-end/utils"
)

func (s *Service) GetGitHubRedirectURL(back string) string {
	query := url.Values{
		"client_id": {s.config.GithubAuth.ClientID},
		"state":     {utils.RandomBase16String(6)},
		"scope":     {url.QueryEscape(strings.Join(s.config.GithubAuth.Scopes, " "))},
		"redirect_uri": {
			fmt.Sprintf(
				"%s%s?back=%s",
				s.config.Server.PublicURL,
				s.config.GithubAuth.RedirectCompletePath,
				url.QueryEscape(back),
			),
		},
	}

	return fmt.Sprintf("https://github.com/login/oauth/authorize?%s", query.Encode())
}
