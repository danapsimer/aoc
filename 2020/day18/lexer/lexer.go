package lexer

import (
	"regexp"
	"strconv"
)

type TokenType int

const (
	Noop TokenType = iota
	Value
	Multiplication
	Addition
	OpenParenthesis
	CloseParenthesis
)

type Token struct {
	Type  TokenType
	Value int
}

var tokenRegex = regexp.MustCompile("(\\d+|\\*|\\+|\\(|\\))")

func Tokenize(expr string) ([]*Token, error) {
	matches := tokenRegex.FindAllStringSubmatch(expr, -1)
	if matches == nil {
		return nil, nil
	}
	tokens := make([]*Token, 0, 100)
	for _, match := range matches {
		switch {
		case match[0] == "*":
			tokens = append(tokens, &Token{Multiplication, 0})
		case match[0] == "+":
			tokens = append(tokens, &Token{Addition, 0})
		case match[0] == "(":
			tokens = append(tokens, &Token{OpenParenthesis, 0})
		case match[0] == ")":
			tokens = append(tokens, &Token{CloseParenthesis, 0})
		default:
			value, err := strconv.Atoi(match[0])
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, &Token{Value, value})
		}
	}
	return tokens, nil
}
