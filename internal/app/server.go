package app

import (
	"context"
	"strings"

	"github.com/havus/nekasa-quote-server/internal/infrastructure/config"
	"github.com/havus/nekasa-quote-server/internal/interfaces/api"
	"github.com/havus/nekasa-quote-server/internal/interfaces/middleware"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type Server struct {
	Engine *gin.Engine
	Config *config.Config
}

func NewServer(config *config.Config) *Server {
	engine := gin.New()

	// Set trusted proxies
	if len(config.TrustedProxies) > 0 {
		engine.SetTrustedProxies(strings.Split(config.TrustedProxies[0], ","))
	}

	engine.Use(middleware.AuthMiddleware())
	api.RegisterHealthRoutes(engine)

	return &Server{
		Engine: engine,
		Config: config,
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
