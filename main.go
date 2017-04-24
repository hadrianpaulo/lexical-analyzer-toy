package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	inputFile := os.Args[1]
	outputFile := os.Args[2]
	b, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)

	var tokenQueue []state
	var lexemeQueue []string

	token, lexeme, d := tokenizer(str)

	tokenQueue = append(tokenQueue, convertNumberState(token))
	lexemeQueue = append(lexemeQueue, lexeme)

	for d != "" {
		token, lexeme, d = tokenizer(d)
		if token != WHITESPACE {
			tokenQueue = append(tokenQueue, convertNumberState(token))
			lexemeQueue = append(lexemeQueue, lexeme)
		}
	}
	writeToFile(outputFile, tokenQueue, lexemeQueue)
}
