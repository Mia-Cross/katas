package monad

import (
	"testing"
)

func TestSumOfInts(t *testing.T) {
	//o := ObjectOrError[int]{value: 56}
	baseObject := Unit[int](67)
	transformedObject := baseObject.Bind(square())
}
