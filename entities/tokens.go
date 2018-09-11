package tokens

// Token is a defined struct where it is defined the token body
type Token struct {
	Type       string
	Value      string
	LineNumber int
	Character  int
}

const (
	ReservedWord     = "Palabra Reservada"
	ID               = "Identificador"
	LeftParenthesis  = "("
	RightParenthesis = ")"

	// Arithmetic Operators
	Addition      = "+"
	Substraction  = "-"
	Division      = "/"
	Multplication = "*"
	Power         = "^"

	// Relational Operators
	Equals = "="
	Less   = "<"
	Great  = ">"

	//Logical Operators
	Or  = "|"
	And = "&"
)

// SpecialChars for detecting special chars
var SpecialChars = map[string]bool{
	"!":  true,
	"#":  true,
	"$":  true,
	"%":  true,
	"&":  true,
	"*":  true,
	"+":  true,
	",":  true,
	"-":  true,
	".":  true,
	"/":  true,
	":":  true,
	";":  true,
	"<":  true,
	"=":  true,
	">":  true,
	"?":  true,
	"@":  true,
	"[":  true,
	"\\": true,
	"]":  true,
	"^":  true,
	"_":  true,
	"`":  true,
	"{":  true,
	"|":  true,
	"}":  true,
	"~":  true,
}
