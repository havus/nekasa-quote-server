package app

import (
	"context"
	"strings"

	"github.com/havus/nekasa-quote-server/internal/infrastructure/config"
	"github.com/havus/nekasa-quote-server/internal/infrastructure/logger"
	"github.com/havus/nekasa-quote-server/internal/interfaces/api"
	"github.com/havus/nekasa-quote-server/internal/interfaces/middleware"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Server struct {
	Engine *gin.Engine
	Config *config.Config
	Logger *logger.Logger
}

func NewServer(config *config.Config, logger *logger.Logger) *Server {
	ginMode := gin.DebugMode
	if config.EnvMode == "PRODUCTION" {
		ginMode = gin.ReleaseMode
	}
	gin.SetMode(ginMode)

	engine := gin.New()

	// Set trusted proxies
	if len(config.TrustedProxies) > 0 {
		engine.SetTrustedProxies(strings.Split(config.TrustedProxies[0], ","))
	}

	engine.Use(logger.GinMiddleware())
	engine.Use(middleware.AuthMiddleware())
	api.RegisterHealthRoutes(engine)

	return &Server{
		Engine: engine,
		Config: config,
		Logger: logger,
	}
}

func StartServer(lc fx.Lifecycle, server *Server) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			server.Logger.GeneralLog(
				"info",
				"internal/app/server.StartServer",
				"Server is starting on port "+server.Config.Port,
				nil,
			)
			go server.Engine.Run(":" + server.Config.Port)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			server.Logger.GeneralLog("info", "StartServer", "Server is stopping", nil)
			return nil
		},
	})
}

func NewLogger() (*zap.Logger, error) {
	return zap.NewProduction()
}
