package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type operation int

const (
	Nop operation = iota
	Acc
	Jmp
)

var opNames = []string{"nop", "acc", "jmp"}

func operationFromString(opStr string) operation {
	switch opStr {
	case "nop":
		return Nop
	case "acc":
		return Acc
	case "jmp":
		return Jmp
	default:
		panic(fmt.Errorf("unknown operation: %s", opStr))
	}
}

func (o operation) String() string {
	return opNames[o]
}

type instruction struct {
	Op       operation
	Argument int
	Seen     bool
}

func (i *instruction) String() string {
	return fmt.Sprintf("%s %d", i.Op.String(), i.Argument)
}

func (i *instruction) Run(acc int) (int, int) {
	i.Seen = true
	switch i.Op {
	case Nop:
		return acc, 1
	case Acc:
		return acc + i.Argument, 1
	case Jmp:
		return acc, i.Argument
	default:
		panic(fmt.Errorf("unknown instruction: %s", i.String()))
	}
}

type program []*instruction

func (p program) Clone() program {
	clone := make(program, len(p))
	for idx, i := range p {
		clone[idx] = &instruction{i.Op, i.Argument, false}
	}
	return clone
}

func (p program) Run() (int, bool) {
	accumulator := 0
	for pos := 0; pos < len(p); {
		if p[pos].Seen {
			return accumulator, false
		}
		var delta int
		accumulator, delta = p[pos].Run(accumulator)
		pos += delta
	}
	return accumulator, true
}

func ReadProgram(reader io.Reader) program {
	scanner := bufio.NewScanner(reader)
	prg := make(program, 0, 1000)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			parts := strings.Split(line, " ")
			op := operationFromString(parts[0])
			arg, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(err)
			}
			prg = append(prg, &instruction{op, arg, false})
		}
	}
	return prg
}

func main() {

	prg := ReadProgram(os.Stdin)
	acc, ok := prg.Clone().Run()
	if ok {
		log.Print("Un-fixed program completed normally?")
	}
	log.Printf("Accumulator = %d", acc)
	for idx, i := range prg {
		if i.Op == Acc {
			continue
		}
		clone := prg.Clone()
		if i.Op == Nop {
			clone[idx] = &instruction{Jmp, i.Argument, false}
		} else if i.Op == Jmp {
			clone[idx] = &instruction{Nop, i.Argument, false}
		}
		acc, ok = clone.Run()
		if ok {
			log.Printf("%d: Accumulator = %d", idx, acc)
			break
		}
	}
}
