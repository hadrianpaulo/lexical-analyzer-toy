package main

import (
	"bufio"
	"fmt"
	"os"
)

//go:generate stringer -type=state
type state int

const (
	start state = iota
	// Accepting states
	IDENT
	strSingle
	strDouble
	NUMBER
	numberPeriod
	numberDecimal
	numberExp
	numberDecimalExp
	numberTerminal
	MULT
	EXP
	COMMENT
	PLUS
	MINUS
	DIV
	MOD
	LPAREN
	RPAREN
	COMMA
	SEMICOLON
	EQUALS
	PERIOD
	STRING
	WHITESPACE
	BANG
	// relational operator states
	EQ
	GT
	LT
	GTOREQ
	LTOREQ
	NOTEQ
	// termination state/signal
	terminated
	// reserved words
	IF
	SQRT
	PRINT
	// Error States
	ILLEGALCHARACTER
	UNTERMINATEDSTRING
	BADLYFORMEDNUMBER
	unknownState
)

func convertNumberState(s state) state {
	switch s {
	case numberDecimalExp:
		return NUMBER
	case numberExp:
		return NUMBER
	case numberTerminal:
		return NUMBER
	case numberDecimal:
		return NUMBER
	}
	return s
}

func writeToFile(fn string, tokens []state, lexemes []string) {
	f, err := os.Create(fn)
	if err != nil {
		fmt.Print(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)

	for i := range tokens {
		_, err = fmt.Fprintf(w, "%v %v\n", tokens[i], lexemes[i])
		check(err)
		w.Flush()
	}
	f.Sync()
}
