package monad

type ObjectOrError[T any] struct {
	Value  T
	Errors []error
}
type V any

//var globalError error
//
//func handle_error() error {
//	if globalError.Error() != "" {
//		return globalError
//	}
//	return nil
//}
//func square[W int|float32|float64](num W) ObjectOrError[V] {
//	return ObjectOrError[V]{ value: num * num, errors: []error{}}
//}

func (o ObjectOrError[T]) Bind(f func(num T) ObjectOrError[V]) ObjectOrError[V] {
	return f(o.value)
}

func Unit[T any](num T) ObjectOrError[T] {
	return ObjectOrError[T]{value: num, errors: []error{}}
}
