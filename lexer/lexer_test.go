package lexer

import "testing"

func TestScanNext(t *testing.T) {
	input := `let message string = "hello world"`

	//	let message := "hello world"
	//	let ten int = 10
	//
	//	func sum(num1 int, num2 int) int {
	//		return num1 + num2
	//	}
	//	9 == 9
	//	(10 != 9)
	//	1 < 20
	//	3 > 30

	src := []byte(input)

	tests := []struct {
		expectedType   tokenType
		expectedValue  string
		expectedLine   int
		expectedColumn int
	}{
		{LET, "let", 1, 1},
		{IDENT, "message", 1, 5},
		{STRING, "string", 1, 13},
		{EQUAL, "=", 1, 15},
		{IDENT, "hello world", 1, 18},
	}

	lexer := NewLexer(src)

	for i, tt := range tests {
		tok := lexer.ScanNext()

		if tok.tokType != tt.expectedType {
			t.Fatalf("tests[%d] - tokType wrong. expected=%q, got=%q", i, tt.expectedType, tok.tokType)
		}
		if tok.value != tt.expectedValue {
			t.Fatalf("tests[%d] - tokType wrong. expected=%q, got=%q", i, tt.expectedType, tok.tokType)
		}
		if tok.position.line != tt.expectedLine {
			t.Fatalf("tests[%d] - tokType wrong. expected=%q, got=%q", i, tt.expectedType, tok.tokType)
		}
		if tok.position.column != tt.expectedColumn {
			t.Fatalf("tests[%d] - tokType wrong. expected=%q, got=%q", i, tt.expectedType, tok.tokType)
		}
	}
}
