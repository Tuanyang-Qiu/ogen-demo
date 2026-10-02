// implement
package handler

import (
	"cmp"
	"context"
	"slices"

	"github.com/tuanyang-qiu/ogen-demo/internal/oas"
)

type ProductHandler struct {
	products map[int64]oas.Product
}

// Check if all interfaces are implemented or not
// if not, 'go build' will throw errors at compile time rathan than run time
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

func (h *ProductHandler) ListProducts(ctx context.Context) ([]oas.Product, error) {
	allProducts := make([]oas.Product, 0, len(h.products))

	for _, e := range h.products {
		allProducts = append(allProducts, e)
	}

	//sort.Slice(p, func(i int, j int) bool { return p[i].ID < p[j].ID }) //old GO syntax
	slices.SortFunc(allProducts, func(a, b oas.Product) int { //New GO syntax
		return cmp.Compare(a.ID, b.ID)
	})

	return allProducts, nil
}
