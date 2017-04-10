package main

import (
	"strconv"
	"unicode"
)

type state int

const (
	start state = iota
	// Accepting states
	ident
	strSingle
	strDouble
	number
	mult
	exp
	// terminal states
	plus
	minus
	div
	mod
	lParen
	rParen
	comma
	semicolon
	equal
	// Error States
	illegalCharacter
	unterminatedString
	badlyFormedNumber
	unknownState
)

// single character analyzers
func isLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return true
}

func isPlus(s string) bool {
	if s != "+" {
		return false
	}
	return true
}

func isMinus(s string) bool {
	if s != "-" {
		return false
	}
	return true
}

func isStar(s string) bool {
	if s != "*" {
		return false
	}
	return true
}

func isDiv(s string) bool {
	if s != "/" {
		return false
	}
	return true
}

func isMod(s string) bool {
	if s != "%" {
		return false
	}
	return true
}

func isLParen(s string) bool {
	if s != "(" {
		return false
	}
	return true
}

func isRParen(s string) bool {
	if s != ")" {
		return false
	}
	return true
}

func isComma(s string) bool {
	if s != "," {
		return false
	}
	return true
}

func isSemicolon(s string) bool {
	if s != ";" {
		return false
	}
	return true
}

func isEqualSymbol(s string) bool {
	if s != "=" {
		return false
	}
	return true
}

func isSpace(s string) bool {
	if s != " " {
		return false
	}
	return true
}

func isSingleQuote(s string) bool {
	if s != `'` {
		return false
	}
	return true
}

func isDoubleQuote(s string) bool {
	if s != `"` {
		return false
	}
	return true
}

func isDot(s string) bool {
	if s != "." {
		return false
	}
	return true
}
