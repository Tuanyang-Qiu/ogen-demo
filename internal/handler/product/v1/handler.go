// Package product defines products v1
package product

import (
	"log/slog"

	"github.com/tuanyang-qiu/ogen-demo/internal/handler/ogen"
)

type Handler struct {
	logger   *slog.Logger
	products map[int64]ogen.Product
}

func NewHandler(logger *slog.Logger) *Handler {
	if logger == nil {
		logger = slog.New(slog.DiscardHandler)

	}

	return &Handler{
		logger: logger,
		products: map[int64]ogen.Product{
			1: {ID: 1, Name: "Flat white milk 2L", Price: 4.2},
			2: {ID: 2, Name: "Sourdough Loaf", Price: 6.50}},
	}
}

var _ ogen.ProductHandler = (*Handler)(nil)
