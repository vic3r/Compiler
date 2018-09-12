package main

import s "github.com/golang-collections/collections/stack"

func validateParenthesis(word string) bool {
	charStack := s.New()
	if word == "" {
		return true
	}

	runes := []rune(word)
	for i := 0; i < len(runes); i++ {
		if runes[i] == '(' {
			charStack.Push(runes[i])
		} else {
			if charStack.Len() == 0 && runes[i] == ')' {
				return false
			}
			if runes[i] == ')' {
				charStack.Pop()
			}
		}
		characterNumber++
	}
	if charStack.Len() != 0 {
		return false
	}
	return true
}
