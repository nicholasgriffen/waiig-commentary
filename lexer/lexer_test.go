package lexer 

import (
	"testing"

	"interpreter/token"
)

func TestNextToken(test *testing.T) {
	input := `let six = 6;
	let nineteen = 19;
	
	let add = fn(x, y) {
			x + y;
		}
		
	let result = add(six, 19)`

	tests := []struct {
		expectedType	token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "six"},
		{token.ASSIGN, "="},
		{token.INT, "6"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "nineteen"},
		{token.ASSIGN, "="},
		{token.INT, "19"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LROUND, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RROUND, ")"},
		{token.SEMICOLON, ";"},
		{token.LCURLY, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.RCURLY, "}"},
		{token.LET, "let"},
		{token.IDENT, " result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},		
		{token.LROUND, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RROUND, ")"},
		{token.SEMICOLON, ";"},
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