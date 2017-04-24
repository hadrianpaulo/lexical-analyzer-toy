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
	c, l, d := tokenizer(str)
	w := bufio.NewWriter(f)

	_, err = fmt.Fprintf(w, "%v %v\n", prettyPrint(c), l)
	check(err)
	w.Flush()
	for d != "" {
		c, l, d = tokenizer(d)
		if c != WHITESPACE {
			_, err = fmt.Fprintf(w, "%v %v\n", prettyPrint(c), l)
			check(err)
			w.Flush()
		}
	}
	f.Sync()
}
