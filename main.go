package main

import (
	"fmt"
)

// A.1 Remember: if in starting state or any error state, clear out the lexemeHolder
// BUT DONT throw away the new character
// A.2 this case shouldnt happen WHEN note A.1 is fulfilled ???
// B.1 variable mutation does not get reflected unless return immediately in nested case > if
// C.1 Handle termination state. Should output lexemeHolder and reassign state to start
// C.2 Exp state is still an error

func main() {
	// b, err := ioutil.ReadFile("test/input/samp1.txt") // just pass the file name
	// if err != nil {
	// 	fmt.Print(err)
	// }

	// str := string(b) // convert content to a 'string'
	// for _, r := range str {
	// 	c := string(r)
	// 	fmt.Println(c)
	// }

	// char := `'`
	// lexemeHolder := ""

	// currentState, lexemeHolder = stateMachine(currentState, char, lexemeHolder)
	// fmt.Println("State: ", currentState, "Lexeme: ", lexemeHolder)
	// char = "a"
	// currentState, lexemeHolder = stateMachine(currentState, char, lexemeHolder)
	// fmt.Println("State: ", currentState, "Lexeme: ", lexemeHolder)
	// char = `"`
	// currentState, lexemeHolder = stateMachine(currentState, char, lexemeHolder)
	// fmt.Println("State: ", currentState, "Lexeme: ", lexemeHolder)
	// char = `'`
	// currentState, lexemeHolder = stateMachine(currentState, char, lexemeHolder)
	// fmt.Println("State: ", currentState, "Lexeme: ", lexemeHolder)
	// char = "*"
	// lexemeHolder = ""
	// currentState, lexemeHolder = stateMachine(currentState, char, lexemeHolder)
	// fmt.Println("State: ", currentState, "Lexeme: ", lexemeHolder)
	// char = "*"
	// currentState, lexemeHolder = stateMachine(currentState, char, lexemeHolder)
	// fmt.Println("State: ", currentState, "Lexeme: ", lexemeHolder)
	// str := `*`
	// c, l := stateMachineLooper(str)
	str := `+`
	c, l := stateMachineLooper(str)
	fmt.Println("State: ", c, "Lexeme: ", l)

}

func stateMachineLooper(s string) (state, string) {
	lexemeHolder := ""
	currentState := start
	prevState := start
	for currentState != terminated {
		for _, char := range s {
			prevState = currentState
			currentState, lexemeHolder = stateMachine(currentState, string(char), lexemeHolder)
		}
	}
	return prevState, lexemeHolder[:len(lexemeHolder)-1]
}

func stateMachine(currentState state, char string, lexemeHolder string) (state, string) {
	var newState state
	switch currentState {
	case start:
		switch {
		case isLetter(char):
			newState = ident
		case isNumber(char):
			newState = number
		case isSpace(char):
			newState = start
		case isComma(char):
			newState = comma
		case isPlus(char):
			newState = plus
		case isMinus(char):
			newState = minus
		case isDiv(char):
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
			newState = equal
		case isSingleQuote(char):
			newState = strSingle
		case isDoubleQuote(char):
			newState = strDouble
		case isStar(char):
			newState = mult
		default:
			newState = illegalCharacter
		}
	case ident:
		if isLetter(char) || isNumber(char) {
			newState = ident
		} else if isSpace(char) {
			newState = terminated
		} else {
			newState = illegalCharacter
		}
	case strSingle:
		if isSingleQuote(char) {
			newState = terminated
			lexemeHolder += char + `'`
			// B.1
			return newState, lexemeHolder
		}
		newState = strSingle
	case strDouble:
		if isDoubleQuote(char) {
			newState = terminated
			lexemeHolder += char + `"`
			// B.1
			return newState, lexemeHolder
		}
		newState = strDouble
	case mult:
		if isStar(char) {
			newState = exp
			return newState, lexemeHolder
		}
		newState = terminated
	case exp: //C.2
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
	case equal:
		newState = terminated
	case comma:
		newState = terminated
	case terminated:
		newState = start
		lexemeHolder = ""
	case illegalCharacter:
		fmt.Println("Illegal character in lexeme ", lexemeHolder)
		newState = illegalCharacter
	case unterminatedString: // A.2
		fmt.Println("Encountered EOF before string termination: ", lexemeHolder)
		newState = unterminatedString
	case badlyFormedNumber: // A.2
		fmt.Println("Badly formed number: ", lexemeHolder)
	default:
		newState = unknownState
		fmt.Println("Unhandled Case! Lexeme: ", lexemeHolder, " Last Recorded state: ", newState)
	}
	lexemeHolder += char
	return newState, lexemeHolder
}
