// Package product defines list products api
package product

import (
	"context"

	"github.com/tuanyang-qiu/ogen-demo/internal/handler/ogen"
	"github.com/tuanyang-qiu/ogen-demo/internal/handler/product/v1/pack"
)

func (h *Handler) ListProducts(ctx context.Context) ([]ogen.Product, error) {

	prodAll, err := h.svc.List(ctx)
	if err != nil {
		return nil, err
	}

	return pack.Products(prodAll), nil
}
