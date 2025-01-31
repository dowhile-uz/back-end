package githubauthservicefx

import (
	"fmt"
	"net/url"
	"strings"

	"dowhile.uz/back-end/utils"
)

func (s *Service) GetRedirectURL(back string) string {
	query := url.Values{
		"client_id": {s.Config.GithubAuth.ClientID},
		"state":     {utils.RandomBase16String(6)},
		"scope":     {url.QueryEscape(strings.Join(s.Config.GithubAuth.Scopes, " "))},
		"redirect_uri": {
			fmt.Sprintf(
				"%s%s?back=%s",
				s.Config.Server.PublicURL,
				s.Config.GithubAuth.RedirectCompletePath,
				url.QueryEscape(back),
			),
		},
	}

	return fmt.Sprintf("https://github.com/login/oauth/authorize?%s", query.Encode())
}
