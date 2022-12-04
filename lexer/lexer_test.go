package lexer_test

import (
	"sorrow/lexer"
	"sorrow/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	/* input := `let five=5
	let ten=10
	let add = fn(x,y){
		x+y;
	};
	let result=add(five,ten);` */
	input := `let five = 5;
              if (five > 10) {
				true;
			} else {
				false;
			}`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.EOF, " "},
		{token.VARIABLE, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
	}
	l := lexer.New(input)
	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
