package main

import (
	"strconv"
	"unicode"
)

//go:generate stringer -type=state
type state int

const (
	start state = iota
	// Accepting states
	ident
	strSingle
	strDouble
	number
	numberDecimal
	numberExp
	numberDecimalExp
	numberTerminal
	mult
	exp
	comment
	plus
	minus
	div
	mod
	lParen
	rParen
	comma
	semicolon
	equal
	// termination state/signal
	terminated
	// Error States
	illegalCharacter
	unterminatedString
	badlyFormedNumber
	unknownState
	period
	STRING
)

// single character analyzers
func isLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	if s == "" {
		return false
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

func isSlash(s string) bool {
	if s != "/" {
		return false
	}
	return true
}

func isE(s string) bool {
	if s == "E" || s == "e" {
		return true
	}
	return false
}

func isNewLine(s string) bool {
	if s != "\n" {
		return false
	}
	return true
}
