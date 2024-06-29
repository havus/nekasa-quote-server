package di

import (
	"github.com/havus/nekasa-quote-server/internal/app"
	"github.com/havus/nekasa-quote-server/internal/infrastructure/config"

	"go.uber.org/fx"
)

func New() *fx.App {
	return fx.New(
		fx.Provide(
			config.LoadConfig,
			app.NewServer,
		),
		fx.Invoke(app.StartServer),
	)
}
