package string_calculator

import (
	"kata-error-handling/monad"
	"strconv"
	"strings"
)

func outputErrors(errors []error) string {
	//fmt.Println(errors)
	var negNumbers []string
	var toPrint []string
	for _, err := range errors {
		negValueError, ok := err.(NegativeValueError)
		if ok {
			if len(negNumbers) == 0 {
				toPrint = append(toPrint, negValueError.Error())
			}
			negNumbers = append(negNumbers, negValueError.ValueStr())
		}
		missValueError, ok := err.(MissingValueError)
		if ok {
			toPrint = append(toPrint, missValueError.Error())
		}
	}
	output := ""
	for index, errorStr := range toPrint {
		if strings.HasPrefix(errorStr, "Negative ") {
			for i := 0; i < len(negNumbers)-1; i++ {
				errorStr += negNumbers[i] + ", "
			}
			errorStr += negNumbers[len(negNumbers)-1]
		}
		output += errorStr
		if index < len(toPrint)-1 {
			output += "\n"
		}
	}
	return output
}

func checkEmpty(s string) monad.ObjectOrError {
	if len(s) == 0 {
		return monad.ObjectOrError{
			Errors: []error{MissingValueError{}},
		}
	}
}

func convertValueFromStr(s string) monad.ObjectOrError {
	val, err := strconv.ParseFloat(s, 64)
	o := monad.ObjectOrError{ Value: val, Errors: append(o.Errors, err) }
}

func checkNegative(n any) monad.ObjectOrError {
	if
}

func getNextValue(str string, index int, separators string) (string, int) {
	if index >= len(str) && len(str) > 1 {
		return "", len(str)
	}
	for i := index; i < len(str); i++ {
		for j := range separators {
			if separators[j] == str[i] {
				return string(str[index:i]), i
			}
		}
	}
	return str[index:], len(str)
}

func Add(numbers string) string {
	if len(numbers) == 0 {
		return "0"
	}
	//var allErrors []error
	sum := 0.0
	for i := -1; i < len(numbers); {
		value, index := getNextValue(numbers, i+1, ",\n")
		o := monad.ObjectOrError{value}
		o.Bind(checkEmpty(value)).Bind(checkNegative())

		if len(value) == 0 {

		}

	}
	return
}
