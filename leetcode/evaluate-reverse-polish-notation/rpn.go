package rpn

import "strconv"

type operator func(int, int) int

var operators = map[string]operator{
	"+": func(a, b int) int { return a + b },
	"-": func(a, b int) int { return a - b },
	"*": func(a, b int) int { return a * b },
	"/": func(a, b int) int { return a / b },
}

func evalRPN(tokens []string) int {
	operands := []int{}
	for _, token := range tokens {
		if op, ok := operators[token]; ok {
			n := len(operands)
			a, b := operands[n-2], operands[n-1]
			result := op(a, b)
			operands = append(operands[:n-2], result)
		} else {
			operand, err := strconv.Atoi(token)
			if err != nil {
				panic(err.Error())
			}
			operands = append(operands, operand)
		}
	}
	return operands[0]
}
