package lexer

import (
	"toy/token"
	"unicode"
	"unicode/utf8"
)

type Lexer struct {
	input        string
	position     int  // posição atual do input
	readPosition int  // lendo a posição atual do input
	character    rune // caracter atual sendo lido/examidado
	line         int  // linha atual
	column       int  // coluna atual
}

func New(input string) *Lexer {
	l := &Lexer{
		input:  input,
		line:   1,
		column: 0,
	}

	l.readChar() // inicializa o primeiro caracter
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.character = 0
	} else {
		var size int
		l.character, size = utf8.DecodeRuneInString(l.input[l.readPosition:])

		if l.character == '\n' {
			l.line++
			l.column = 0
		} else {
			l.column++
		}

		l.position = l.readPosition
		l.readPosition += size
	}
}

func (l *Lexer) peekChar() rune {
	if l.readPosition >= len(l.input) {
		return 0
	}

	r, _ := utf8.DecodeRuneInString(l.input[l.readPosition:])
	return r
}

// o nome da função já diz tudo
func (l *Lexer) skipWhitespace() {
	for unicode.IsSpace(l.character) {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position

	for unicode.IsLetter(l.character) || l.character == '_' {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position

	for unicode.IsDigit(l.character) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) readString() string {
