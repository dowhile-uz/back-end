package githubAuthServiceFx

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	usermodelfx "dowhile.uz/back-end/models/user"
)

type (
	RefreshAccessTokenRequest struct {
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
		RefreshToken string `json:"refresh_token"`
		GrantType    string `json:"grant_type"`
	}
)

func (s *Service) RefreshTokens(ctx context.Context, refreshToken string) (*usermodelfx.GithubAuthTokensResponse, error) {
	body := &RefreshAccessTokenRequest{
		ClientID:     s.config.GithubAuth.ClientID,
		ClientSecret: s.config.GithubAuth.ClientSecret,
		RefreshToken: refreshToken,
		GrantType:    "refresh_token",
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

	if responseBody.Error != "" {
		return nil, fmt.Errorf("github auth error: %s", responseBody.Error)
	}

	return &responseBody, nil
}
