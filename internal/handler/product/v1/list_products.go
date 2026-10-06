// Package product defines list products api
package product

import (
	"cmp"
	"context"
	"slices"

	"github.com/tuanyang-qiu/ogen-demo/internal/handler/ogen"
)

func (h *Handler) ListProducts(ctx context.Context) ([]ogen.Product, error) {
	allProducts := make([]ogen.Product, 0, len(h.products))

	for _, p := range h.products {
		allProducts = append(allProducts, p)
	}

	slices.SortFunc(allProducts, func(a, b ogen.Product) int {
		return cmp.Compare(a.ID, b.ID)
	})

	return allProducts, nil
}
