package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
)

func countOrbits(k string, orbits map[string]string) int {
	v, ok := orbits[k]
	if ok {
		return 1 + countOrbits(v, orbits)
	} else {
		return 0
	}
}

var regex = regexp.MustCompile("^([A-Z0-9]{1,3})\\)([A-Z0-9]{1,3})$")

func ReadOrbits(in io.Reader) map[string]string {
	orbits := make(map[string]string)
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		text := scanner.Text()
		matches := regex.FindStringSubmatch(text)
		if matches != nil {
			orbits[matches[2]] = matches[1]
		}
	}
	return orbits
}

func Day06(orbits map[string]string) int {
	count := 0
	for k, _ := range orbits {
		count += countOrbits(k, orbits)
	}
	return count
}

func orbitsOf(start string, orbits map[string]string) []string {
	result := make([]string, 0, 10)
	k := start
	for {
		o, ok := orbits[k]
		if !ok {
			return result
		}
		result = append(result, o)
		k = o
	}
}

func Day06Part2(orbits map[string]string) int {
	orbitsOfSAN := orbitsOf("SAN", orbits)
	orbitsOfYou := orbitsOf("YOU", orbits)
	min := math.MaxInt64
	for san, vsan := range orbitsOfSAN {
		for you, vyou := range orbitsOfYou {
			if vsan == vyou {
				if san+you < min {
					min = san + you
				}
			}
		}
	}
	return min
}

func main() {
	orbits := ReadOrbits(os.Stdin)
	fmt.Printf("part1 = %d\n", Day06(orbits))
	fmt.Printf("part2 = %d\n", Day06Part2(orbits))
}
