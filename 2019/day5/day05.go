package main

import (
	"aoc/2019/intCode"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
		in := make(chan int)
		prg := intCode.NewIntCodeProgramWithInput(program, in)
		go prg.RunProgram()
		in <- 5
		var o int
		for o = range prg.GetOutput() {
			fmt.Printf("%d,", o)
		}
		fmt.Println()
	}
}
