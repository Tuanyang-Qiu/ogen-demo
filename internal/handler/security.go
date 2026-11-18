// Package handler
package handler

import (
	"context"
	"log/slog"

	"github.com/tuanyang-qiu/ogen-demo/internal/handler/ogen"
)

type SecurityHandler struct {
	logger *slog.Logger
}

func NewSecurityHandler(logger *slog.Logger) *SecurityHandler {
	return &SecurityHandler{logger: logger}
}

func (h *SecurityHandler) HandleJwtAuth(ctx context.Context,
	operationName ogen.OperationName,
	t ogen.JwtAuth,
) (context.Context, error) {
	h.logger.InfoContext(ctx,
		"HandleJwtAuth called",
		slog.String("operation", operationName),
		slog.Any("roles", t.Roles))

	return ctx, nil
}

var _ ogen.SecurityHandler = (*SecurityHandler)(nil)
