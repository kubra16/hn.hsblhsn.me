package backend

import (
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Options(
		fx.Invoke(RegisterRoutes),
	)
}