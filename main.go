package main

import (
	"fmt"
	"os"
	"toy/lexer"
	"toy/parser"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: toy <source_file>")
		os.Exit(1)
	}

	// Read the source file
	filename := os.Args[1]
	source, err := os.ReadFile(filename)

	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", filename, err)
		os.Exit(1)
	}

	// Phase 1: Lexical Analysis
	fmt.Println("Phase 1: Lexical Analysis")
	l := lexer.New(string(source))
	for tok := l.NextToken(); tok.Type != "EOF"; tok = l.NextToken() {
		fmt.Printf("Token: {Type: %s, Literal: %q, Line: %d, Column: %d}\n", tok.Type, tok.Literal, tok.Line, tok.Column)
	}

	// Reinitialize lexer for parsing
	l = lexer.New(string(source))

	// Phase 2: Parsing
	fmt.Println("\nPhase 2: Parsing")
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		fmt.Println("Parser errors:")
		for _, err := range p.Errors() {
			fmt.Println("\t", err)
		}
		os.Exit(1)
	}

	fmt.Println("AST:")
	fmt.Println(program)
}
