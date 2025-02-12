package modelsFx

import (
	userModelFx "dowhile.uz/back-end/models/user"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"models",
	userModelFx.Module,
)
