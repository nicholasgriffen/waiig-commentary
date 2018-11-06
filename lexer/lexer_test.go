package lexer 

import (
	"testing"

	"interpreter/token"
)

func TestNextToken(test *testing.T) {
	input := `=+,;(){}`

	tests := []struct {
		expectedType	token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.LROUND, "("},
		{token.RROUND, ")"},
		{token.LCURLY, "{"},
		{token.RCURLY, "}"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tokenType := range tests {
		token := l.NextToken()

		if token.Type != tokenType.expectedType {
			test.Fatalf("tests[%d] - expected %q got tokenType %q", i, tokenType.expectedType, token.Type)
		}

		if token.Literal != tokenType.expectedLiteral {
			test.Fatalf("tests[%d] - expected %q got literal %q", i, tokenType.expectedLiteral, token.Literal)
		}
		
	}
}