package parser

import (
	"fmt"
	"toy/lexer"
	"toy/token"
)

// segura/guarda o estado do parser
type Parser struct {
	l         *lexer.Lexer
	errors    []string
	curToken  token.Token
	peekToken token.Token
}

// inicializa o parser
func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	// le dois token para inicializar curToken e peekToken
	p.nextToken()
	p.nextToken()
	return p
}

// nextToken avança ambos curToken e peekToken
func (p *Parser) nextToken() {
	// o token atual recebe o caracter coletado pelo peekToken
	p.curToken = p.peekToken
	// enquanto peekToken ja avança para coletar o próximo token
	p.peekToken = p.l.NextToken()
}

// ParseProgram "parsea" o input por completo e a partir disso
// produz o node Program
func (p *Parser) ParseProgram() *Program {
	program := &Program{}
	program.Statements = []Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *LetStatement {
	stmt := &LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// enquanto o token atual não for SEMICOLON `;`
	for !p.curTokenIs(token.SEMICOLON) {
		// continuo avançando para o proximo token
		p.nextToken()
	}

	return stmt
}

// helpers
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}

}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) Errors() []string {
	return p.errors
}
