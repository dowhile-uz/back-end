package githubAuthServiceFx

import (
	"fmt"
	"net/url"
	"strings"

	"dowhile.uz/back-end/utils"
)

func (s *Service) GetGithubRedirectURL(back string) string {
	query := url.Values{
		"client_id":    {s.config.GithubAuth.ClientID},
		"state":        {utils.RandomBase16String(6)},
		"scope":        {url.QueryEscape(strings.Join(s.config.GithubAuth.Scopes, ","))},
		"allow_signup": {"true"},
		"redirect_uri": {
			fmt.Sprintf(
				"%s%s?back=%s",
				s.config.Server.PublicURL,
				s.config.GithubAuth.RedirectCompletePath,
				url.QueryEscape(back),
			),
		},
	}

	// https://github.com/apps/vercel/installations/new

	// return fmt.Sprintf("https://github.com/apps/%s/installations/new?%s", s.config.GithubAuth.AppName, query.Encode())
	return fmt.Sprintf("https://github.com/login/oauth/authorize?%s", query.Encode())
}
