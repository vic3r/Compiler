package main

import (
	"fmt"
	"io/ioutil"
	"os"

	s "github.com/golang-collections/collections/stack"
	u "github.com/vic3r/Compiler/entities"
	e "github.com/vic3r/Compiler/errors"
)

var symbols = make([]*u.Symbol, 0)

func generateSynError(synErr *e.SyntacticalError) error {
	data := make([]byte, 0)
	data = append(data, []byte(fmt.Sprintf("ID						|	type	|	data_type	|	value  |  line number  |  character number\n"))...)
	data = append(data, []byte(fmt.Sprintf("%s  |  %s  |  %d  |  %d\n", synErr.Type, synErr.Value, synErr.NumberLine, synErr.NumberCharacter))...)

	if err := ioutil.WriteFile("syntacticError.txt", data, os.ModePerm); err != nil {
		return err
	}
	return nil
}

func generateSymbolSyntacticTable() error {
	data := make([]byte, 0)
	data = append(data, []byte(fmt.Sprintf("ID						|	var_type	|	data_type	|	type	| value  |  line number  |  character number\n"))...)
	for _, v := range symbols {
		data = append(data, []byte(fmt.Sprintf("	%s	|		%s  |	%s  |  %s  |  %d  |  %d\n", v.VarType, v.TypeData, v.Type, v.Value, v.LineNumber, v.Character))...)
	}
	if err := ioutil.WriteFile("symbolSyntatic.txt", data, os.ModePerm); err != nil {
		return err
	}
	return nil
}

func validateParenthesis(word string) bool {
	charStack := s.New()
	if word == "" {
		return true
	}

	runes := []rune(word)
	for i := 0; i < len(runes); i++ {
		if runes[i] == '(' || runes[i] == '{' {
			charStack.Push(runes[i])
		} else {
			if charStack.Len() == 0 && (runes[i] == ')' || runes[i] == '}') {
				return false
			}
			if runes[i] == ')' || runes[i] == '}' {
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

func validateReturn() error {

	return nil
}

func bodyFunction() bool {
	for _, v := range dictionary {
		varType := ""
		if e := functions[v.Value]; e != nil {
			varType = "Funcion"
		} else {
			switch {
			case u.Addition == v.Value:
				varType = "Operador Aritmetico"
				break
			case u.And == v.Value:
				varType = "Operador logico"
				break
			case u.Division == v.Value:
				varType = "Operador Aritmetico"
				break
			case u.DotComma == v.Value:
				varType = ";"
				break
			case u.Equals == v.Value:
				varType = "Asignacion"
				break
			case u.Great == v.Value:
				varType = "Operador Relacional"
				break
			case u.LeftBracket == v.Value:
				varType = "{"
				break
			case u.LeftParenthesis == v.Value:
				varType = "("
				break
			case u.Less == v.Value:
				varType = "Operador Relacional"
				break
			case u.Multplication == v.Value:
				varType = "Operdor Aritmetico"
				break
			case u.Or == v.Value:
				varType = "Operador Logico"
				break
			case u.Power == v.Value:
				varType = "Operador Aritmetico"
				break
			case u.ReservedWord == v.Value:
				varType = "Real"
				break
			case u.RightBracket == v.Value:
				varType = "}"
				break
			case u.RightParenthesis == v.Value:
				varType = ")"
				break
			case u.Substraction == v.Value:
				varType = "Operador Aritmetico"
				break
			default:
				varType = "Identificador"
				break
			}
		}
		symbol := &u.Symbol{
			Type:       v.Type,
			VarType:    varType,
			TypeData:   v.Value,
			Value:      v.Value,
			LineNumber: v.LineNumber,
			Character:  v.Character,
		}
		symbols = append(symbols, symbol)
	}

	return true
}

func validateFunctions() *e.SyntacticalError {
	for i := 0; i < len(tokens); i++ {
		v := tokens[i]
		funToken := functions[v.Value]
		if funToken != nil {
			if tokens[i+1].Value != "(" {
				return &e.SyntacticalError{
					Value:           tokens[i+1].Value,
					Type:            tokens[i+1].Type,
					NumberLine:      tokens[i+1].LineNumber,
					NumberCharacter: tokens[i+1].Character,
				}
			}
			i++
			for i < len(tokens) {
				if tokens[i].Value == ")" {
					i++
					if tokens[i].Value != u.LeftBracket {
						return &e.SyntacticalError{
							Value:           tokens[i].Value,
							Type:            tokens[i].Type,
							NumberLine:      tokens[i].LineNumber,
							NumberCharacter: tokens[i].Character,
						}
					}

					for i < len(tokens) {
						if tokens[i].Value == u.RightBracket {
							return nil
						}
						if e := dictionary[tokens[i].Value]; e != nil {
							elem := functions[tokens[i].Value]

							symbol := &u.Symbol{
								Type:       tokens[i].Type,
								LineNumber: tokens[i].LineNumber,
								Character:  tokens[i].Character,
							}
							if elem != nil {
								symbol.TypeData = elem.Type
							} else {
								symbol.TypeData = u.ID
							}
						}
						symbols = append(symbols)
						i++
					}

				}
				if tokens[i].Type == u.ReservedWord || tokens[i].Type == u.LeftBracket || tokens[i].Type == u.RightBracket || tokens[i].Type == u.DotComma || u.SpecialChars[tokens[i].Value] {
					return &e.SyntacticalError{
						Value:           tokens[i].Value,
						Type:            tokens[i].Type,
						NumberLine:      tokens[i].LineNumber,
						NumberCharacter: tokens[i].Character,
					}
				}
				i++
			}
		}
	}
	return nil
}

func validateParameters() error {
	return nil
}

func validatePrincipal() error {

	return nil
}

func synAnalysis() bool {
	bodyFunction()
	if synErr := validateFunctions(); synErr != nil {
		if e := generateSynError(synErr); e != nil {
			fmt.Printf("Error creating symbols text file in Syn Analysis: %v", synErr)
		}
		return false
	}
	return true
}
