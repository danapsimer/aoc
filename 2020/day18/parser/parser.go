package parser

import (
	"aoc/2020/day18/expr"
	"aoc/2020/day18/lexer"
	"errors"
	"fmt"
)

// expr := addExpr { * addExpr }
// addExpr := subExpr { + subExpr }
// subExpr := "(" expr ")" | number
// number := [0-9]+

type Parser struct {
	tokens    []*lexer.Token
	nextToken int
}

func NewParser(tokens []*lexer.Token) *Parser {
	return &Parser{tokens: tokens}
}

func (p *Parser) ParseSubExpr() (*expr.Expr, error) {
	// ( ... ) | number
	if len(p.tokens) == 0 {
		return nil, errors.New("unexpected end of expression")
	}
	token := p.tokens[p.nextToken]
	p.nextToken += 1
	if token.Type == lexer.OpenParenthesis {
		expr, err := p.ParseExpr()
		if err != nil {
			return nil, err
		}
		if p.tokens[p.nextToken].Type != lexer.CloseParenthesis {
			return nil, fmt.Errorf("expected ')' but got %+v", p.tokens[p.nextToken])
		}
		p.nextToken += 1
		return expr, nil
	} else if token.Type == lexer.Value {
		return &expr.Expr{Operation: expr.Const, Value: token.Value}, nil
	}
	return nil, fmt.Errorf("invalid subexpression: %+v", token)
}

func (p *Parser) ParseAddExpr() (*expr.Expr, error) {
	left, err := p.ParseSubExpr()
	if err != nil {
		return nil, err
	}
	for p.nextToken < len(p.tokens) && p.tokens[p.nextToken].Type == lexer.Addition {
		p.nextToken += 1
		right, err := p.ParseSubExpr()
		if err != nil {
			return nil, err
		}
		left = &expr.Expr{Operation: expr.Add, Arguments: []*expr.Expr{left, right}}
	}
	return left, nil
}

func (p *Parser) ParseExpr() (*expr.Expr, error) {
	left, err := p.ParseAddExpr()
	if err != nil {
		return nil, err
	}
	for p.nextToken < len(p.tokens) && p.tokens[p.nextToken].Type == lexer.Multiplication {
		p.nextToken += 1
		right, err := p.ParseAddExpr()
		if err != nil {
			return nil, err
		}
		left = &expr.Expr{Operation: expr.Mul, Arguments: []*expr.Expr{left, right}}
	}
	return left, nil
}

func (p *Parser) Parse() (*expr.Expr, error) {
	expr, err := p.ParseExpr()
	if err != nil {
		return nil, err
	}
	if p.nextToken < len(p.tokens) {
		return nil, fmt.Errorf("there are tokens left over")
	}
	return expr, nil
}

//func(p *Parser) Parse(tokens []*Token) (*Expr, error) {
//	exprStack := NewExpressionStack()
//	var pushValueExprFn func(*Expr)
//	pushValueExprFn = func(expr *Expr) {
//		last := exprStack.Peek()
//		if last != nil && last.Operation == Add {
//			var op, arg1 *Expr
//			exprStack, op = exprStack.Pop()
//			exprStack, arg1 = exprStack.Pop()
//			op.Arguments = []*Expr{arg1, expr}
//			pushValueExprFn(op)
//		} else {
//			exprStack = exprStack.Push(expr)
//		}
//	}
//	for p := 0; p < len(tokens); p++ {
//		token := tokens[p]
//		switch token.Type {
//		case Value:
//			pushValueExprFn(&Expr{Operation: Const, Value: token.Value})
//		case Multiplication:
//			exprStack = exprStack.Push(&Expr{Operation: Mul})
//		case Addition:
//			exprStack = exprStack.Push(&Expr{Operation: Add})
//		case OpenParenthesis:
//			exprStack = exprStack.Push(&Expr{Operation: SubExpr})
//		case CloseParenthesis:
//			var subExpr *Expr
//			exprStack, subExpr = exprStack.Pop()
//			openSubExpr := exprStack.Peek()
//			if openSubExpr == nil || openSubExpr.Operation != SubExpr {
//				return nil, fmt.Errorf("unmatch close parenthesis at %d", p)
//			}
//			exprStack, _ = exprStack.Pop() // throw away subexpr
//			pushValueExprFn(subExpr)
//		default:
//			return nil, fmt.Errorf("unsupported token: %+v", token)
//		}
//	}
//	exprStack, expr := exprStack.Pop()
//	if len(exprStack) > 0 {
//		return nil, fmt.Errorf("parsing ended with expressions still on the stack: %+v", exprStack)
//	}
//	return expr, nil
//}
