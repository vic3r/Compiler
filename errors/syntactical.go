package errors

import (
	"fmt"
	"time"
)

// SyntacticalError is a body defined for
type SyntacticalError struct {
	Value           string
	Type            string
	NumberLine      int
	NumberCharacter int
}

var _ error = &SyntacticalError{}

func (l *SyntacticalError) Error() string {
	return fmt.Sprintf("Syntatical error in: %s of type: %s in line: %d at :%d Character. In time: %v", l.Value, l.Type, l.NumberLine, l.NumberCharacter, time.Now())
}
