// Package unpack
package unpack

import (
	"github.com/tuanyang-qiu/ogen-demo/internal/handler/ogen"
	productsvc "github.com/tuanyang-qiu/ogen-demo/internal/service/product"
	"github.com/tuanyang-qiu/ogen-demo/pkg/optional"
)

func ProductPatch(id int64, req *ogen.ProductPatch) productsvc.UpdateInput {
	return productsvc.UpdateInput{
		ID:          id,
		Name:        optString(req.Name),
		Price:       optFloat64(req.Price),
		Description: optNilString(req.Description),
	}
}

func optString(o ogen.OptString) optional.Value[string] {
	if !o.Set {
		return optional.Value[string]{}
	}
	return optional.From(o.Value)
}

func optFloat64(o ogen.OptFloat64) optional.Value[float64] {
	if !o.Set {
		return optional.Value[float64]{}
	}

	return optional.From(o.Value)
}

func optNilString(o ogen.OptNilString) optional.Value[*string] {

	if !o.Set {
		return optional.Value[*string]{}
	}
	if o.Null {
		return optional.Null[string]()
	}

	value := o.Value
	return optional.From(&value)
}
