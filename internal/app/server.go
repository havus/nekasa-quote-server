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
	engine := gin.New()

	// Set trusted proxies
	if len(config.TrustedProxies) > 0 {
		engine.SetTrustedProxies(strings.Split(config.TrustedProxies[0], ","))
	}

	engine.Use(logger.GinMiddleware())      // Menggunakan middleware logger
	engine.Use(middleware.AuthMiddleware()) // Menggunakan middleware auth
	api.RegisterHealthRoutes(engine)        // Register endpoint /health

	return &Server{
		Engine: engine,
		Config: config,
		Logger: logger,
	}
}

func StartServer(lc fx.Lifecycle, server *Server) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go server.Engine.Run(":" + server.Config.Port)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}

func NewLogger() (*zap.Logger, error) {
	return zap.NewProduction()
}
