package main

import "fmt"

// A.1 Remember: if in starting state or any error state, clear out the lexemeHolder
// BUT DONT throw away the new character
// A.2 this case shouldnt happen WHEN note A.1 is fulfilled ???
// B.1 variable mutation does not get reflected unless return immediately in nested case > if

func main() {
	char := `'`
	lexemeHolder := ""
	currentState := start
	currentState, lexemeHolder = stateMachine(currentState, char, lexemeHolder)
	fmt.Println("State: ", currentState, "Lexeme: ", lexemeHolder)
	char = "a"
	currentState, lexemeHolder = stateMachine(currentState, char, lexemeHolder)
	fmt.Println("State: ", currentState, "Lexeme: ", lexemeHolder)
	char = "d"
	currentState, lexemeHolder = stateMachine(currentState, char, lexemeHolder)
	fmt.Println("State: ", currentState, "Lexeme: ", lexemeHolder)
	char = `'`
	currentState, lexemeHolder = stateMachine(currentState, char, lexemeHolder)
	fmt.Println("State: ", currentState, "Lexeme: ", lexemeHolder)
	char = "*"
	lexemeHolder = ""
	currentState, lexemeHolder = stateMachine(currentState, char, lexemeHolder)
	fmt.Println("State: ", currentState, "Lexeme: ", lexemeHolder)
	char = "*"
	currentState, lexemeHolder = stateMachine(currentState, char, lexemeHolder)
	fmt.Println("State: ", currentState, "Lexeme: ", lexemeHolder)
	char = "("
	lexemeHolder = ""
	currentState, lexemeHolder = stateMachine(currentState, char, lexemeHolder)
	fmt.Println("State: ", currentState, "Lexeme: ", lexemeHolder)
	char = "1"
	lexemeHolder = ""
	currentState, lexemeHolder = stateMachine(currentState, char, lexemeHolder)
	fmt.Println("State: ", currentState, "Lexeme: ", lexemeHolder)
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
		case isMod(char):
			newState = mod
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
			newState = start
		} else {
			newState = illegalCharacter
		}
	case strSingle:
		if isSingleQuote(char) {
			newState = start
			lexemeHolder += char
			// B.1
			return newState, lexemeHolder
		}
		newState = strSingle
	case strDouble:
		if isDoubleQuote(char) {
			newState = start
			lexemeHolder += char
			// B.1
			return newState, lexemeHolder
		}
		newState = strDouble
	case mult:
		if isStar(char) {
			newState = exp
			lexemeHolder += char
			return newState, lexemeHolder
		}
		newState = start
	case exp:
		newState = start
	case plus:
		newState = start
	case minus:
		newState = start
	case mod:
		newState = start
	case semicolon:
		newState = start
	case lParen:
		newState = start
	case rParen:
		newState = start
	case equal:
		newState = start
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
