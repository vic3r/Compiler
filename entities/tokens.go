package tokens

// Token is a defined struct where it is defined the token body
type Token struct {
	ID         int
	Value      string
	LineNumber int
	Symbol     string
}
