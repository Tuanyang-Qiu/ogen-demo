// Package product
package product

import (
	"cmp"
	"context"
	"errors"
	"slices"
	"sync"
)

var ErrNoRows = errors.New("no product found")

type Product struct {
	ID          int64
	Name        string
	Price       float64
	Description *string
}

type Repository interface {
	GetProduct(ctx context.Context, id int64) (Product, error)
	ListProducts(ctx context.Context) ([]Product, error)
	UpdateProduct(ctx context.Context, p Product) (Product, error)
}

var _ Repository = (*MemoryRepository)(nil)

type MemoryRepository struct {
	mu       sync.RWMutex
	products map[int64]Product
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		products: map[int64]Product{
			1: {ID: 1, Name: "Flat white milk 2L", Price: 4.2},
			2: {ID: 2, Name: "Sourdough Loaf", Price: 6.50},
		},
	}
}

func (r *MemoryRepository) GetProduct(_ context.Context, id int64) (Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	p, ok := r.products[id]
	if !ok {
		return Product{}, ErrNoRows
	}
	return p, nil
}

func (r *MemoryRepository) ListProducts(_ context.Context) ([]Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var allProducts []Product
	for _, e := range r.products {
		allProducts = append(allProducts, e)
	}

	slices.SortFunc(allProducts, func(a, b Product) int {
		return cmp.Compare(a.ID, b.ID)
	})

	return allProducts, nil
}
func (r *MemoryRepository) UpdateProduct(ctx context.Context, p Product) (Product, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.products[p.ID]; !ok {
		return Product{}, ErrNoRows
	}

	r.products[p.ID] = p
	return r.products[p.ID], nil
}
