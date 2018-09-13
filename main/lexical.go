package main

import (
	"regexp"
	"strings"

	u "github.com/vic3r/Compiler/entities"
	e "github.com/vic3r/Compiler/errors"
	w "github.com/vic3r/Compiler/predefined"
)

const floatRegex = "^[0-9]{1,}.[0-9]{1,}$"

var frx = regexp.MustCompile(floatRegex)

func addTypeError(token *u.Token) {
	var tokenType string
	switch {
	case w.Set[token.Value]:
		tokenType = u.ReservedWord
		break
	case w.SetArithmetic[token.Value]:
		tokenType = "Operador Aritmetico"
		break
	case w.SetLogicalOperators[token.Value]:
		tokenType = "Operador Logico"
		break
	case w.SetRelationalOperators[token.Value]:
		tokenType = "Operador Relacional"
		break
	case token.Value == u.LeftParenthesis:
		tokenType = u.LeftParenthesis
		break
	case token.Value == u.RightParenthesis:
		tokenType = u.RightParenthesis
		break
	case token.Value == u.LeftBracket:
		tokenType = u.LeftBracket
		break
	case token.Value == u.RightBracket:
		tokenType = u.RightBracket
		break
	case token.Value == "[":
		tokenType = "["
		break
	case token.Value == "]":
		tokenType = "]"
		break
	default:
		tokenType = u.ID
		break
	}

	token.Type = tokenType
}

func validateToken(token *u.Token) *e.LexicalError {
	addTypeError(token)
	if token.Type == "Identificador" {
		if frx.MatchString(token.Value) {
			return nil
		}
		if token.Value == "===" {
			token.Value = "=="
			return nil
		}
		for k := range u.SpecialChars {
			if strings.Contains(token.Value, k) {
				index := strings.Index(token.Value, k) - token.Character
				return &e.LexicalError{
					Value:           token.Value,
					Type:            token.Type,
					NumberLine:      token.LineNumber,
					NumberCharacter: index,
				}
			}
		}
	}
	return nil
}
