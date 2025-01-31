package githubauthservicefx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type (
	GetAccessTokenRequest struct {
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
		Code         string `json:"code"`
		RedirectURL  string `json:"redirect_url"`
	}
	GetAccessTokenResponse struct {
		Error            string `json:"error"`
		ErrorDescription string `json:"error_description"`
		ErrorURI         string `json:"error_uri"`

		AccessToken string `json:"access_token"`
		Scope       string `json:"scope"`
		TokenType   string `json:"token"`
	}
)

func (s *Service) FetchUserCredentials(code, back string) (any, error) {
	body := &GetAccessTokenRequest{
		ClientID:     s.Config.GithubAuth.ClientID,
		ClientSecret: s.Config.GithubAuth.ClientSecret,
		Code:         code,
		RedirectURL: fmt.Sprintf(
			"%s%s?back=%s",
			s.Config.Server.PublicURL,
			s.Config.GithubAuth.RedirectCompletePath,
			url.QueryEscape(back),
		),
	}

	bodyRaw, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", "https://github.com/login/oauth/access_token", bytes.NewBuffer(bodyRaw))
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Accept", "application/json")

	client := http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	var responseBody GetAccessTokenResponse

	if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
		return nil, err
	}

	return responseBody, nil
}
