package middleware

import (
	"time"

	"github.com/havus/nekasa-quote-server/internal/infrastructure/config"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func LoggingMiddleware(logger *zap.Logger, cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		logger.Info("incoming request",
			zap.String("level", "INFO"),
			zap.String("msg", "incoming request"),
			zap.String("timestamp", time.Now().Format(time.RFC3339Nano)),
			zap.String("caller", "middleware/logging_middleware.go:20"),
			zap.Object("resource", zapcore.ObjectMarshalerFunc(func(enc zapcore.ObjectEncoder) error {
				enc.AddString("hostname", c.Request.Host)
				enc.AddString("version", cfg.Version)
				return nil
			})),
			zap.Object("request", zapcore.ObjectMarshalerFunc(func(enc zapcore.ObjectEncoder) error {
				enc.AddString("id", c.GetString("request_id"))
				enc.AddString("user_agent", c.Request.UserAgent())
				enc.AddString("remote_ip", c.ClientIP())
				enc.AddString("method", c.Request.Method)
				enc.AddString("full_path", c.Request.URL.Path)
				return nil
			})),
		)

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()

		logger.Info("completed request",
			zap.String("level", "INFO"),
			zap.String("msg", "completed request"),
			zap.String("timestamp", time.Now().Format(time.RFC3339Nano)),
			zap.String("caller", "middleware/logging_middleware.go:39"),
			zap.Object("resource", zapcore.ObjectMarshalerFunc(func(enc zapcore.ObjectEncoder) error {
				enc.AddString("hostname", c.Request.Host)
				enc.AddString("version", cfg.Version)
				return nil
			})),
			zap.Object("request", zapcore.ObjectMarshalerFunc(func(enc zapcore.ObjectEncoder) error {
				enc.AddString("id", c.GetString("request_id"))
				enc.AddString("user_agent", c.Request.UserAgent())
				enc.AddString("remote_ip", c.ClientIP())
				enc.AddString("method", c.Request.Method)
				enc.AddString("full_path", c.Request.URL.Path)
				return nil
			})),
			zap.Object("response", zapcore.ObjectMarshalerFunc(func(enc zapcore.ObjectEncoder) error {
				enc.AddInt("status", status)
				enc.AddDuration("latency", latency)
				return nil
			})),
		)
	}
}
