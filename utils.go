package main

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

func prettyPrint(s state) string {
	switch s {
	case numberDecimalExp:
		return "NUMBER"
	case numberExp:
		return "NUMBER"
	case numberTerminal:
		return "NUMBER"
	case numberDecimal:
		return "NUMBER"
	}
	return string(s.String())
}
