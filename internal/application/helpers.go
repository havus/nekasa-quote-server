package application

import (
	"context"

	"github.com/google/uuid"
)

func generateUID() string {
	return uuid.New().String()
}

func getCurrentIP(ctx context.Context) string {
	clientIP, ok := ctx.Value("clientIP").(string)
	if ok {
		return clientIP
	}
	return "unknown"
}
