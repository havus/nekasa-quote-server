package app

import (
	"context"
	"strings"

	"github.com/havus/nekasa-quote-server/internal/infrastructure/config"
	"github.com/havus/nekasa-quote-server/internal/infrastructure/logger"
	"github.com/havus/nekasa-quote-server/internal/interfaces/middleware"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Server struct {
	Engine *gin.Engine
	Config *config.Config
	Logger *logger.Logger
	DB     *gorm.DB
}

func NewServer(config *config.Config, logger *logger.Logger, db *gorm.DB) *Server {
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

	return &Server{
		Engine: engine,
		Config: config,
		Logger: logger,
		DB:     db,
	}
}

func StartServer(lc fx.Lifecycle, server *Server) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			server.Logger.GeneralLog(
				"INFO",
				"internal/app/server.StartServer",
				"Server is starting on port "+server.Config.Port,
				nil,
			)
			go func() {
				if err := server.Engine.Run(":" + server.Config.Port); err != nil {
					server.Logger.GeneralLog("ERROR", "StartServer", "Failed to start server", map[string]interface{}{"error": err})
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			server.Logger.GeneralLog("INFO", "StartServer", "Server is stopping", nil)
			return nil
		},
	})
}

func NewLogger() (*zap.Logger, error) {
	return zap.NewProduction()
}
