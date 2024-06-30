package api

import (
	"github.com/havus/nekasa-quote-server/internal/app"
	"github.com/havus/nekasa-quote-server/internal/domain/user"
	v1 "github.com/havus/nekasa-quote-server/internal/interfaces/api/v1"
	"go.uber.org/fx"
)

type Params struct {
	fx.In

	Server      *app.Server
	UserService user.UserService
}

func RegisterRoutes(params Params) {
	RegisterHealthRoutes(params.Server.Engine)
	v1.RegisterRoutes(params.Server.Engine, params.UserService)
}
