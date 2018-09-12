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
	w "github.com/vic3r/Compiler/predefined"
)

var (
	tokenID         = 0
	lineNumber      = 0
	characterNumber = 0
	reader          *bufio.Reader
	dictionary      = make(map[string]*u.Token)
	tokens          = make([]*u.Token, 0)
)

func generateTokenTable() error {
	data := make([]byte, 0)
	data = append(data, []byte(fmt.Sprintf("ID  |  value  |  line number  |  character number\n"))...)
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
	data = append(data, []byte(fmt.Sprintf("ID  |  value  |  line number  |  character number\n"))...)
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
	if _, v := w.Set[token.Value]; v {
		tokenType = u.ReservedWord
	} else {
		tokenType = u.ID
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
	if _, v := w.Set[token.Value]; v {
		tokenType = u.ReservedWord
	} else {
		tokenType = u.ID
	}
	token.Type = tokenType
	tokens = append(tokens, token)
}

func analyzeChar(c *rune, builder *strings.Builder) error {
	actualChar := string(*c)
	if actualChar != "\n" && actualChar != "\t" && actualChar != " " {
		if actualChar != u.DotComma && actualChar != u.LeftParenthesis && actualChar != u.RightParenthesis && actualChar != u.LeftBracket && actualChar != u.RightBracket && actualChar != "!" {
			builder.WriteString(fmt.Sprintf("%s", actualChar))
		} else {
			if builder.String() != " " {
				token := &u.Token{Value: builder.String(), LineNumber: lineNumber, Character: characterNumber}
				insertIntoTokenMap(token)
				insertIntoSymbolsMap(token)
				builder.Reset()
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
		if builder.String() != " " {
			token := &u.Token{Value: builder.String(), LineNumber: lineNumber, Character: characterNumber}
			insertIntoTokenMap(token)
			insertIntoSymbolsMap(token)
		}

		builder.Reset()
	}
	characterNumber++
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
	if err := generateTokenTable(); err != nil {
		fmt.Printf("Tokens can not be generated %v", err)
	}
	if err := generateSymbolTable(); err != nil {
		fmt.Printf("Symbols can not be generated %v", err)
	}
}
