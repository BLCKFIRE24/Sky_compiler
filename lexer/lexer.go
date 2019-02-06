Package lexer 

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

func (l *Lexer) readChar() {
	if l.readPosition >= len(.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = 1.readPosition
	l.readPosition += 1
}