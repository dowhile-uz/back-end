package controllersfx

import (
	githubauthcontrollerfx "dowhile.uz/back-end/controllers/github-auth"
	profilecontrollerfx "dowhile.uz/back-end/controllers/profile"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"controllers",
	githubauthcontrollerfx.Module,
	profilecontrollerfx.Module,
)
