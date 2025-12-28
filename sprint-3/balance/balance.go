package balance

import (
	"slices"
)

type Stack []rune

// Добавление элемента
func (s *Stack) push(str rune) {
	*s = append(*s, str)
}

// Удаление элемента с вершины стека и возвращение Top элемента
func (s *Stack) pop() rune {
	sizeStack := len(*s)

	var topElement rune

	if sizeStack > 0 {
		lastIndex := sizeStack - 1
		topElement = (*s)[lastIndex]
		*s = (*s)[:lastIndex]
	}
	return topElement
}

var containsSymbols = map[rune]rune{
	')': '(',
	'}': '{',
	']': '[',
}

var containsSymbolsOpen = []rune{
	'(',
	'{',
	'[',
}

var containsSymbolsClose = []rune{
	')',
	'}',
	']',
}

func Balance(s string) bool {
	stack := make(Stack, 0)
	for _, str := range s {
		if slices.Contains(containsSymbolsOpen, str) {
			stack.push(str)
			continue
		}

		if slices.Contains(containsSymbolsClose, str) {
			symbolOpen := containsSymbols[str]
			top := stack.pop()

			if top != symbolOpen {
				return false
			}
		}
	}

	return stack.pop() == 0

}
