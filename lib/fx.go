package libfx

import (
	authorizedmiddlewarelibfx "dowhile.uz/back-end/lib/authorized-middleware"
	configlibfx "dowhile.uz/back-end/lib/config"
	githubclientlibfx "dowhile.uz/back-end/lib/github-client"
	postgreslibfx "dowhile.uz/back-end/lib/postgres"
	redislibfx "dowhile.uz/back-end/lib/redis"
	routerlibfx "dowhile.uz/back-end/lib/router"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"lib",
	configlibfx.Module,
	routerlibfx.Module,
	githubclientlibfx.Module,
	postgreslibfx.Module,
	redislibfx.Module,
	authorizedmiddlewarelibfx.Module,
)
