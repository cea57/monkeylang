package token

type TokenType string

type Token struct {
	Type 	TokenType
	Literal string

}

const (
	ILLEGAL = "ILLEGAL"
	EOF 	= "EOF"

	// Identifiers + Literals
	IDENT = "IDENT" // add, foobar, x, y, z, ...
	INT   = "INT"

	// Operators
	ASSIGN   = "="
	PLUS 	 = "+"

	// Delimeters
	COMMA 	  = ","
	SEMICOLON = ";"
	LPAREN 	  = "("
	RPAREN 	  = ")"
	LBRACE 	  = "{"
	RBRACE 	  = "}"

	// Keywords 
	FUNCTION = "FUNCTION"
	LET 	 = "LET"
)
