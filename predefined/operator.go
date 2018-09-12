package predefined

//SetArithmetic is a set of arithmetic operations
var SetArithmetic = map[string]bool{
	"=": true,
	"+": true,
	"-": true,
	"^": true,
	"/": true,
	"*": true,
}

//SetRelationalOperators is a set of arithmetic operations
var SetRelationalOperators = map[string]bool{
	"<":  true,
	">":  true,
	"==": true,
	"<=": true,
	">=": true,
}

// SetLogicalOperators is a set of logical operators
var SetLogicalOperators = map[string]bool{
	"&": true,
	"|": true,
	"!": true,
}
