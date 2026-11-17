// Package ogen defines files for code generation
package ogen

//go:generate go tool ogen --config config.yaml --package ogen --target ../../internal/handler/ogen --clean ../../api/openapi.yaml
