package valid_parentheses

type Stack []byte

func createStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(char byte) bool {
	// println(string(char))

	if char == '(' || char == '{' || char == '[' {
		*s = append(*s, char)
		return true
	}

	if len(*s) == 0 {
		return false
	}

	last := (*s)[len(*s)-1]

	// fmt.Println(string(last), " ", string(char))

	if char == ')' || char == '}' || char == ']' {
		switch last {
		case '(':
			{
				if char == ')' {
					s.Pop()
					return true
				}
			}
		case '{':
			{
				if char == '}' {
					s.Pop()
					return true
				}
			}
		case '[':
			{
				if char == ']' {
					s.Pop()
					return true
				}
			}
		}
	}

	return false
}

func (s *Stack) Pop() (byte, bool) {
	i := len(*s)
	if i == 0 {
		*s = []byte{}
		return 0, false
	}
	item := (*s)[i-1]
	*s = (*s)[:i-1]
	return item, true
}

func isValid(s string) bool {
	stack := createStack()
	if len(s) < 2 {
		return false
	}
	for i := 0; i < len(s); i++ {
		check := stack.Push(s[i])
		if !check {
			return false
		}
	}
	if len(*stack) > 0 {
		return false
	}
	return true
}
