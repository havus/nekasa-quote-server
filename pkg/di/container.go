package di

import (
	"github.com/havus/nekasa-quote-server/internal/app"
	"github.com/havus/nekasa-quote-server/internal/infrastructure/config"
	"github.com/havus/nekasa-quote-server/internal/infrastructure/database"
	"github.com/havus/nekasa-quote-server/internal/infrastructure/logger"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func New() *fx.App {
	return fx.New(
		fx.Provide(
			app.NewServer,
			database.Connect,
			logger.NewLogger,
			config.LoadConfig,
		),
		fx.WithLogger(func(log *logger.Logger) fxevent.Logger {
			return log.FxLogger()
		}),
		fx.Invoke(app.StartServer),
	)
}
