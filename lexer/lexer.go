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
	position := l.position + 1 // skip opening quote
	for {
		l.readChar()
		if l.character == '"' || l.character == 0 {
			break
		}
	}
	str := l.input[position:l.position]
	l.readChar() // skip closing quote
	return str
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	tok.Line = l.line
	tok.Column = l.column

	switch l.character {
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case '=':
		tok.Type = token.ASSIGN
		tok.Literal = string(l.character)
	case '+':
		tok.Type = token.PLUS
		tok.Literal = string(l.character)
	case '-':
		tok.Type = token.MINUS
		tok.Literal = string(l.character)
	case '*':
		tok.Type = token.ASTERISK
		tok.Literal = string(l.character)
	case '/':
		tok.Type = token.SLASH
		tok.Literal = string(l.character)
	case '(':
		tok.Type = token.LPAREN
		tok.Literal = string(l.character)
	case ')':
		tok.Type = token.RPAREN
		tok.Literal = string(l.character)
	case '{':
		tok.Type = token.LBRACE
		tok.Literal = string(l.character)
	case '}':
		tok.Type = token.RBRACE
		tok.Literal = string(l.character)
	case ',':
		tok.Type = token.COMMA
		tok.Literal = string(l.character)
	case ';':
		tok.Type = token.SEMICOLON
		tok.Literal = string(l.character)
	case 0:
		tok.Type = token.EOF
		tok.Literal = ""
	default:
		if unicode.IsLetter(l.character) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok
		} else if unicode.IsDigit(l.character) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok.Type = token.ILLEGAL
			tok.Literal = string(l.character)
		}
	}

	l.readChar()
	return tok
}
