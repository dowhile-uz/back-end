package githubauthservicefx

import (
	"fmt"
	"net/url"

	"dowhile.uz/back-end/utils"
)

func (s *Service) GetRedirectURL(back string) string {
	return fmt.Sprintf(
		"https://github.com/login/oauth/authorize?client_id=%s&state=%s&redirect_uri=%s",
		s.Config.Githubauth.Clientid,
		utils.RandomBase16String(6),
		url.QueryEscape(
			fmt.Sprintf(
				"%s%s?back=%s",
				s.Config.Server.Publicurl,
				s.Config.Githubauth.Redirectpath,
				url.QueryEscape(back),
			),
		),
	)
}
