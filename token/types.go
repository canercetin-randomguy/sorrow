package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var Keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"add":    ADD,
}

const (
	VARIABLE = "VARIABLE"
	ILLEGAL  = "ILLEGAL"
	EOF      = "EOF"
	// IDENT & INT Identifiers + literals
	INT = "INT" // 1343456
	// ASSIGN & PLUS Operators
	ASSIGN = "="
	PLUS   = "+"

	// COMMA & and other Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	SPACE     = " "
	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"

	// Operators
	EQ     = "=="
	NOT_EQ = "!="
	ADD    = "ADD"
)
