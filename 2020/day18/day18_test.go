package main

import (
	"aoc/2020/day18/expr"
	"aoc/2020/day18/lexer"
	"aoc/2020/day18/parser"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestExpression struct {
	expr           string
	expectedValue  int
	expectedTokens []*lexer.Token
	expectedExpr   *expr.Expr
}

var testExpressions []*TestExpression = []*TestExpression{
	{
		"1 + 2 * 3 + 4 * 5 + 6", 231,
		[]*lexer.Token{
			{lexer.Value, 1},
			{lexer.Addition, 0},
			{lexer.Value, 2},
			{lexer.Multiplication, 0},
			{lexer.Value, 3},
			{lexer.Addition, 0},
			{lexer.Value, 4},
			{lexer.Multiplication, 0},
			{lexer.Value, 5},
			{lexer.Addition, 0},
			{lexer.Value, 6},
		},
		&expr.Expr{
			Operation: expr.Add,
			Arguments: []*expr.Expr{
				{Operation: expr.Mul, Arguments: []*expr.Expr{
					{Operation: expr.Add, Arguments: []*expr.Expr{
						{Operation: expr.Mul, Arguments: []*expr.Expr{
							{Operation: expr.Add, Arguments: []*expr.Expr{
								{Operation: expr.Const, Value: 1},
								{Operation: expr.Const, Value: 2},
							}},
							{Operation: expr.Const, Value: 3},
						}},
						{Operation: expr.Const, Value: 4},
					}},
					{Operation: expr.Const, Value: 5},
				}},
				{Operation: expr.Const, Value: 6},
			},
		},
	},
	{"2 * 3 + (4 * 5)", 46,
		[]*lexer.Token{
			{lexer.Value, 2},
			{lexer.Multiplication, 0},
			{lexer.Value, 3},
			{lexer.Addition, 0},
			{lexer.OpenParenthesis, 0},
			{lexer.Value, 4},
			{lexer.Multiplication, 0},
			{lexer.Value, 5},
			{lexer.CloseParenthesis, 0},
		},
		&expr.Expr{
			Operation: expr.Add,
			Arguments: []*expr.Expr{
				{Operation: expr.Mul, Arguments: []*expr.Expr{
					{Operation: expr.Const, Value: 2},
					{Operation: expr.Const, Value: 3},
				}},
				{Operation: expr.Mul, Arguments: []*expr.Expr{
					{Operation: expr.Const, Value: 4},
					{Operation: expr.Const, Value: 5},
				}},
			},
		},
	},
	{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 1445, []*lexer.Token{
		{lexer.Value, 5},
		{lexer.Addition, 0},
		{lexer.OpenParenthesis, 0},
		{lexer.Value, 8},
		{lexer.Multiplication, 0},
		{lexer.Value, 3},
		{lexer.Addition, 0},
		{lexer.Value, 9},
		{lexer.Addition, 0},
		{lexer.Value, 3},
		{lexer.Multiplication, 0},
		{lexer.Value, 4},
		{lexer.Multiplication, 0},
		{lexer.Value, 3},
		{lexer.CloseParenthesis, 0},
	}, nil},
	{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 669060, []*lexer.Token{
		{lexer.Value, 5},
		{lexer.Multiplication, 0},
		{lexer.Value, 9},
		{lexer.Multiplication, 0},
		{lexer.OpenParenthesis, 0},
		{lexer.Value, 7},
		{lexer.Multiplication, 0},
		{lexer.Value, 3},
		{lexer.Multiplication, 0},
		{lexer.Value, 3},
		{lexer.Addition, 0},
		{lexer.Value, 9},
		{lexer.Multiplication, 0},
		{lexer.Value, 3},
		{lexer.Addition, 0},
		{lexer.OpenParenthesis, 0},
		{lexer.Value, 8},
		{lexer.Addition, 0},
		{lexer.Value, 6},
		{lexer.Multiplication, 0},
		{lexer.Value, 4},
		{lexer.CloseParenthesis, 0},
		{lexer.CloseParenthesis, 0},
	}, nil},
	{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 23340, []*lexer.Token{
		{lexer.OpenParenthesis, 0},
		{lexer.OpenParenthesis, 0},
		{lexer.Value, 2},
		{lexer.Addition, 0},
		{lexer.Value, 4},
		{lexer.Multiplication, 0},
		{lexer.Value, 9},
		{lexer.CloseParenthesis, 0},
		{lexer.Multiplication, 0},
		{lexer.OpenParenthesis, 0},
		{lexer.Value, 6},
		{lexer.Addition, 0},
		{lexer.Value, 9},
		{lexer.Multiplication, 0},
		{lexer.Value, 8},
		{lexer.Addition, 0},
		{lexer.Value, 6},
		{lexer.CloseParenthesis, 0},
		{lexer.Addition, 0},
		{lexer.Value, 6},
		{lexer.CloseParenthesis, 0},
		{lexer.Addition, 0},
		{lexer.Value, 2},
		{lexer.Addition, 0},
		{lexer.Value, 4},
		{lexer.Multiplication, 0},
		{lexer.Value, 2},
	}, nil},
}

func TestTokenize(t *testing.T) {
	for _, testExpression := range testExpressions {
		t.Run(testExpression.expr, func(t *testing.T) {
			tokens, err := lexer.Tokenize(testExpression.expr)
			if assert.NoError(t, err) {
				if assert.NotNil(t, tokens) {
					assert.EqualValues(t, testExpression.expectedTokens, tokens)
				}
			}
		})
	}
}

func TestParse(t *testing.T) {
	for _, testExpression := range testExpressions {
		t.Run(testExpression.expr, func(t *testing.T) {
			expr, err := parser.NewParser(testExpression.expectedTokens).Parse()
			if assert.NoError(t, err) {
				if assert.NotNil(t, expr) {
					assert.EqualValues(t, testExpression.expectedExpr, expr)
				}
			}
		})
	}
}

func TestExpr_Evaluate(t *testing.T) {
	for _, testExpression := range testExpressions {
		t.Run(testExpression.expr, func(t *testing.T) {
			expr, err := parser.NewParser(testExpression.expectedTokens).Parse()
			if assert.NoError(t, err) {
				if assert.NotNil(t, expr) {
					value, err := expr.Evaluate()
					if assert.NoError(t, err) {
						assert.EqualValues(t, testExpression.expectedValue, value)
					}
				}
			}
		})
	}
}
