package servicesfx

import (
	githubauthservicefx "dowhile.uz/back-end/services/github-auth"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"services",
	githubauthservicefx.Module,
)
