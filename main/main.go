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
	dictionary      = make(map[string]*u.Token)
)

func generateTokens() error {
	data := make([]byte, 0)
	data = append(data, []byte(fmt.Sprintf("ID  |  value  |  line number  |  character number\n"))...)
	for _, v := range dictionary {
		data = append(data, []byte(fmt.Sprintf("%s  |  %s  |  %d  |  %d\n", v.Type, v.Value, v.LineNumber, v.Character))...)
	}
	if err := ioutil.WriteFile("tokens.txt", data, os.ModePerm); err != nil {
		return err
	}
	return nil
}

func analyzeChar(c *rune, builder *strings.Builder) {
	actualChar := string(*c)
	if actualChar != "\n" && actualChar != "\t" && actualChar != " " {
		builder.WriteString(fmt.Sprintf("%s", actualChar))
	} else {
		tokenID++
		// Add token
		token := &u.Token{Value: builder.String(), LineNumber: lineNumber}
		if _, v := w.Set[token.Value]; v {
			dictionary[token.Value] = &u.Token{
				Type:       u.ReservedWord,
				Value:      token.Value,
				LineNumber: lineNumber,
				Character:  characterNumber,
			}
		} else {
			if _, exists := dictionary[token.Value]; !exists {
				dictionary[token.Value] = &u.Token{
					Type:       u.ID,
					Value:      token.Value,
					LineNumber: lineNumber,
					Character:  characterNumber,
				}
			}
		}
		characterNumber++
		if actualChar == "\n" {
			lineNumber++
			characterNumber = 0
		}
		builder.Reset()
	}
}

func main() {
	file, _ := ioutil.ReadFile("./test.txt")

	text := string(file[:])

	r := bufio.NewReader(strings.NewReader(text))
	var builder strings.Builder
	for {
		if c, _, err := r.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		} else {
			analyzeChar(&c, &builder)
		}
	}
	if err := generateTokens(); err != nil {
		fmt.Printf("The map is not accessible %v", err)
	}
}
