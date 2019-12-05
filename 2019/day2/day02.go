package main

import (
	"aoc/2019/intCode"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RunWithNounAndVerb(noun, verb int, program []int) []int {
	cpy := make([]int, len(program))
	copy(cpy, program)
	cpy[1] = noun
	cpy[2] = verb
	prg := intCode.NewIntCodeProgram(cpy)
	go prg.RunProgram()
	<-prg.GetOutput()
	return prg.GetProgram()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		programStrs := strings.Split(scanner.Text(), ",")
		program := make([]int, len(programStrs))
		var err error
		for i, s := range programStrs {
			program[i], err = strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
		}
		for noun := 0; noun < 100; noun += 1 {
			for verb := 0; verb < 100; verb += 1 {
				out := RunWithNounAndVerb(noun, verb, program)
				if out[0] == 19690720 {
					fmt.Printf("%d\n", noun*100+verb)
					return
				}
			}
		}
	}
}
