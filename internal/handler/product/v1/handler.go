// Package product defines products v1
package product

import (
	"log/slog"

	"github.com/tuanyang-qiu/ogen-demo/internal/handler/ogen"
	productSvc "github.com/tuanyang-qiu/ogen-demo/internal/service/product"
)

type Handler struct {
	logger *slog.Logger
	svc    productSvc.Service
}

func NewHandler(logger *slog.Logger, svc productSvc.Service) *Handler {
	if logger == nil {
		logger = slog.New(slog.DiscardHandler)

	}

	return &Handler{
		logger: logger,
		svc:    svc,
	}
}

var _ ogen.ProductHandler = (*Handler)(nil)
