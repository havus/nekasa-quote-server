package di

import (
	"github.com/havus/nekasa-quote-server/internal/app"
	"github.com/havus/nekasa-quote-server/internal/application"
	"github.com/havus/nekasa-quote-server/internal/infrastructure/config"
	"github.com/havus/nekasa-quote-server/internal/infrastructure/database"
	"github.com/havus/nekasa-quote-server/internal/infrastructure/logger"
	"github.com/havus/nekasa-quote-server/internal/infrastructure/repository"
	"github.com/havus/nekasa-quote-server/internal/interfaces/api"
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

			application.NewUserService,
			repository.NewGormUserRepository,
		),
		fx.WithLogger(func(log *logger.Logger) fxevent.Logger {
			return log.FxLogger()
		}),
		fx.Invoke(app.StartServer),
		fx.Invoke(api.RegisterRoutes),
	)
}
