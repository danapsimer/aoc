package intCode

import "fmt"

type IntCodeProgram struct {
	program []int
	in      <-chan int
	out     chan int
}

type parameter struct {
	mode  int
	value int
}

func NewIntCodeProgram(program []int) *IntCodeProgram {
	return NewIntCodeProgramWithInput(program, nil)
}

func NewIntCodeProgramWithInput(program []int, in <-chan int) *IntCodeProgram {
	cpy := make([]int, len(program))
	copy(cpy, program)
	return &IntCodeProgram{cpy, in, make(chan int)}
}

func (prg *IntCodeProgram) GetOutput() <-chan int {
	return prg.out
}

func (prg *IntCodeProgram) GetProgram() []int {
	return prg.program
}

func (prg *IntCodeProgram) get(p *parameter) int {
	switch p.mode {
	case 0:
		return prg.program[p.value]
	case 1:
		return p.value
	default:
		panic(fmt.Sprintf("Unknown parameter mode: %d", p.mode))
	}
}

func (prg *IntCodeProgram) put(p *parameter, value int) {
	switch p.mode {
	case 0:
		prg.program[p.value] = value
	case 1:
		panic(fmt.Sprintf("Cannot write to an immediate parameter: %d", p.mode))
	default:
		panic(fmt.Sprintf("Unknown parameter mode: %d", p.mode))
	}
}

func (icp *IntCodeProgram) extractParameters(pos, n int) []*parameter {
	args := icp.program[pos+1 : pos+n+1]
	modes := icp.program[pos] / 100
	parameters := make([]*parameter, len(args))
	for i, arg := range args {
		mode := modes % 10
		modes = modes / 10
		parameters[i] = &parameter{mode, arg}
	}
	return parameters
}

func (icp *IntCodeProgram) RunProgram() {
	for pos := 0; pos < len(icp.program); {
		switch (icp.program[pos] % 100) {
		case 1:
			// Addition
			params := icp.extractParameters(pos, 3)
			icp.put(params[2], icp.get(params[0])+icp.get(params[1]))
			pos += 4
		case 2:
			// Muliplication
			params := icp.extractParameters(pos, 3)
			icp.put(params[2], icp.get(params[0])*icp.get(params[1]))
			pos += 4
		case 3:
			// Input
			if icp.in == nil {
				panic("No input provided")
			}
			params := icp.extractParameters(pos, 1)
			icp.put(params[0], <-icp.in)
			pos += 2
		case 4:
			// Output
			params := icp.extractParameters(pos, 1)
			icp.out <- icp.get(params[0])
			pos += 2
		case 5:
			params := icp.extractParameters(pos, 2)
			if icp.get(params[0]) != 0 {
				pos = icp.get(params[1])
			} else {
				pos += 3
			}
		case 6:
			params := icp.extractParameters(pos, 2)
			if icp.get(params[0]) == 0 {
				pos = icp.get(params[1])
			} else {
				pos += 3
			}
		case 7:
			params := icp.extractParameters(pos, 3)
			if icp.get(params[0]) < icp.get(params[1]) {
				icp.put(params[2], 1)
			} else {
				icp.put(params[2], 0)
			}
			pos += 4
		case 8:
			params := icp.extractParameters(pos, 3)
			if icp.get(params[0]) == icp.get(params[1]) {
				icp.put(params[2], 1)
			} else {
				icp.put(params[2], 0)
			}
			pos += 4
		case 99:
			close(icp.out)
			return
		default:
			panic(fmt.Sprintf("unrecognized opcode %d at %d", icp.program[pos], pos))
		}
	}
	panic("unexpected end of program.")
}
