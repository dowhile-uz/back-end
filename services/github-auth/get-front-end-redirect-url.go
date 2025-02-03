package githubauthservicefx

import (
	"fmt"
	"net/url"
)

func (s *Service) GetFrontEndRedirectURL(accessToken, back string) string {
	query := url.Values{
		"access_token": {accessToken},
		"back":         {back},
	}

	return fmt.Sprintf(
		"%s%s?%s",
		s.config.Server.FrontEndURL,
		s.config.GithubAuth.RedirectFrontEndPath,
		query.Encode(),
	)
}
