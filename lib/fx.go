package libFx

import (
	authorizedMiddlewareLibFx "dowhile.uz/back-end/lib/authorized-middleware"
	configLibFx "dowhile.uz/back-end/lib/config"
	errorCodesLibFx "dowhile.uz/back-end/lib/error-codes"
	githubClientLibFx "dowhile.uz/back-end/lib/github-client"
	postgresLibFx "dowhile.uz/back-end/lib/postgres"
	redisLibFx "dowhile.uz/back-end/lib/redis"
	routerLibFx "dowhile.uz/back-end/lib/router"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"lib",
	configLibFx.Module,
	routerLibFx.Module,
	githubClientLibFx.Module,
	postgresLibFx.Module,
	redisLibFx.Module,
	authorizedMiddlewareLibFx.Module,
	errorCodesLibFx.Module,
)
