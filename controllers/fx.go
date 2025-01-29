package controllersfx

import (
	githubauthcontrollerfx "dowhile.uz/back-end/controllers/github-auth"
	helloworldcontrollerfx "dowhile.uz/back-end/controllers/hello-world"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"controllers",
	helloworldcontrollerfx.Module,
	githubauthcontrollerfx.Module,
)
