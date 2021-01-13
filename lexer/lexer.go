package lexer

type Position struct {
	filename string
	line     int
	column   int
}

type Lexer struct {
	src          []byte
	position     int
	nextPosition int
	ch           byte
}

func NewLexer(src []byte) *Lexer {
	return &Lexer{src: src}
}

func (lex *Lexer) ScanNext() Token {
	var tok Token

}

func newToken(tokType tokenType, ch byte) Token {
	return Token{tokType: tokType, value: string(ch)}
}

// nextChar reads the next char and advances the position
func (lex *Lexer) nextChar() {
	if lex.nextPosition < len(lex.src) {
		lex.ch = lex.src[lex.nextPosition]
	} else {
		lex.ch = 0
	}

	lex.position = lex.nextPosition
	lex.nextPosition += 1
}

// peekChar reads the next char but does not advance the position
func (lex *Lexer) peekChar() byte {
	if lex.nextPosition < len(lex.src) {
		return lex.src[lex.nextPosition]
	} else {
		return 0
	}
}

//
//func (lex *Lexer) scanNumber() string {
//
//}
//
// skipWhitespace, scanNumber, scanString, etc
