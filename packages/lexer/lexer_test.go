package lexer

import (
	"monkey/packages/token"
	"testing"
)

type Lexer struct {
    input string
    position int
    readPosition int
    ch byte
    // 34r left off here
}

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType   token.TokenType
		expecteLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
	}

    l := New(input)

    for i, tt := range tests {
        tok := l.nextToken()

        if tok.Type != tt.expectedType {
            t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
        }

        if tok.Type != tt.expecteLiteral {
            t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expecteLiteral, tok.Literal)
        }
    }

}
