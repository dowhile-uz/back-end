package libfx

import (
	configlibfx "dowhile.uz/back-end/lib/config"
	githubclientlibfx "dowhile.uz/back-end/lib/github-client"
	routerlibfx "dowhile.uz/back-end/lib/router"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"lib",
	configlibfx.Module,
	routerlibfx.Module,
	githubclientlibfx.Module,
)
