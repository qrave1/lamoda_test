package http

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/qrave1/lamoda_test/internal/interface/http/gen"
)

type key int

const (
	requestIDKey key = iota
)

func RequestIDMiddleware(next gen.StrictHandlerFunc, operationID string) gen.StrictHandlerFunc {
	return func(c *fiber.Ctx, args interface{}) (interface{}, error) {
		requestID := c.Get("X-Request-ID")
		if requestID == "" {
			requestID = uuid.New().String()
		}

		c.SetUserContext(context.WithValue(c.Context(), "requestID", requestID))

		return next(c, args)
	}
}

func RequestIDFromContext(ctx context.Context) string {
	if requestID, ok := ctx.Value(requestIDKey).(string); ok {
		return requestID
	}
	return ""
}
