package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	inputFile := os.Args[1]
	b, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)

	outputFile := os.Args[2]
	f, err := os.Create(outputFile)
	if err != nil {
		fmt.Print(err)
	}
	defer f.Close()
	c, l, d := stateMachineLooper(str)
	w := bufio.NewWriter(f)

	_, err = fmt.Fprintf(w, "%v %v\n", prettyPrint(c), l)
	check(err)
	w.Flush()
	for d != "" {
		c, l, d = stateMachineLooper(d)
		if c != WHITESPACE {
			_, err = fmt.Fprintf(w, "%v %v\n", prettyPrint(c), l)
			check(err)
			w.Flush()
		}
	}
	f.Sync()
}

func stateMachineLooper(s string) (state, string, string) {
	lexemeHolder := ""
	currentState := start
	prevState := start
	counter := 0
	over := false
	for currentState != terminated {
		str := ""
		if len(s) > counter {
			str = string(s[counter])
		} else {
			over = true
		}
		prevState = currentState
		currentState, lexemeHolder = stateMachine(currentState, str, lexemeHolder)
		counter++
	}
	rem := ""
	if counter-1 <= len(s) {
		rem = s[counter-1:]
	}
	if !over {
		lexemeHolder = lexemeHolder[:len(lexemeHolder)-1]
	}

	return prevState, lexemeHolder[:len(lexemeHolder)], rem
}

func stateMachine(currentState state, char string, lexemeHolder string) (state, string) {
	var newState state
	lexemeHolder += char
	switch currentState {
	case start:
		switch {
		case isLetter(char):
			newState = IDENT
		case isNumber(char):
			newState = NUMBER
		case isSpace(char) || isNewLine(char):
			newState = WHITESPACE
		case isComma(char):
			newState = COMMA
		case isPlus(char):
			newState = PLUS
		case isMinus(char):
			newState = MINUS
		case isSlash(char):
			newState = DIV
		case isMod(char):
			newState = MOD
		case isComma(char):
			newState = COMMA
		case isSemicolon(char):
			newState = SEMICOLON
		case isLParen(char):
			newState = LPAREN
		case isRParen(char):
			newState = RPAREN
		case isEqualSymbol(char):
			newState = EQUALS
		case isSingleQuote(char):
			newState = strSingle
		case isDoubleQuote(char):
			newState = strDouble
		case isStar(char):
			newState = MULT
		case isDot(char):
			newState = PERIOD
		case char == "":
			newState = terminated
		default:
			newState = ILLEGALCHARACTER
		}
	case IDENT:
		if isLetter(char) {
			newState = IDENT
			if isSQRT(lexemeHolder) {
				newState = SQRT
			} else if isPRINT(lexemeHolder) {
				newState = PRINT
			} else if isIF(lexemeHolder) {
				newState = IF
			}
		} else {
			newState = terminated
		}
	case NUMBER:
		if isNumber(char) {
			newState = NUMBER
		} else if isDot(char) {
			newState = numberPeriod
		} else if isE(char) {
			newState = numberExp
		} else {
			newState = terminated
		}
	case numberPeriod:
		if isNumber(char) {
			newState = numberDecimal
		} else {
			newState = BADLYFORMEDNUMBER
		}
	case numberDecimal:
		if isNumber(char) {
			newState = numberDecimal
		} else if isE(char) {
			newState = numberExp
		} else {
			newState = terminated
		}
	case numberExp:
		if isNumber(char) || isMinus(char) || isPlus(char) {
			newState = numberTerminal
		} else {
			newState = BADLYFORMEDNUMBER
		}
	case numberTerminal:
		if isNumber(char) {
			newState = numberTerminal
		} else {
			newState = terminated
		}
	case strSingle:
		if isSingleQuote(char) {
			newState = STRING
		} else if isNewLine(char) {
			newState = UNTERMINATEDSTRING
		} else {
			newState = strSingle
		}
	case strDouble:
		if isDoubleQuote(char) {
			newState = STRING
		} else if isNewLine(char) {
			newState = UNTERMINATEDSTRING
		} else {
			newState = strDouble
		}
	case MULT:
		if isStar(char) {
			newState = EXP
		} else {
			newState = terminated
		}
	case DIV:
		if isSlash(char) {
			newState = COMMENT
		} else {
			newState = terminated
		}
	case COMMENT:
		if isNewLine(char) {
			newState = terminated
		} else {
			newState = COMMENT
		}
	case EXP:
		newState = terminated
	case PLUS:
		newState = terminated
	case MINUS:
		newState = terminated
	case MOD:
		newState = terminated
	case SEMICOLON:
		newState = terminated
	case LPAREN:
		newState = terminated
	case RPAREN:
		newState = terminated
	case EQUALS:
		newState = terminated
	case COMMA:
		newState = terminated
	case PERIOD:
		newState = terminated
	case STRING:
		newState = terminated
	case WHITESPACE:
		newState = terminated
	case SQRT:
		newState = terminated
	case IF:
		newState = terminated
	case PRINT:
		newState = terminated
	case terminated:
		newState = start
		lexemeHolder = ""
	case ILLEGALCHARACTER:
		newState = terminated
	case UNTERMINATEDSTRING:
		newState = terminated
	case BADLYFORMEDNUMBER:
		newState = terminated
	default:
		newState = ILLEGALCHARACTER
	}
	return newState, lexemeHolder
}
