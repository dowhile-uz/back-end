package githubAuthServiceFx

import "github.com/golang-jwt/jwt/v5"

type AccessTokenClaims struct {
	jwt.RegisteredClaims
	UserID int64 `json:"user_id"`
}
