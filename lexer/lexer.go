package lexer

type Position struct {
	filename string
	line     int
	column   int
}

type Lexer struct {
	src    []byte
	ch     byte
	nextCh byte
}

func NewLexer(src []byte) *Lexer {
	return &Lexer{src: src}
}
