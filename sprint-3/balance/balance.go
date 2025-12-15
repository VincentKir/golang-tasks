package balance

import (
	"slices"
)

type Stack []string

// Добавление элемента
func (s Stack) push(str string) Stack {
	return append(s, str)
}

// Удаление элемента с вершины стека
func (s Stack) pop() Stack {
	return s[:len(s)-1]
}

// Просмотр последнего элемента (без удаления)
func (s Stack) top() string {
	if len(s) > 0 {
		return s[len(s)-1]
	}
	return ""
}

var containsSymbols = map[string]string{
	")": "(",
	"}": "{",
	"]": "[",
}

var containsSymbolsOpen = []string{
	"(",
	"{",
	"[",
}

var containsSymbolsClose = []string{
	")",
	"}",
	"]",
}

func Balance(s string) bool {
	stack := make(Stack, 0)
	for _, str := range s {
		str := string(str)
		if slices.Contains(containsSymbolsOpen, str) {
			stack = stack.push(str)
			continue
		}

		if slices.Contains(containsSymbolsClose, str) {
			symbolOpen := containsSymbols[str]
			top := stack.top()

			if top != symbolOpen {
				return false
			}

			stack = stack.pop()
		}
	}

	if len(stack) == 0 {
		return true
	}

	return false

}
