package main

import (
	"testing"
)

func TestIdentifySingleQuoteString(t *testing.T) {
	str := `'sda'`
	c, l := stateMachineLooper(str)
	if c != strSingle && l != str {
		t.Error(`Expected strSingle: 'sda', got `, c, ": ", l)
	}
}

func TestIdentifyDoubleQuoteString(t *testing.T) {
	str := `"sda"`
	c, l := stateMachineLooper(str)
	if c != strDouble && l != str {
		t.Error(`Expected strDouble: "sda", got `, c, ": ", l)
	}
}

func TestIdentifyPlus(t *testing.T) {
	str := `+`
	c, l := stateMachineLooper(str)
	if c != plus && l != str {
		t.Error(`Expected plus: `+str+`, got `, c, ": ", l)
	}
}

func TestIdentifyMinus(t *testing.T) {
	str := `-`
	c, l := stateMachineLooper(str)
	if c != minus && l != str {
		t.Error(`Expected: minus`+str+`, got `, c, ": ", l)
	}
}

func TestIdentifyMult(t *testing.T) {
	str := `*`
	c, l := stateMachineLooper(str)
	if c != mult && l != str {
		t.Error(`Expected: mult`+str+`, got `, c, ": ", l)
	}
}

func TestIdentifyMod(t *testing.T) {
	str := `%`
	c, l := stateMachineLooper(str)
	if c != mod && l != str {
		t.Error(`Expected: mod`+str+`, got `, c, ": ", l)
	}
}

func TestIdentifyLParen(t *testing.T) {
	str := `(`
	c, l := stateMachineLooper(str)
	if c != lParen && l != str {
		t.Error(`Expected: lParen`+str+`, got `, c, ": ", l)
	}
}

func TestIdentifyRParen(t *testing.T) {
	str := `)`
	c, l := stateMachineLooper(str)
	if c != rParen && l != str {
		t.Error(`Expected: rParen`+str+`, got `, c, ": ", l)
	}
}

func TestIdentifyComma(t *testing.T) {
	str := `,`
	c, l := stateMachineLooper(str)
	if c != comma && l != str {
		t.Error(`Expected: comma`+str+`, got `, c, ": ", l)
	}
}

func TestIdentifySemicolon(t *testing.T) {
	str := `;`
	c, l := stateMachineLooper(str)
	if c != semicolon && l != str {
		t.Error(`Expected: semicolon`+str+`, got `, c, ": ", l)
	}
}

func TestIdentifyEqual(t *testing.T) {
	str := `=`
	c, l := stateMachineLooper(str)
	if c != equal && l != str {
		t.Error(`Expected: equal`+str+`, got `, c, ": ", l)
	}
}

func TestIdentifyIdent(t *testing.T) {
	str := `sad1 `
	c, l := stateMachineLooper(str)
	if c != ident && l != str {
		t.Error(`Expected: ident`+str+`, got `, c, ": ", l)
	}
}

func TestIdentifyNumber(t *testing.T) {
	str := `123 `
	c, l := stateMachineLooper(str)
	if c != number && l != str {
		t.Error(`Expected: number`+str+`, got `, c, ": ", l)
	}
}

// func TestIdentifyExp(t *testing.T) {
// 	str := `**`
// 	c, l := stateMachineLooper(str)
// 	if c != exp && l != str {
// 		t.Error(`Expected exp: **, got `, c, ": ", l)
// 	}
// }
