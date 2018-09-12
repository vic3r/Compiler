package errors

import (
	"fmt"
	"time"
)

// LexicalError is a body defined for
type LexicalError struct {
	Value           string
	Type            string
	NumberLine      int
	NumberCharacter int
}

var _ error = &LexicalError{}

func (l *LexicalError) Error() string {
	return fmt.Sprintf("Lexical error in: %s of type: %s in line: %d at :%d Character. In time: %v", l.Value, l.Type, l.NumberLine, l.NumberCharacter, time.Now())
}
