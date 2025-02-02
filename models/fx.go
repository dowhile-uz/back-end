package modelsfx

import (
	usermodelfx "dowhile.uz/back-end/models/user"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"models",
	usermodelfx.Module,
)
