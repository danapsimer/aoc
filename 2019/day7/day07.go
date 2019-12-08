package main

import (
	"aoc/2019/intCode"
	"log"
	"os"
)

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func swap(s string, p1, p2 int) string {
	if p1 == p2 {
		panic("swap called with equal values")
	}
	s1 := s[p1 : p1+1]
	s2 := s[p2 : p2+1]
	if p2 == len(s) {
		return s[:p1] + s2 + s[p1+1:p2] + s1
	}
	return s[:p1] + s2 + s[p1+1:p2] + s1 + s[p2+1:]
}

func permutations(s string, l, r int) []string {
	if l == r {
		return []string{s}
	} else {
		result := make([]string, 0, factorial(r-l+1))
		for i := l; i <= r; i++ {
			if l != i {
				s = swap(s, l, i)
			}
			result = append(result, permutations(s, l+1, r)...)
			if l != i {
				s = swap(s, l, i)
			}
		}
		return result
	}
}

func RunPermutation(prg *intCode.IntCodeProgram, p string) int {
	amplifiers := []*intCode.IntCodeProgram{
		prg.Copy(),
		prg.Copy(),
		prg.Copy(),
		prg.Copy(),
		prg.Copy(),
	}

	for n, a := range amplifiers {
		go func(n int, a *intCode.IntCodeProgram) {
			a.RunProgram()
		}(n, a)
		if n < len(amplifiers)-1 {
			go func(n int, a *intCode.IntCodeProgram) {
				output := <-a.GetOutput()
				amplifiers[n+1].GetInput() <- output
			}(n, a)
		}
	}

	for n, c := range p {
		phase := int(c - '0')
		amplifiers[n].GetInput() <- phase
	}

	amplifiers[0].GetInput() <- 0

	output := <-amplifiers[len(amplifiers)-1].GetOutput()

	return output
}

func FindHighestOutput(prg *intCode.IntCodeProgram) int {
	highestOutput := 0
	highestPermutation := ""
	for _, p := range permutations("01234", 0, 4) {
		output := RunPermutation(prg, p)
		if output > highestOutput {
			highestOutput = output
			highestPermutation = p
		}
	}
	log.Printf("highest permutation = %s\n", highestPermutation)
	return highestOutput
}

func RunPermutationPart2(prg *intCode.IntCodeProgram, p string) int {
	amplifiers := []*intCode.IntCodeProgram{
		prg.Copy(),
		prg.Copy(),
		prg.Copy(),
		prg.Copy(),
		prg.Copy(),
	}

	seriesOutput := make(chan int)
	for n, a := range amplifiers {
		go func(n int, a *intCode.IntCodeProgram) {
			a.RunProgram()
		}(n, a)
		nn := n + 1
		if nn >= len(amplifiers) {
			nn = 0
		}
		aa := amplifiers[nn]
		go func(n, nn int, a, aa *intCode.IntCodeProgram) {
			for output := range a.GetOutput() {
				if nn != 0 || a.IsRunning() && aa.IsRunning() {
					amplifiers[nn].GetInput() <- output
				} else if nn == 0 {
					seriesOutput <- output
				}
			}
		}(n, nn, a, aa)
	}

	for n, c := range p {
		phase := int(c - '0')
		amplifiers[n].GetInput() <- phase
	}

	amplifiers[0].GetInput() <- 0

	output := <-seriesOutput

	return output
}

func FindHighestOutputPart2(prg *intCode.IntCodeProgram) int {
	highestOutput := 0
	highestPermutation := ""
	for _, p := range permutations("56789", 0, 4) {
		output := RunPermutationPart2(prg, p)
		if output > highestOutput {
			highestOutput = output
			highestPermutation = p
		}
	}
	log.Printf("highest permutation = %s\n", highestPermutation)
	return highestOutput
}
func main() {
	prg := intCode.ReadIntCodeProgram(os.Stdin)
	log.Printf("part01 = %d\n", FindHighestOutput(prg))
	log.Printf("part02 = %d\n", FindHighestOutputPart2(prg))
}
