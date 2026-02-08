package lexer

import (
	"testing"

	"github.com/i-vinayak/interpreter-in-go/src/internal/token"
)

func TestNextToken(t *testing.T) {
	testInput := `
		let five = 5;
		let ten = 10;

		let add = fn(x, y){
			x + y
		};
		let result = add(five, ten);
	`

	tts := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.OPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.CPAREN, ")"},
		{token.OBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.CBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.OPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.CPAREN, ")"},
		{token.SEMICOLON, ";"},
	}

	l := New(testInput)

	for i, tt := range tts {
		curToken := l.NextToken()

		if curToken.Literal != tt.expectedLiteral {
			t.Fatalf("Failed test %d - expected %q, got %q", i, tt.expectedLiteral, curToken.Literal)
		}

		if curToken.Type != tt.expectedType {
			t.Fatalf("Failed test %d - expected %q, got %q", i, tt.expectedType, curToken.Type)
		}
	}
}
