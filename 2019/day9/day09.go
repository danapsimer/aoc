package main

import (
	"aoc/2019/intCode"
	"fmt"
	"os"
)

func main() {
	prg := intCode.ReadIntCodeProgram(os.Stdin)
	fmt.Print("part1 = ")
	day09(1, prg.Copy())
	fmt.Print("part2 = ")
	day09(2, prg.Copy())
}

func day09(mode int, prg *intCode.IntCodeProgram) int {
	go prg.RunProgram()
	prg.GetInput() <- mode
	errCount := 0
	for out := range prg.GetOutput() {
		if errCount > 0 {
			fmt.Print(",")
		}
		fmt.Printf("%d", out)
		errCount += 1
	}
	fmt.Println()
	return errCount - 1
}
