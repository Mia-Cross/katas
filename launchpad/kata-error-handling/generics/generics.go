package generics

import (
	"reflect"
)

// type ValueType struct {
// 	value string
// }
type ValueType[T float32|float64|string] struct {
	value T
}

func (v ValueType[T])add[[T float32|float64|string](total T, value T) T {
	return total + value
}

func Sum[T float32|float64|string|ValueType](arr []T) T {
	var result T
	reflect.ArrayOf(len(arr), )

	for _, elem := range arr {
		result = add(result, elem)
	}
	return result
}