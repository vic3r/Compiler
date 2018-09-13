package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	u "github.com/vic3r/Compiler/entities"
	e "github.com/vic3r/Compiler/errors"
	w "github.com/vic3r/Compiler/predefined"
)

var (
	tokenID         = 0
	lineNumber      = 0
	characterNumber = 0
	reader          *bufio.Reader
	dictionary      = make(map[string]*u.Token)
	tokens          = make([]*u.Token, 0)
	errors          = make([]*e.LexicalError, 0)
	prevChar        = ""
	functions       = make(map[string]*u.Token)
)

func generateErrors() error {
	data := make([]byte, 0)
	data = append(data, []byte(fmt.Sprintf("ID						|  value  |  line number  |  character number\n"))...)
	for _, v := range errors {
		data = append(data, []byte(fmt.Sprintf("%s  |  %s  |  %d  |  %d\n", v.Type, v.Value, v.NumberLine, v.NumberCharacter))...)
	}
	if err := ioutil.WriteFile("errors.txt", data, os.ModePerm); err != nil {
		return err
	}
	return nil
}

func generateTokenTable() error {
	data := make([]byte, 0)
	data = append(data, []byte(fmt.Sprintf("ID						|  value  |  line number  |  character number\n"))...)
	for _, v := range tokens {
		data = append(data, []byte(fmt.Sprintf("%s  |  %s  |  %d  |  %d\n", v.Type, v.Value, v.LineNumber, v.Character))...)
	}
	if err := ioutil.WriteFile("tokens.txt", data, os.ModePerm); err != nil {
		return err
	}
	return nil
}

func generateSymbolTable() error {
	data := make([]byte, 0)
	data = append(data, []byte(fmt.Sprintf("ID						|  value  |  line number  |  character number\n"))...)
	for _, v := range dictionary {
		if v.Type == u.ID {
			data = append(data, []byte(fmt.Sprintf("%s  |  %s  |  %d  |  %d\n", v.Type, v.Value, v.LineNumber, v.Character))...)
		}
	}
	if err := ioutil.WriteFile("symbols.txt", data, os.ModePerm); err != nil {
		return err
	}
	return nil
}

func insertIntoSymbolsMap(token *u.Token) {
	var tokenType string
	switch {
	case w.Set[token.Value]:
		tokenType = token.Value
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
	if _, exists := dictionary[token.Value]; !exists {
		dictionary[token.Value] = &u.Token{
			Type:       tokenType,
			Value:      token.Value,
			LineNumber: lineNumber,
			Character:  characterNumber,
		}
	}
}

func insertIntoTokenMap(token *u.Token) {
	var tokenType string
	switch {
	case w.Set[token.Value]:
		tokenType = token.Value
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
	tokens = append(tokens, token)
}

func analyzeChar(c *rune, builder *strings.Builder) error {
	actualChar := string(*c)
	if actualChar != "\n" && actualChar != "\t" && actualChar != " " {
		if prevChar == u.DotComma && actualChar != "\n" {
			token := &u.Token{
				Value:      actualChar,
				LineNumber: lineNumber,
				Character:  characterNumber,
			}
			addTypeError(token)
			lexErr := &e.LexicalError{
				Type:            token.Type,
				Value:           token.Value,
				NumberLine:      token.LineNumber,
				NumberCharacter: token.Character,
			}
			errors = append(errors, lexErr)
		}
		if actualChar != u.DotComma && actualChar != u.LeftParenthesis && actualChar != u.RightParenthesis &&
			actualChar != u.LeftBracket && actualChar != u.RightBracket && actualChar != "!" &&
			actualChar != "," && actualChar != "\\" && actualChar != "&" && actualChar != "<" &&
			actualChar != ">" && actualChar != "/" && actualChar != "*" && actualChar != "-" && actualChar != "+" && actualChar != "^" {

			builder.WriteString(fmt.Sprintf("%s", actualChar))

		} else {
			if actualChar == u.DotComma {
				prevChar = actualChar
			}
			if actualChar == u.LeftParenthesis {
				tokFunc := &u.Token{
					Type:       "Function",
					Value:      builder.String(),
					LineNumber: lineNumber,
					Character:  characterNumber,
				}
				if _, v := functions[builder.String()]; !v {
					functions[builder.String()] = tokFunc
				}
			}
			if builder.String() != "" {
				token := &u.Token{Value: builder.String(), LineNumber: lineNumber, Character: characterNumber}

				if lexErr := validateToken(token); lexErr != nil {
					errors = append(errors, lexErr)
				} else {
					insertIntoTokenMap(token)
					insertIntoSymbolsMap(token)
					builder.Reset()
				}
			}

			token := &u.Token{Value: actualChar, LineNumber: lineNumber, Character: characterNumber}
			insertIntoTokenMap(token)
			insertIntoSymbolsMap(token)
			characterNumber++
		}
	} else {

		if actualChar == "\n" {
			lineNumber++
			characterNumber = 0
		}
		if builder.String() != "" {
			if builder.String() == "===" {
				token := &u.Token{Value: "==", LineNumber: lineNumber, Character: characterNumber}
				insertIntoTokenMap(token)
				insertIntoSymbolsMap(token)
				characterNumber++
				token = &u.Token{Value: "=", LineNumber: lineNumber, Character: characterNumber}
				insertIntoTokenMap(token)
				insertIntoSymbolsMap(token)
			} else {
				if builder.String() == "====" {
					token := &u.Token{Value: "==", LineNumber: lineNumber, Character: characterNumber}
					insertIntoTokenMap(token)
					insertIntoSymbolsMap(token)
					characterNumber += 2
					token = &u.Token{Value: "==", LineNumber: lineNumber, Character: characterNumber}
					insertIntoTokenMap(token)
					insertIntoSymbolsMap(token)
				} else {
					token := &u.Token{Value: builder.String(), LineNumber: lineNumber, Character: characterNumber}
					if lexErr := validateToken(token); lexErr != nil {
						errors = append(errors, lexErr)
					} else {
						insertIntoTokenMap(token)
						insertIntoSymbolsMap(token)
					}
				}
			}
		}

		builder.Reset()
	}
	prevChar = ""
	characterNumber++
	return nil
}

func generateFiles() error {
	if err := generateTokenTable(); err != nil {
		return fmt.Errorf("Tokens can not be generated %v", err)
	}
	if err := generateSymbolTable(); err != nil {
		return fmt.Errorf("Symbols can not be generated %v", err)
	}
	if err := generateErrors(); err != nil {
		return fmt.Errorf("Errors can not be generated %v", err)
	}
	return nil
}

func main() {
	file, _ := ioutil.ReadFile("./test.txt")

	text := string(file[:])

	reader = bufio.NewReader(strings.NewReader(text))
	var builder strings.Builder
	for {
		if c, _, err := reader.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		} else {
			if err := analyzeChar(&c, &builder); err != nil {
				fmt.Println("There was an error in the lexical Analysis")
			}
		}
	}
	if err := generateFiles(); err != nil {
		fmt.Println("Can not create tables for lexical Analysis")
		return
	}
	// if flag := synAnalysis(); flag {
	// 	fmt.Println("Can not do the syntactical analysis")
	// }

	if flag := synAnalysis(); flag {
		//fmt.Println("Can not do the syntactical analysis")
	}
	if err := generateSymbolSyntacticTable(); err != nil {
		fmt.Println("Can not generate syn table")
	}

}
