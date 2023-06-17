package main

import (
	"fmt"
)

type TokenType int

const (
	// Single char tokens.
	LeftParen TokenType = iota
	RightParen
	LeftBrace
	RightBrace
	Comma
	Dot
	Plus
	Minus
	Star
	Semicolon
	Slash

	// One or two character tokens.
	Bang
	BangEqual
	Equal
	EqualEqual
	Greater
	GreaterEqual
	Less
	LessEqual

	// Literals.
	Identifier
	String
	Number

	// Keywords.
	And
	Class
	Else
	False
	Fun
	For
	If
	Nil
	Or
	Print
	Return
	Super
	This
	True
	Var
	While
)

var (
	ttypes = map[TokenType]string{
		LeftParen:    "LeftParen",
		RightParen:   "RightParen",
		LeftBrace:    "LeftBrace",
		RightBrace:   "RightBrace",
		Comma:        "Comma",
		Dot:          "Dot",
		Plus:         "Plus",
		Minus:        "Minus",
		Star:         "Star",
		Semicolon:    "Semicolon",
		Slash:        "Slash",
		Bang:         "Bang",
		BangEqual:    "BangEqual",
		Equal:        "Equal",
		EqualEqual:   "EqualEqual",
		Greater:      "Greater",
		GreaterEqual: "GreaterEqual",
		Less:         "Less",
		LessEqual:    "LessEqual",
		Identifier:   "Identifier",
		String:       "String",
		Number:       "Number",
		And:          "And",
		Class:        "Class",
		Else:         "Else",
		False:        "False",
		Fun:          "Fun",
		For:          "For",
		If:           "If",
		Nil:          "Nil",
		Or:           "Or",
		Print:        "Print",
		Return:       "Return",
		Super:        "Super",
		This:         "This",
		True:         "True",
		Var:          "Var",
		While:        "While",
	}
)

type token struct {
	tokenType TokenType
	lexeme    string
	literal   interface{}
	line      int
}

func (t token) debug() {
	fmt.Println("~~Token~~")
	fmt.Printf("Lexeme: %2s\n", t.lexeme)
	fmt.Printf("Literal: %2s\n", t.literal)
	fmt.Printf("Type: %2s\n", ttypes[t.tokenType])
	fmt.Printf("Line: %2d\n", t.line)
	fmt.Println()
}
