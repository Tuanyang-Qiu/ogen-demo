package product

import "errors"

var (
	ErrProductNotFound  = errors.New("product not found")
	ErrProductNameTaken = errors.New("another product already use this name")
)
