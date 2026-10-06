// Package handler
package handler

import (
	"log/slog"

	"github.com/tuanyang-qiu/ogen-demo/internal/handler/health"
	"github.com/tuanyang-qiu/ogen-demo/internal/handler/ogen"
	productv1 "github.com/tuanyang-qiu/ogen-demo/internal/handler/product/v1"
)

type (
	productHandler = productv1.Handler
)

type CombinedHandler struct {
	*health.Handler
	*productHandler
}

func NewCombinedHandler(logger *slog.Logger) *CombinedHandler {
	return &CombinedHandler{
		health.NewHealthHandler(logger),
		productv1.NewHandler(logger),
	}
}

var _ ogen.Handler = (*CombinedHandler)(nil)
