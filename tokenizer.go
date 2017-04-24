package main

func tokenizer(s string) (token state, lexeme string, remaining string) {
	lexemeHolder := ""
	currentState := start
	token = start
	counter := 0
	over := false
	for currentState != terminated {
		str := ""
		if len(s) > counter {
			str = string(s[counter])
		} else {
			over = true
		}
		token = currentState
		currentState, lexemeHolder = tokenIdentifier(currentState, str, lexemeHolder)
		counter++
	}
	remaining = ""
	if counter-1 <= len(s) {
		remaining = s[counter-1:]
	}
	if !over {
		lexemeHolder = lexemeHolder[:len(lexemeHolder)-1]
	}

	return token, lexemeHolder[:len(lexemeHolder)], remaining
}

func tokenIdentifier(currentState state, char string, lexemeHolder string) (state, string) {
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
		case isBang(char):
			newState = BANG
		case isGT(char):
			newState = GT
		case isLT(char):
			newState = LT
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
		if isEqualSymbol(char) {
			newState = EQ
		} else {
			newState = terminated
		}
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
	case EQ:
		newState = terminated
	case BANG:
		if isEqualSymbol(char) {
			newState = NOTEQ
		} else {
			newState = ILLEGALCHARACTER
		}
	case NOTEQ:
		newState = terminated
	case GT:
		if isEqualSymbol(char) {
			newState = GTOREQ
		} else {
			newState = terminated
		}
	case LT:
		if isEqualSymbol(char) {
			newState = LTOREQ
		} else {
			newState = terminated
		}
	case GTOREQ:
		newState = terminated
	case LTOREQ:
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
