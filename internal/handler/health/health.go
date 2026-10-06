// Package health
package health

import (
	"context"
	"log/slog"

	"github.com/tuanyang-qiu/ogen-demo/internal/handler/ogen"
)

type Handler struct {
	logger *slog.Logger
}

func NewHealthHandler(logger *slog.Logger) *Handler {
	if logger == nil {
		panic("nil logger")
	}

	ps := &Handler{
		logger: logger,
	}

	return ps

}

var _ ogen.HealthHandler = (*Handler)(nil)

func (h *Handler) CheckHealth(ctx context.Context) (*ogen.CheckHealthOK, error) {
	res := new(ogen.CheckHealthOK)
	res.SetStatus(ogen.NewOptString("healthy"))
	return res, nil
}
