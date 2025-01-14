package parser

import "toy/token"

// Isso aqui representa um nó na ast
type Node interface {
	TokenLiteral() string
}

// Statement representa "a complete statement" na ast
type Statement interface {
	Node
	statementNode()
}

// Expression representa uma expressão na ast
type Expression interface {
	Node
	expressionNode()
}

// Program é o nó raiz de toda ast
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}

	return ""
}

// LetStatement representa o `let`
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// representa um identificador na ast
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
