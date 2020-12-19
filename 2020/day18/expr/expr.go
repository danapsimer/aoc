package expr

import "fmt"

type OpType int

const (
	Const OpType = iota
	Mul
	Add
)

type Expr struct {
	Operation OpType
	Arguments []*Expr
	Value     int
}

func (exp *Expr) Evaluate() (int, error) {
	var err error
	switch exp.Operation {
	case Const:
		return exp.Value, nil
	case Mul:
		if len(exp.Arguments) != 2 {
			return 0, fmt.Errorf("invalid number of arguments for Multiply: %+v", exp.Arguments)
		}
		var arg1, arg2 int
		arg1, err = exp.Arguments[0].Evaluate()
		if err != nil {
			return 0, err
		}
		arg2, err = exp.Arguments[1].Evaluate()
		if err != nil {
			return 0, err
		}
		return arg1 * arg2, nil
	case Add:
		if len(exp.Arguments) != 2 {
			return 0, fmt.Errorf("invalid number of arguments for Addition: %+v", exp.Arguments)
		}
		var arg1, arg2 int
		arg1, err = exp.Arguments[0].Evaluate()
		if err != nil {
			return 0, err
		}
		arg2, err = exp.Arguments[1].Evaluate()
		if err != nil {
			return 0, err
		}
		return arg1 + arg2, nil
	default:
		return 0, fmt.Errorf("Invalid expression operation: %d", exp.Operation)
	}
}
