package parser

import (
	"testing"
	"toy/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foo = 838383;
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foo"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		letStmt, ok := stmt.(*LetStatement)
		if !ok {
			t.Fatalf("stmt not *LetStatement. got=%T", stmt)
		}

		if letStmt.Name.Value != tt.expectedIdentifier {
			t.Errorf("letStmt.Name.Value not '%s'. got=%s", tt.expectedIdentifier, letStmt.Name.Value)
		}

		if letStmt.Name.TokenLiteral() != tt.expectedIdentifier {
			t.Errorf("letStmt.Name.TokenLiteral() not '%s'. got=%s", tt.expectedIdentifier, letStmt.Name.TokenLiteral())
		}
	}
}
