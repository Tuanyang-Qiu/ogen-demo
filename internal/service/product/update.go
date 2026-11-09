// Package product
package product

import (
	"context"
	"errors"
	"fmt"

	productRepo "github.com/tuanyang-qiu/ogen-demo/internal/repo/product"
	"github.com/tuanyang-qiu/ogen-demo/pkg/optional"
)

type UpdateInput struct {
	ID          int64
	Name        optional.Value[string]
	Price       optional.Value[float64]
	Description optional.Value[*string]
}

func (s *service) Update(ctx context.Context, in UpdateInput) (productRepo.Product, error) {
	p, err := s.Get(ctx, in.ID)
	if err != nil {
		return productRepo.Product{}, err
	}

	if in.Name.Set {
		p.Name = in.Name.Value
	}

	if in.Price.Set {
		p.Price = in.Price.Value
	}

	if in.Description.Set {
		p.Description = in.Description.Value
	}

	updated, err := s.repo.UpdateProduct(ctx, p)
	if err != nil {
		if errors.Is(err, productRepo.ErrNoRows) {
			return productRepo.Product{}, ErrProductNotFound
		}

		return productRepo.Product{}, fmt.Errorf("update product %d: %w", in.ID, err)
	}

	return updated, nil

}
