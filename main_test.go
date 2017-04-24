package main

import (
	"testing"
)

func TestIdentifySingleQuoteString(t *testing.T) {
	str := `'sda'`
	c, l, d := tokenizer(str)
	if c != STRING || l != str {
		t.Error(`Expected strSingle: 'sda', got `, c, ": ", l, " :rem: ", d)
	}
}

func TestIdentifyDoubleQuoteString(t *testing.T) {
	str := `"sda"`
	c, l, d := tokenizer(str)
	if c != STRING || l != str {
		t.Error(`Expected strDouble: "sda", got `, c, ": ", l, " :rem: ", d)
	}
}

func TestIdentifyPlus(t *testing.T) {
	str := `+`
	c, l, d := tokenizer(str)
	if c != PLUS || l != str {
		t.Error(`Expected plus: `+str+`, got `, c, ": ", l, " :rem: ", d)
	}
}

func TestIdentifyMinus(t *testing.T) {
	str := `-`
	c, l, d := tokenizer(str)
	if c != MINUS || l != str {
		t.Error(`Expected: minus`+str+`, got `, c, ": ", l, " :rem: ", d)
	}
}

func TestIdentifyMult(t *testing.T) {
	str := `*`
	c, l, d := tokenizer(str)
	if c != MULT || l != str {
		t.Error(`Expected: mult`+str+`, got `, c, ": ", l, " :rem: ", d)
	}
}

func TestIdentifyMod(t *testing.T) {
	str := `%`
	c, l, d := tokenizer(str)
	if c != MOD || l != str {
		t.Error(`Expected: mod`+str+`, got `, c, ": ", l, " :rem: ", d)
	}
}

func TestIdentifyLParen(t *testing.T) {
	str := `(`
	c, l, d := tokenizer(str)
	if c != LPAREN || l != str {
		t.Error(`Expected: lParen`+str+`, got `, c, ": ", l, " :rem: ", d)
	}
}

func TestIdentifyRParen(t *testing.T) {
	str := `)`
	c, l, d := tokenizer(str)
	if c != RPAREN || l != str {
		t.Error(`Expected: rParen`+str+`, got `, c, ": ", l, " :rem: ", d)
	}
}

func TestIdentifyComma(t *testing.T) {
	str := `,`
	c, l, d := tokenizer(str)
	if c != COMMA || l != str {
		t.Error(`Expected: comma`+str+`, got `, c, ": ", l, " :rem: ", d)
	}
}

func TestIdentifySemicolon(t *testing.T) {
	str := `;`
	c, l, d := tokenizer(str)
	if c != SEMICOLON || l != str {
		t.Error(`Expected: semicolon`+str+`, got `, c, ": ", l, " :rem: ", d)
	}
}

func TestIdentifyEqual(t *testing.T) {
	str := `=`
	c, l, d := tokenizer(str)
	if c != EQUALS || l != str {
		t.Error(`Expected: equal`+str+`, got `, c, ": ", l, " :rem: ", d)
	}
}

func TestIdentifyIdent(t *testing.T) {
	str := `sad `
	c, l, d := tokenizer(str)
	if c != IDENT || l != str[:len(str)-1] {
		t.Error(`Expected: ident`+str+`, got `, c, ": ", l, " :rem: ", d)
	}
}

func TestIdentifyNormalNumber(t *testing.T) {
	str := `123`
	c, l, d := tokenizer(str)
	if c != NUMBER || l != str {
		t.Error(`Expected: number`+str+`, got `, c, ": ", l, " :rem: ", d)
	}
}

func TestIdentifyDecimalNumber(t *testing.T) {
	str := `123.2`
	c, l, d := tokenizer(str)
	if c != numberDecimal || l != str {
		t.Error(`Expected: numberDecimal`+str+`, got `, c, ": ", l, " :rem: ", d)
	}
}

func TestIdentifyExpNumber(t *testing.T) {
	str := `123e12`
	c, l, d := tokenizer(str)
	if c != numberTerminal || l != str {
		t.Error(`Expected: numberTerminal`+str+`, got `, c, ": ", l, " :rem: ", d)
	}
	str1 := `123E12`
	c1, l1, d1 := tokenizer(str1)
	if c1 != numberTerminal || l1 != str1 {
		t.Error(`Expected: numberTerminal `+str1+`, got `, c1, ": ", l1, " :rem: ", d1)
	}
}

func TestIdentifyDecimalExpNumber(t *testing.T) {
	str := `123.12e1`
	c, l, d := tokenizer(str)
	if c != numberTerminal || l != str {
		t.Error(`Expected: numberTerminal`+str+`, got `, c, ": ", l, " :rem: ", d)
	}
	str1 := `123.123E124`
	c1, l1, d1 := tokenizer(str1)
	if c1 != numberTerminal || l1 != str1 {
		t.Error(`Expected: number`+str1+`, got `, c1, ": ", l1, " :rem: ", d1)
	}
}

func TestIdentifyDecimalExpMinusNumber(t *testing.T) {
	str := `123.12e-12`
	c, l, d := tokenizer(str)
	if c != numberTerminal || l != str {
		t.Error(`Expected: number`+str+`, got `, c, ": ", l, " :rem: ", d)
	}
}

func TestIdentifyExp(t *testing.T) {
	str := `**`
	c, l, d := tokenizer(str)
	if c != EXP || l != str {
		t.Error(`Expected exp: **, got `, c, ": ", l, " :rem: ", d)
	}
}

func TestIdentifyDiv(t *testing.T) {
	str := `/`
	c, l, d := tokenizer(str)
	if c != DIV || l != str {
		t.Error(`Expected div:`+str+`, got `, c, ": ", l, " :rem: ", d)
	}
}

func TestIdentifyComment(t *testing.T) {
	str := "// dsadas \n"
	c, l, d := tokenizer(str)
	if c != COMMENT || l != str[:len(str)-1] {
		t.Error(`Expected comment:`+str+`, got `, c, ": ", l, " :rem: ", d)
	}
}

func TestIdentifyBadNormalNumber(t *testing.T) {
	str := `12..`
	c, l, d := tokenizer(str)
	if c != BADLYFORMEDNUMBER || l != str {
		t.Error(`Expected: badlyFormedNumber`+str+`, got `, c, ": ", l, " :rem: ", d)
	}
}

func TestIdentifyBadDecimalNumber(t *testing.T) {
	str := `123.=`
	c, l, d := tokenizer(str)
	if c != BADLYFORMEDNUMBER || l != str {
		t.Error(`Expected: badlyFormedNumber`+str+`, got `, c, ": ", l, " :rem: ", d)
	}
}

func TestIdentifyBadExpNumber(t *testing.T) {
	str := `123e=`
	c, l, d := tokenizer(str)
	if c != BADLYFORMEDNUMBER || l != str {
		t.Error(`Expected: badlyFormedNumber`+str+`, got `, c, ": ", l, " :rem: ", d)
	}
	str1 := `123E.`
	c1, l1, d1 := tokenizer(str1)
	if c1 != BADLYFORMEDNUMBER || l1 != str1 {
		t.Error(`Expected: badlyFormedNumber `+str1+`, got `, c1, ": ", l1, " :rem: ", d1)
	}
}

func TestIdentifyIllegalCharacter(t *testing.T) {
	str := `^`
	c, l, d := tokenizer(str)
	if c != ILLEGALCHARACTER || l != str {
		t.Error(`Expected: illegal character`+str+`, got `, c, ": ", l, " :rem: ", d)
	}
}

func TestPrettyPrintNumber(t *testing.T) {
	c := prettyPrint(numberDecimalExp)
	if c != "NUMBER" {
		t.Error(`Expected: NUMBER, got `, c)
	}
}

func TestPrettyPrintOther(t *testing.T) {
	c := prettyPrint(IDENT)
	if c != "IDENT" {
		t.Error(`Expected: IDENT, got `, c)
	}
}

func TestIdentifyNewLine(t *testing.T) {
	str := "\n"
	c, l, d := tokenizer(str)
	if c != WHITESPACE || l != str {
		t.Error(`Expected: whitespace`+str+`, got `, c, ": ", l, " :rem: ", d)
	}
}

func TestIdentifyWhiteSpace(t *testing.T) {
	str := " "
	c, l, d := tokenizer(str)
	if c != WHITESPACE || l != str {
		t.Error(`Expected: whitespace`+str+`, got `, c, ": ", l, " :rem: ", d)
	}
}

func TestIdentifySQRT(t *testing.T) {
	str := "SQRT"
	c, l, d := tokenizer(str)
	if c != SQRT || l != str {
		t.Error(`Expected: SQRT `+str+`, got `, c, ": ", l, " :rem: ", d)
	}
}

func TestIdentifyPRINT(t *testing.T) {
	str := "PRINT"
	c, l, d := tokenizer(str)
	if c != PRINT || l != str {
		t.Error(`Expected: PRINT `+str+`, got `, c, ": ", l, " :rem: ", d)
	}
}

func TestIdentifyIF(t *testing.T) {
	str := "IF"
	c, l, d := tokenizer(str)
	if c != IF || l != str {
		t.Error(`Expected: IF `+str+`, got `, c, ": ", l, " :rem: ", d)
	}
}

func TestIdentifyEQ(t *testing.T) {
	str := "=="
	c, l, d := tokenizer(str)
	if c != EQ || l != str {
		t.Error(`Expected: EQ `+str+`, got `, c, ": ", l, " :rem: ", d)
	}
}

func TestIdentifyNOTEQ(t *testing.T) {
	str := "!="
	c, l, d := tokenizer(str)
	if c != NOTEQ || l != str {
		t.Error(`Expected: NOTEQ `+str+`, got `, c, ": ", l, " :rem: ", d)
	}
}

func TestIdentifyGT(t *testing.T) {
	str := ">"
	c, l, d := tokenizer(str)
	if c != GT || l != str {
		t.Error(`Expected: GT `+str+`, got `, c, ": ", l, " :rem: ", d)
	}
}

func TestIdentifyGTOREQ(t *testing.T) {
	str := ">="
	c, l, d := tokenizer(str)
	if c != GTOREQ || l != str {
		t.Error(`Expected: GTOREQ `+str+`, got `, c, ": ", l, " :rem: ", d)
	}
}

func TestIdentifyLT(t *testing.T) {
	str := "<"
	c, l, d := tokenizer(str)
	if c != LT || l != str {
		t.Error(`Expected: LT `+str+`, got `, c, ": ", l, " :rem: ", d)
	}
}

func TestIdentifyLTOREQ(t *testing.T) {
	str := "<="
	c, l, d := tokenizer(str)
	if c != LTOREQ || l != str {
		t.Error(`Expected: LTOREQ `+str+`, got `, c, ": ", l, " :rem: ", d)
	}
}
