package githubAuthControllerFx

import (
	"context"
	"net/http"
	"time"

	githubAuthServiceFx "dowhile.uz/back-end/services/github-auth"
	"github.com/danielgtaylor/huma/v2"
	"github.com/golang-jwt/jwt/v5"
)

type (
	CompleteInput struct {
		Back string `query:"back"`
		Code string `query:"code" required:"true"`
	}
	CompleteOutput struct {
		// SetCookie http.Cookie `header:"Set-Cookie"`
		Status int
		URL    string `header:"Location"`
	}
)

func (c *Controller) CompleteHandler(ctx context.Context, input *CompleteInput) (*CompleteOutput, error) {
	o := &CompleteOutput{}

	tokens, err := c.service.FetchAuthorizationTokens(ctx, input.Code, input.Back)
	if err != nil {
		return nil, err
	}

	if (*tokens).Error != "" {
		return nil, huma.Error400BadRequest((*tokens).Error)
	}

	githubUser, err := c.service.FetchUserData(ctx, (*tokens).AccessToken)
	if err != nil {
		return nil, err
	}

	user, err := c.userModel.CreateOrUpdateWithTokens(ctx, githubUser, tokens)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	expiresAt := now.Add(180 * 24 * time.Hour)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, githubAuthServiceFx.AccessTokenClaims{
		UserID: *user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	})

	accessToken, err := token.SignedString([]byte(c.config.Server.JWTSecret))
	if err != nil {
		return nil, err
	}

	// o.SetCookie = http.Cookie{
	// 	Name:    "access_token",
	// 	Value:   accessToken,
	// 	Expires: expiresAt,
	// }
	o.Status = http.StatusTemporaryRedirect
	o.URL = c.service.GetFrontEndRedirectURL(accessToken, input.Back)

	return o, nil
}
