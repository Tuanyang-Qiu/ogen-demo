// Package product
package product

import (
	"context"
	"errors"
	"fmt"

	productRepo "github.com/tuanyang-qiu/ogen-demo/internal/repo/product"
)

type Service interface {
	Get(ctx context.Context, id int64) (productRepo.Product, error)
	List(ctx context.Context) ([]productRepo.Product, error)
}

var _ Service = (*service)(nil)

type service struct {
	repo productRepo.Repository
}

func NewService(repo productRepo.Repository) Service {
	if repo == nil {
		panic("missing repo in product.NewService. Aborting!")
	}
	return &service{repo: repo}
}

func (s *service) Get(ctx context.Context, id int64) (productRepo.Product, error) {
	p, err := s.repo.GetProduct(ctx, id)
	if err != nil {
		if errors.Is(err, productRepo.ErrNoRows) {
			return productRepo.Product{}, ErrProductNotFound
		}

		return productRepo.Product{}, fmt.Errorf("get product %d: %w", id, err)
	}

	return p, nil
}

func (s *service) List(ctx context.Context) ([]productRepo.Product, error) {
	ps, err := s.repo.ListProducts(ctx)
	if err != nil {
		return []productRepo.Product{}, fmt.Errorf("list products: %w", err)
	}

	return ps, nil
}
