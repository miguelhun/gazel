package lexer

type tokenType int

const (
	ILLEGAL tokenType = iota
	EOF

	// data types
	IDENT
	INT
	DOUBLE
	STRING
	BOOL

	// arithmetic operators
	PLUS
	MINUS
	ASTERISK
	SLASH
	MODULO

	// assignment operators
	EQUAL
	PLUS_EQUAL
	MINUS_EQUAL
	ASTERISK_EQUAL
	SLASH_EQUAL
	NOT

	// comparison operators
	EQUAL_TO
	NOT_EQUAL
	LESS
	GREATER
	LESS_EQUAL
	GREATER_EQUAL
	AND
	OR

	// delimiters
	COMMA
	SEMICOLON
	COLON
	LPAREN
	RPAREN
	LBRACKET
	RBRACKET
	LBRACE
	RBRACE

	// keywords
	LET
	FUNC
	TRUE
	FALSE
	IF
	ELSE
	FOR
	WHILE
	RETURN
	TYPE
)

type Token struct {
	tokType  tokenType
	value    string
	position Position
}
