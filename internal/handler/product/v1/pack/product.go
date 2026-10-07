// Package pack
package pack

import (
	"github.com/tuanyang-qiu/ogen-demo/internal/handler/ogen"
	productRepo "github.com/tuanyang-qiu/ogen-demo/internal/repo/product"
)

func Product(product productRepo.Product) *ogen.Product {
	return &ogen.Product{
		ID:          product.ID,
		Name:        product.Name,
		Description: optString(product.Description),
		Price:       product.Price,
	}
}
func Products(products []productRepo.Product) []ogen.Product {
	packed := make([]ogen.Product, 0, len(products))
	for _, p := range products {

		r := ogen.Product{
			ID:          p.ID,
			Name:        p.Name,
			Price:       p.Price,
			Description: optString(p.Description),
		}

		packed = append(packed, r)
	}

	return packed
}

func optString(s *string) ogen.OptString {
	if s == nil {
		return ogen.OptString{}
	}

	return ogen.NewOptString(*s)
}
