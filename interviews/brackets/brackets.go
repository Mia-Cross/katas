package brackets

import "fmt"

func MultipleBrackets(str string) string {
	nbBrackets := 0
	stack := []rune(nil)

	openingCharacters := map[rune]rune{
		'(': ')',
		'[': ']',
		'!': '$',
		'&': '*',
	}
	closingCharacters := map[rune]rune{}
	for oc, cc := range openingCharacters {
		closingCharacters[cc] = oc
	}

	for _, c := range str {
		if _, ok := openingCharacters[c]; ok {
			stack = append(stack, c)
			continue
		}
		if cc, ok := closingCharacters[c]; ok {
			if len(stack) > 0 && stack[len(stack)-1] == cc {
				stack = stack[:len(stack)-1]
				nbBrackets++
			} else {
				return "0"
			}
		}
	}

	if len(stack) > 0 {
		return "0"
	}
	return fmt.Sprintf("1 %d", nbBrackets)
}

func MultipleBrackets2(str string) string {
	nbBrackets := 0
	stack := []rune(nil)

	for _, c := range str {
		switch c {
		case '(', '[':
			stack = append(stack, c)
		case ')':
			if len(stack) > 0 && stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
				nbBrackets++
			} else {
				return "0"
			}
		case ']':
			if len(stack) > 0 && stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1]
				nbBrackets++
			} else {
				return "0"
			}
		}
	}

	if len(stack) > 0 {
		return "0"
	}
	return fmt.Sprintf("1 %d", nbBrackets)
}
