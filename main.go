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
		if c != whitespace {
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
			newState = ident
		case isNumber(char):
			newState = number
		case isSpace(char) || isNewLine(char):
			newState = whitespace
		case isComma(char):
			newState = comma
		case isPlus(char):
			newState = plus
		case isMinus(char):
			newState = minus
		case isSlash(char):
			newState = div
		case isMod(char):
			newState = mod
		case isComma(char):
			newState = comma
		case isSemicolon(char):
			newState = semicolon
		case isLParen(char):
			newState = lParen
		case isRParen(char):
			newState = rParen
		case isEqualSymbol(char):
			newState = equals
		case isSingleQuote(char):
			newState = strSingle
		case isDoubleQuote(char):
			newState = strDouble
		case isStar(char):
			newState = mult
		case isDot(char):
			newState = period
		case char == "":
			newState = terminated
		default:
			newState = illegalCharacter
		}
	case ident:
		if isLetter(char) {
			newState = ident
		} else {
			newState = terminated
		}
	case number:
		if isNumber(char) {
			newState = number
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
			newState = badlyFormedNumber
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
			newState = badlyFormedNumber
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
			newState = unterminatedString
		} else {
			newState = strSingle
		}
	case strDouble:
		if isDoubleQuote(char) {
			newState = STRING
		} else if isNewLine(char) {
			newState = unterminatedString
		} else {
			newState = strDouble
		}
	case mult:
		if isStar(char) {
			newState = exp
		} else {
			newState = terminated
		}
	case div:
		if isSlash(char) {
			newState = comment
		} else {
			newState = terminated
		}
	case comment:
		if isNewLine(char) {
			newState = terminated
		} else {
			newState = comment
		}
	case exp:
		newState = terminated
	case plus:
		newState = terminated
	case minus:
		newState = terminated
	case mod:
		newState = terminated
	case semicolon:
		newState = terminated
	case lParen:
		newState = terminated
	case rParen:
		newState = terminated
	case equals:
		newState = terminated
	case comma:
		newState = terminated
	case period:
		newState = terminated
	case STRING:
		newState = terminated
	case whitespace:
		newState = terminated
	case terminated:
		newState = start
		lexemeHolder = ""
	case illegalCharacter:
		newState = terminated
	case unterminatedString:
		newState = terminated
	case badlyFormedNumber:
		newState = terminated
	default:
		newState = illegalCharacter
	}
	return newState, lexemeHolder
}
