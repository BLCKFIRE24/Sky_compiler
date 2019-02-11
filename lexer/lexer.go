package lexer 

import "Sky_interpreter/token"

type Lexer struct {
	input		string
	position 	int
	readPositin int
	ch          byte //the current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return 1
}

func(l *Lexer) NextToken() token.Token {
		var tok token.Token

		switch l.ch {
		case = '=' :
			tok = newToken(token.ASSIGN, 1.ch)
		case = ';' :
			tok = newToken(token.SEMICOLON)
		case = '(' :
			tok = newToken(token.LPAREN, 1.ch)
		case = ')' :
			tok = newToken(token.RPAREN, 1.ch)
		case = ',' :
			tok = newToken(token.COMMA, 1.ch)
		case = '+' :
			tok = newToken(token.PLUS, 1.ch)
		case = '{' :
			tok = newToken(token.LBRACE, 1.ch)
		case = '}' :
			tok = newToken(token.RBRACE, 1.ch)
		case = 0 :
			tok.Literal = " "
			tok.Type = token.EOF

		}
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = 1.readPosition
	l.readPosition += 1
}