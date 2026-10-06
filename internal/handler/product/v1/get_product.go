package product

import (
	"context"

	"github.com/tuanyang-qiu/ogen-demo/internal/handler/ogen"
)

func (h *Handler) GetProduct(_ context.Context, params ogen.GetProductParams) (ogen.GetProductRes, error) {
	p, ok := h.products[params.ID]
	if !ok {
		return &ogen.Error{Message: "product not found"}, nil
	}
	return &p, nil
}
