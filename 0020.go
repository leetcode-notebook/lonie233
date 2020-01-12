package main

func isValid(s string) bool {

	if len(s)%2 != 0 || len(s) == 0 {
		return false
	}

	var leftStack = &Stack{}

	for _, c := range s {
		if c == '{' || c == '[' || c == '(' {
			push(c, leftStack)
		} else if c == '}' {
			if '{' != pop(leftStack) {
				return false
			}
		} else if c == ']' {
			if '[' != pop(leftStack) {
				return false
			}
		} else if c == ')' {
			if '(' != pop(leftStack) {
				return false
			}
		} else {
			return false
		}
	}
	return len(leftStack.ele) == 0
}

type Stack struct {
	ele []int32
}

func pop(stack *Stack) int32 {
	if len(stack.ele) == 0 {
		return 654321
	}
	s := stack.ele[len(stack.ele)-1]
	stack.ele = stack.ele[:len(stack.ele)-1]
	return s
}

func push(c int32, stack *Stack) {
	stack.ele = append(stack.ele, c)
}
