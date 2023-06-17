package main

import (
	"fmt"
	"log"
	"strconv"
	"unicode"
)

var (
	keywords = map[string]TokenType{
		"and":    And,
		"class":  Class,
		"else":   Else,
		"false":  False,
		"fun":    Fun,
		"for":    For,
		"if":     If,
		"nil":    Nil,
		"or":     Or,
		"print":  Print,
		"return": Return,
		"super":  Super,
		"this":   This,
		"true":   True,
		"var":    Var,
		"while":  While,
	}
)

type scanner struct {
	source  []rune
	current int
	start   int
	line    int
	tokens  []token
}

func newscanner(src string) *scanner {

	return &scanner{
		source:  []rune(src),
		current: 0,
		start:   0,
		line:    1,
		tokens:  []token{},
	}
}

func (s scanner) clip() string {
	return string(s.source[s.start:s.current])
}

func (s scanner) isatend() bool {
	return s.current >= len(s.source)
}

func (s *scanner) scantokens() []token {
	for !s.isatend() {
		s.start = s.current
		scantoken(s)
	}

	return s.tokens
}

func scantoken(s *scanner) {
	c := advance(s)

	switch c {
	case '(':
		addtoken(s, LeftParen, nil)
		break
	case ')':
		addtoken(s, RightParen, nil)
		break
	case '{':
		addtoken(s, LeftBrace, nil)
		break
	case '}':
		addtoken(s, RightBrace, nil)
		break
	case ',':
		addtoken(s, Comma, nil)
		break
	case '+':
		addtoken(s, Plus, nil)
		break
	case '-':
		addtoken(s, Minus, nil)
		break
	case '*':
		addtoken(s, Star, nil)
		break
	case ';':
		addtoken(s, Semicolon, nil)
		break
	case '.':
		addtoken(s, Dot, nil)
		break
	case '!':
		if match(s, '=') {
			addtoken(s, BangEqual, nil)
		} else {
			addtoken(s, Bang, nil)
		}
		break
	case '=':
		if match(s, '=') {
			addtoken(s, EqualEqual, nil)
		} else {
			addtoken(s, Equal, nil)
		}
		break
	case '<':
		if match(s, '=') {
			addtoken(s, LessEqual, nil)
		} else {
			addtoken(s, Less, nil)
		}
		break
	case '>':
		if match(s, '=') {
			addtoken(s, GreaterEqual, nil)
		} else {
			addtoken(s, Greater, nil)
		}
		break
	case '/':
		if match(s, '/') {
			for peek(s) != '\n' && !s.isatend() {
				advance(s)
			}
		} else {
			addtoken(s, Slash, nil)
		}
		break
	case ' ':
	case '\r':
	case '\t':
		break
	case '\n':
		s.line += 1
		break
	case '"':
		str(s)
		break
	default:
		if isdigit(c) {
			number(s, c)
		} else if isalpha(c) {
			identifier(s, c)
		} else {
			log.Print("Unexpected character: %s\n", string(c))
		}
	}
}

func isdigit(c rune) bool {
	return unicode.IsNumber(c)
}

func isalpha(c rune) bool {
	return unicode.IsLetter(c)
}

func identifier(s *scanner, c rune) {
	for isalpha(peek(s)) && !s.isatend() {
		advance(s)
	}

	lexeme := s.clip()
	if keywords[lexeme] != 0 {
		addtoken(s, keywords[lexeme], nil)
	} else {
		addtoken(s, Identifier, nil)
	}
}

func number(s *scanner, c rune) {
	for (isdigit(peek(s)) || peek(s) == '.') && !s.isatend() {
		advance(s)
	}

	lexeme := s.clip()

	v, err := strconv.ParseFloat(lexeme, 64)
	if err != nil {
		log.Print(err)
	}
	addtoken(s, Number, v)
}

func str(s *scanner) {
	for {
		if peek(s) == '"' || s.isatend() {
			break
		}
		if peek(s) == '\n' {
			s.line += 1
		}
		advance(s)
	}

	if s.isatend() {
		fmt.Errorf("Unterminated string.: line %d", s.line)
		return
	}

	// the last '"'
	advance(s)

	addtoken(s, String, s.clip())
}

func peek(s *scanner) rune {
	if s.isatend() {
		return 0
	}

	return s.source[s.current]
}

func peeknext(s *scanner) rune {
	if s.current+1 >= len(s.source) {
		return 0
	}

	return s.source[s.current+1]
}

func addtoken(s *scanner, ttype TokenType, literal interface{}) {
	t := token{
		ttype,
		s.clip(),
		literal,
		s.line,
	}
	s.tokens = append(s.tokens, t)
}

func advance(s *scanner) rune {
	c := s.source[s.current]
	s.current = s.current + 1
	return c
}

func match(s *scanner, c rune) bool {
	match := c == s.source[s.current]

	if match {
		advance(s)
	}

	return match
}
