package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// identifiers, literals
	IDENT = "IDENT" //variables
	INT   = "INT"   // Integer value

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimeters
	OPAREN = "("
	CPAREN = ")"
	CBRACE = "}"
	OBRACE = "{"

	COMMA     = ","
	SEMICOLON = ";"

	// keywords
	LET      = "LET"
	FUNCTION = "FUNCTION"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {
	if tokenType, ok := keywords[ident]; ok {
		return tokenType
	}
	return IDENT
}
