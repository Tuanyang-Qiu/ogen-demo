// Package product
package product

import (
	"context"
	"errors"

	"github.com/tuanyang-qiu/ogen-demo/internal/handler/ogen"
	"github.com/tuanyang-qiu/ogen-demo/internal/handler/product/v1/pack"
	"github.com/tuanyang-qiu/ogen-demo/internal/handler/product/v1/unpack"
	productsvc "github.com/tuanyang-qiu/ogen-demo/internal/service/product"
)

func (h *Handler) PatchProduct(ctx context.Context, req *ogen.ProductPatch, params ogen.PatchProductParams) (ogen.PatchProductRes, error) {
	p, err := h.svc.Update(ctx, unpack.ProductPatch(params.ID, req))
	if err != nil {
		if errors.Is(err, productsvc.ErrProductNotFound) {
			return &ogen.Error{Message: err.Error()}, nil
		}

		return nil, err
	}

	return pack.Product(p), nil
}
