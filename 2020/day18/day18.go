package main

import (
	"aoc/2020/day18/lexer"
	"aoc/2020/day18/parser"
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)


func SumExpressions(reader io.Reader) int {
	scanner := bufio.NewScanner(reader)
	count := 0
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		tokens, err := lexer.Tokenize(line)
		if err != nil {
			panic(fmt.Errorf("error parsing line %d: %s", count, line))
		}
		expr, err := parser.NewParser(tokens).Parse()
		if err != nil {
			panic(err)
		}
		value, err := expr.Evaluate()
		if err != nil {
			panic(err)
		}
		sum += value
	}
	return sum
}

func main() {
	log.Printf("sum of all expressions = %d", SumExpressions(os.Stdin))
}
