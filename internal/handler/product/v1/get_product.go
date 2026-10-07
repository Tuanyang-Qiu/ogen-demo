package product

import (
	"context"
	"errors"

	"github.com/tuanyang-qiu/ogen-demo/internal/handler/ogen"
	"github.com/tuanyang-qiu/ogen-demo/internal/handler/product/v1/pack"
	productSvc "github.com/tuanyang-qiu/ogen-demo/internal/service/product"
)

func (h *Handler) GetProduct(ctx context.Context, params ogen.GetProductParams) (ogen.GetProductRes, error) {

	p, err := h.svc.Get(ctx, params.ID)
	if err != nil {
		if errors.Is(err, productSvc.ErrProductNotFound) {
			return &ogen.Error{Message: err.Error()}, nil
		}

		return nil, err

	}
	return pack.Product(p), nil
}
