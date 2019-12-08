package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

var (
	three = big.NewInt(3)
	two   = big.NewInt(2)
	zero  = big.NewInt(0)
)

func calcMass(mass int) int {
	m := mass/3 - 2
	if m > 0 {
		fmt.Printf(" + %d", m)
		return calcMass(m) + mass
	}
	return mass
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var totalMass int
	for scanner.Scan() {
		var mass int
		fmt.Sscanf(scanner.Text(), "%d", &mass)
		fmt.Printf("calcMass(%d) = %d", mass, mass)
		m := calcMass(mass) - mass
		fmt.Printf(" = %d\n", m)
		totalMass += m
	}
	fmt.Printf("%d\n", totalMass)
}
