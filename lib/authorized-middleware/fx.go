package authorizedMiddlewareLibFx

import (
	"context"
	"net/http"
	"strings"

	configLibFx "dowhile.uz/back-end/lib/config"
	userModelFx "dowhile.uz/back-end/models/user"
	"github.com/danielgtaylor/huma/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/fx"
)

var Module = fx.Module("lib.authorized-middleware", fx.Provide(New))

type (
	Params struct {
		fx.In
		Config    *configLibFx.Config
		UserModel *userModelFx.Model
	}
	Middleware struct {
		config    *configLibFx.Config
		userModel *userModelFx.Model
	}
)

func (m *Middleware) GetMiddleware(api huma.API) func(huma.Context, func(huma.Context)) {
	return func(ctx huma.Context, next func(huma.Context)) {
		authorization := ctx.Header("Authorization")
		authorization = strings.TrimPrefix(authorization, "Bearer ")

		token, err := jwt.Parse(authorization, func(_ *jwt.Token) (any, error) {
			return []byte(m.config.Server.JWTSecret), nil
		})
		if err != nil {
			huma.WriteErr(api, ctx, http.StatusUnauthorized, "Token parse error", err)
			return
		}

		if !token.Valid {
			huma.WriteErr(api, ctx, http.StatusUnauthorized, "Token invalid")
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			huma.WriteErr(api, ctx, http.StatusInternalServerError, "Token claims is not jwt.MapClaims")
			return
		}

		userID, ok := claims["user_id"].(float64)

		if !ok {
			huma.WriteErr(api, ctx, http.StatusUnauthorized, "Token claims doesn't contain user_id")
			return
		}

		user, err := m.userModel.GetUser(context.Background(), int64(userID))
		if err != nil {
			huma.WriteErr(api, ctx, http.StatusUnauthorized, "User not found", err)
			return
		}

		ctx = huma.WithValue(ctx, "user", user)
		next(ctx)
	}
}

func New(p Params) *Middleware {
	return &Middleware{
		config:    p.Config,
		userModel: p.UserModel,
	}
}
