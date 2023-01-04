package validparentheses

func isValid(s string) bool {
	stack := make([]rune, 0)
	brackets := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, symbol := range s {
		switch symbol {
		case '(', '[', '{':
			stack = append([]rune{symbol}, stack...)
		case ')', ']', '}':
			if len(stack) == 0 || stack[0] != brackets[symbol] {
				return false
			}
			stack = stack[1:]
		}
	}

	return len(stack) == 0
}
