// Package optional
package optional

type Value[T any] struct {
	Value T
	Set   bool
}

func From[T any](value T) Value[T] {
	return Value[T]{Value: value, Set: true}
}

func Null[T any]() Value[*T] {
	return Value[*T]{Set: true}
}
