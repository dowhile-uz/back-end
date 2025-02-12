package servicesFx

import (
	editorServiceFx "dowhile.uz/back-end/services/editor"
	githubAuthServiceFx "dowhile.uz/back-end/services/github-auth"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"services",
	githubAuthServiceFx.Module,
	editorServiceFx.Module,
)
