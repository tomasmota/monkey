package lexer

import (
	"monkey/token"
)

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
    l := &Lexer{input: input}
    l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.position >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.position]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) nextToken() token.Token {
    var tok token.Token

    
    return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
    return token.Token{Type: tokenType, Literal: string(ch)}
}
