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
		switch {
		case errors.Is(err, productsvc.ErrProductNotFound):
			return &ogen.PatchProductNotFound{Message: err.Error()}, nil
		case errors.Is(err, productsvc.ErrProductNameTaken):
			return &ogen.PatchProductConflict{Message: err.Error()}, nil
		}

		return nil, err
	}

	return pack.Product(p), nil
}
