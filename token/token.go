package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// Character we don't know
	ILLEGAL = "ILLEGAL"
	// Tells the parser to stop
	EOF = "EOF"

	// Identifiers + literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"
	TIMES  = "*"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	THEN      = "THEN"
	DO        = "DO"
	END       = "END"

	// Keywords
	SET      = "SET"
	FUNCTION = "FN"
	IF       = "IF"
	ELSEIF   = "ELSE IF"
	ELSE     = "ELSE"
	FOR      = "FOR"
)
