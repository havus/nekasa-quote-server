package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/havus/nekasa-quote-server/internal/infrastructure/config"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	config    *config.Config
	hostname  string
	version   string
	zapLogger *zap.Logger
}

func NewLogger(cfg *config.Config) *Logger {
	hostname, _ := os.Hostname()

	zapLogger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(os.Stdout),
		zapcore.InfoLevel,
	))

	return &Logger{
		zapLogger: zapLogger,
		hostname:  hostname,
		version:   cfg.Version,
		config:    cfg,
	}
}

func (l *Logger) GeneralLog(level, caller, msg string, tags map[string]interface{}) {
	if level == "debug" && l.config.EnvMode == "PRODUCTION" {
		return
	}

	entry := make(map[string]interface{})
	entry["level"] = level
	entry["message"] = msg
	entry["timestamp"] = time.Now().UTC().Format(time.RFC3339)
	entry["resource"] = map[string]interface{}{
		"hostname": l.hostname,
		"version":  l.version,
	}
	entry["caller"] = caller
	entry["tags"] = tags

	log, _ := json.Marshal(entry)
	fmt.Println(string(log))
}

func (l *Logger) GinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now().UTC()

		entry := make(map[string]interface{})
		entry["level"] = "INFO"
		entry["message"] = "request started"
		entry["timestamp"] = start.Format(time.RFC3339)
		entry["resource"] = map[string]interface{}{
			"hostname": l.hostname,
			"version":  l.version,
		}
		entry["request"] = map[string]interface{}{
			"id":         c.GetString("RequestID"),
			"user_agent": c.Request.UserAgent(),
			"remote_ip":  c.ClientIP(),
			"method":     c.Request.Method,
			"full_path":  c.FullPath(),
		}

		log, _ := json.Marshal(entry)
		fmt.Println(string(log))

		c.Next()

		latency := time.Since(start)
		entry["message"] = "request completed"
		entry["response"] = map[string]interface{}{
			"status":  c.Writer.Status(),
			"latency": latency.String(),
		}

		log, _ = json.Marshal(entry)
		fmt.Println(string(log))
	}
}

func (l *Logger) FxLogger() fxevent.Logger {
	return &fxevent.ZapLogger{Logger: l.zapLogger}
}
