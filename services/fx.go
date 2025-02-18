package servicesFx

import (
	githubServiceFx "dowhile.uz/back-end/services/github"
	githubAuthServiceFx "dowhile.uz/back-end/services/github-auth"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"services",
	githubAuthServiceFx.Module,
	githubServiceFx.Module,
)
