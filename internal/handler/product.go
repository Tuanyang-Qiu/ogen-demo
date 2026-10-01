// implement
package handler

import (
	"context"

	"github.com/tuanyang-qiu/ogen-demo/internal/oas"
)

type ProductHandler struct {
	products map[int64]oas.Product
}

var _ oas.Handler = (*ProductHandler)(nil)

func NewProductHandler() *ProductHandler {
	return &ProductHandler{
		products: map[int64]oas.Product{
			1: {ID: 1, Name: "Flat white milk 2L", Price: 4.2},
			2: {ID: 2, Name: "Sourdough Loaf", Price: 6.50},
		},
	}
}

func (h *ProductHandler) GetProduct(ctx context.Context, params oas.GetProductParams) (oas.GetProductRes, error) {
	p, ok := h.products[params.ID]
	if !ok {
		return &oas.Error{Message: "product not found"}, nil
	}

	return &p, nil
}
