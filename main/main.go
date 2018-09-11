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
)

var (
	tokenID    = 0
	lineNumber = 0
	dictionary = make(map[string]interface{})
)

func iterateOverMap() error {
	data := make([]byte, 0)
	for k, v := range dictionary {
		data = append(data, []byte(fmt.Sprintf("%s %v\n", k, v))...)
	}
	if err := ioutil.WriteFile("table.txt", data, os.ModePerm); err != nil {
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
		token := &u.Token{ID: tokenID, Value: builder.String(), LineNumber: lineNumber}
		if _, exists := dictionary[token.Value]; !exists {
			dictionary[token.Value] = token.ID
		}
		if actualChar == "\n" {
			lineNumber++
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
	if err := iterateOverMap(); err != nil {
		fmt.Printf("The map is not accessible %v", err)
	}
}
