package controllersFx

import (
	editorControllerFx "dowhile.uz/back-end/controllers/editor"
	githubControllerFx "dowhile.uz/back-end/controllers/github"
	githubAuthControllerFx "dowhile.uz/back-end/controllers/github-auth"
	profileControllerFx "dowhile.uz/back-end/controllers/profile"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"controllers",
	githubAuthControllerFx.Module,
	profileControllerFx.Module,
	editorControllerFx.Module,
	githubControllerFx.Module,
)
