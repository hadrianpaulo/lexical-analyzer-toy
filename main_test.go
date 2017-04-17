package main

import (
	"testing"
)

func TestIdentifySingleQuoteString(t *testing.T) {
	str := `'sda'`
	c, l, d := stateMachineLooper(str)
	if c != STRING || l != str {
		t.Error(`Expected strSingle: 'sda', got `, c, ": ", l, " :rem: ", d)
	}
}

func TestIdentifyDoubleQuoteString(t *testing.T) {
	str := `"sda"`
	c, l, d := stateMachineLooper(str)
	if c != STRING || l != str {
		t.Error(`Expected strDouble: "sda", got `, c, ": ", l, " :rem: ", d)
	}
}

func TestIdentifyPlus(t *testing.T) {
	str := `+`
	c, l, d := stateMachineLooper(str)
	if c != plus || l != str {
		t.Error(`Expected plus: `+str+`, got `, c, ": ", l, " :rem: ", d)
	}
}

func TestIdentifyMinus(t *testing.T) {
	str := `-`
	c, l, d := stateMachineLooper(str)
	if c != minus || l != str {
		t.Error(`Expected: minus`+str+`, got `, c, ": ", l, " :rem: ", d)
	}
}

func TestIdentifyMult(t *testing.T) {
	str := `*`
	c, l, d := stateMachineLooper(str)
	if c != mult || l != str {
		t.Error(`Expected: mult`+str+`, got `, c, ": ", l, " :rem: ", d)
	}
}

func TestIdentifyMod(t *testing.T) {
	str := `%`
	c, l, d := stateMachineLooper(str)
	if c != mod || l != str {
		t.Error(`Expected: mod`+str+`, got `, c, ": ", l, " :rem: ", d)
	}
}

func TestIdentifyLParen(t *testing.T) {
	str := `(`
	c, l, d := stateMachineLooper(str)
	if c != lParen || l != str {
		t.Error(`Expected: lParen`+str+`, got `, c, ": ", l, " :rem: ", d)
	}
}

func TestIdentifyRParen(t *testing.T) {
	str := `)`
	c, l, d := stateMachineLooper(str)
	if c != rParen || l != str {
		t.Error(`Expected: rParen`+str+`, got `, c, ": ", l, " :rem: ", d)
	}
}

func TestIdentifyComma(t *testing.T) {
	str := `,`
	c, l, d := stateMachineLooper(str)
	if c != comma || l != str {
		t.Error(`Expected: comma`+str+`, got `, c, ": ", l, " :rem: ", d)
	}
}

func TestIdentifySemicolon(t *testing.T) {
	str := `;`
	c, l, d := stateMachineLooper(str)
	if c != semicolon || l != str {
		t.Error(`Expected: semicolon`+str+`, got `, c, ": ", l, " :rem: ", d)
	}
}

func TestIdentifyEqual(t *testing.T) {
	str := `=`
	c, l, d := stateMachineLooper(str)
	if c != equal || l != str {
		t.Error(`Expected: equal`+str+`, got `, c, ": ", l, " :rem: ", d)
	}
}

func TestIdentifyIdent(t *testing.T) {
	str := `sad1 `
	c, l, d := stateMachineLooper(str)
	if c != ident || l != str[:len(str)-1] {
		t.Error(`Expected: ident`+str+`, got `, c, ": ", l, " :rem: ", d)
	}
}

func TestIdentifyNormalNumber(t *testing.T) {
	str := `123`
	c, l, d := stateMachineLooper(str)
	if c != number || l != str {
		t.Error(`Expected: number`+str+`, got `, c, ": ", l, " :rem: ", d)
	}
}

func TestIdentifyDecimalNumber(t *testing.T) {
	str := `123.2`
	c, l, d := stateMachineLooper(str)
	if c != numberTerminal || l != str {
		t.Error(`Expected: numberDecimal`+str+`, got `, c, ": ", l, " :rem: ", d)
	}
}

func TestIdentifyExpNumber(t *testing.T) {
	str := `123e12`
	c, l, d := stateMachineLooper(str)
	if c != numberTerminal || l != str {
		t.Error(`Expected: numberTerminal`+str+`, got `, c, ": ", l, " :rem: ", d)
	}
	str1 := `123E12`
	c1, l1, d1 := stateMachineLooper(str1)
	if c1 != numberTerminal || l1 != str1 {
		t.Error(`Expected: numberTerminal `+str1+`, got `, c1, ": ", l1, " :rem: ", d1)
	}
}

// func TestIdentifyDecimalExpNumber(t *testing.T) {
// 	str := `123.12e1`
// 	c, l, d := stateMachineLooper(str)
// 	if c != numberTerminal || l != str {
// 		t.Error(`Expected: numberTerminal`+str+`, got `, c, ": ", l, " :rem: ", d)
// 	}
// 	str1 := `123.123E124`
// 	c1, l1, d1 := stateMachineLooper(str1)
// 	if c1 != numberTerminal || l1 != str1 {
// 		t.Error(`Expected: number`+str1+`, got `, c1, ": ", l1, " :rem: ", d1)
// 	}
// }

// func TestIdentifyDecimalExpMinusNumber(t *testing.T) {
// 	str := `123.12e-12`
// 	c, l, d := stateMachineLooper(str)
// 	if c != numberTerminal || l != str {
// 		t.Error(`Expected: number`+str+`, got `, c, ": ", l, " :rem: ", d)
// 	}
// }

func TestIdentifyExp(t *testing.T) {
	str := `**`
	c, l, d := stateMachineLooper(str)
	if c != exp || l != str {
		t.Error(`Expected exp: **, got `, c, ": ", l, " :rem: ", d)
	}
}

func TestIdentifyDiv(t *testing.T) {
	str := `/`
	c, l, d := stateMachineLooper(str)
	if c != div || l != str {
		t.Error(`Expected div:`+str+`, got `, c, ": ", l, " :rem: ", d)
	}
}

func TestIdentifyComment(t *testing.T) {
	str := "// dsadas \n"
	c, l, d := stateMachineLooper(str)
	if c != comment || l != str[:len(str)-1] {
		t.Error(`Expected comment:`+str+`, got `, c, ": ", l, " :rem: ", d)
	}
}

func TestIdentifyBadNormalNumber(t *testing.T) {
	str := `12..`
	c, l, d := stateMachineLooper(str)
	if c != badlyFormedNumber || l != str {
		t.Error(`Expected: badlyFormedNumber`+str+`, got `, c, ": ", l, " :rem: ", d)
	}
}

func TestIdentifyBadDecimalNumber(t *testing.T) {
	str := `123.=`
	c, l, d := stateMachineLooper(str)
	if c != badlyFormedNumber || l != str {
		t.Error(`Expected: badlyFormedNumber`+str+`, got `, c, ": ", l, " :rem: ", d)
	}
}

func TestIdentifyBadExpNumber(t *testing.T) {
	str := `123e=`
	c, l, d := stateMachineLooper(str)
	if c != badlyFormedNumber || l != str {
		t.Error(`Expected: badlyFormedNumber`+str+`, got `, c, ": ", l, " :rem: ", d)
	}
	str1 := `123E.`
	c1, l1, d1 := stateMachineLooper(str1)
	if c1 != badlyFormedNumber || l1 != str1 {
		t.Error(`Expected: badlyFormedNumber `+str1+`, got `, c1, ": ", l1, " :rem: ", d1)
	}
}

// func TestIdentifyBadDecimalExpNumber(t *testing.T) {
// 	str := `123.12e.`
// 	c, l, d := stateMachineLooper(str)
// 	if c != badlyFormedNumber || l != str {
// 		t.Error(`Expected: badlyFormedNumber`+str+`, got `, c, ": ", l, " :rem: ", d)
// 	}
// 	str1 := `123.123E1.`
// 	c1, l1, d1 := stateMachineLooper(str1)
// 	if c1 != badlyFormedNumber || l1 != str1 {
// 		t.Error(`Expected: badlyFormedNumber`+str1+`, got `, c1, ": ", l1, " :rem: ", d1)
// 	}
// }

// func TestIdentifyBadDecimalExpMinusNumber(t *testing.T) {
// 	str := `123e-1.`
// 	c, l, d := stateMachineLooper(str)
// 	if c != badlyFormedNumber || l != str {
// 		t.Error(`Expected: badlyFormedNumber`+str+`, got `, c, ": ", l, " :rem: ", d)
// 	}
// }

func TestIdentifyIllegalCharacter(t *testing.T) {
	str := `^`
	c, l, d := stateMachineLooper(str)
	if c != illegalCharacter || l != str {
		t.Error(`Expected: illegal character`+str+`, got `, c, ": ", l, " :rem: ", d)
	}
}

// func TestIdentifyMultiple(t *testing.T) {
// 	str := "123"
// 	c, l, d := stateMachineLooper(str)
// 	if c != comment || l != str {
// 		t.Error(`Expected comment:`+str+`, got `, c, ": ", l, " :rem: ", d)
// 	}
// }
