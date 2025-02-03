package githubauthservicefx

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	usermodelfx "dowhile.uz/back-end/models/user"
)

type (
	GetAccessTokenRequest struct {
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
		Code         string `json:"code"`
		RedirectURL  string `json:"redirect_url"`
	}
)

func (s *Service) FetchAuthorizationTokens(ctx context.Context, code, back string) (*usermodelfx.GithubAuthTokensResponse, error) {
	body := &GetAccessTokenRequest{
		ClientID:     s.config.GithubAuth.ClientID,
		ClientSecret: s.config.GithubAuth.ClientSecret,
		Code:         code,
		RedirectURL: fmt.Sprintf(
			"%s%s?back=%s",
			s.config.Server.PublicURL,
			s.config.GithubAuth.RedirectCompletePath,
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

	var responseBody usermodelfx.GithubAuthTokensResponse

	if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
		return nil, err
	}

	return &responseBody, nil
}
