package string_calculator

import (
	"fmt"
	"strconv"
)

type NegativeValueError struct {
	negValue float64
}

func (n NegativeValueError) Error() string {
	err := "Negative not allowed : "
	return err
}

func (n NegativeValueError) ValueStr() string {
	return strconv.FormatFloat(n.negValue, 'f', -1, 64)
}

type MissingValueError struct {
	index int
	char  byte
}

func (m MissingValueError) Error() string {
	charStr := ""
	if m.char == '\n' {
		charStr += "\\n"
	} else {
		charStr += string(m.char)
	}
	return fmt.Sprintf("Number expected but '%s' found at position %d.", charStr, m.index)
}
