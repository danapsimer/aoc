package polymers

import (
	"fmt"
	"math"
)

type Polymer string

type applyRulesMemoKey struct {
	step    int
	polymer Polymer
}

func (p Polymer) ApplyRules(rules map[string]rune, steps int) map[rune]int {
	var memo = make(map[applyRulesMemoKey]map[rune]int)
	var lastRune rune
	counts := make(map[rune]int)
	for _, currRune := range p {
		if lastRune != 0 {
			for r, c := range Polymer([]rune{lastRune, currRune}).applyRules(rules, steps, memo) {
				counts[r] += c
			}
		} else {
			counts[currRune] += 1
		}
		lastRune = currRune
	}
	return counts
}

func (p Polymer) applyRules(rules map[string]rune, steps int, memo map[applyRulesMemoKey]map[rune]int) map[rune]int {
	asRunes := []rune(p)
	if steps == 0 {
		return map[rune]int{asRunes[1]: 1}
	}
	memoKey := applyRulesMemoKey{steps, p}
	if counts, ok := memo[memoKey]; ok {
		return counts
	}
	if len(asRunes) != 2 {
		return map[rune]int{}
	}
	counts := make(map[rune]int)
	insertion, ok := rules[string(p)]
	if ok {
		for r, c := range Polymer([]rune{asRunes[0], insertion}).applyRules(rules, steps-1, memo) {
			counts[r] += c
		}
		for r, c := range Polymer([]rune{insertion, asRunes[1]}).applyRules(rules, steps-1, memo) {
			counts[r] += c
		}
	}
	memo[memoKey] = counts
	return counts
}

func (p Polymer) FindMostAndLeastCommonElements() (string, int, string, int) {
	mostCommon := ""
	mostCommonCount := 0
	leastCommon := ""
	leastCommonCount := math.MaxInt
	counts := make(map[Polymer]int)
	for c := 0; c < len(p); c++ {
		element := p[c : c+1]
		counts[element] += 1
	}
	for element, count := range counts {
		if mostCommonCount < count {
			mostCommon = string(element)
			mostCommonCount = count
		}
		if leastCommonCount > count {
			leastCommon = string(element)
			leastCommonCount = count
		}
	}
	return mostCommon, mostCommonCount, leastCommon, leastCommonCount
}

func LoadPolymerAndRules(lines <-chan string) (Polymer, map[string]rune, error) {
	readingPolymer := true
	var polymer Polymer
	rules := make(map[string]rune)
	for line := range lines {
		if line == "" {
			readingPolymer = false
			continue
		}
		if readingPolymer {
			polymer = Polymer(line)
		} else {
			var pair, insertion string
			n, err := fmt.Sscanf(line, "%s -> %s", &pair, &insertion)
			if err != nil {
				return "", nil, err
			}
			if n != 2 {
				return "", nil, fmt.Errorf("unable to parse insertion rule: %s", line)
			}
			// Insertion should only have one element
			for _, r := range insertion {
				rules[pair] = r
			}
		}
	}
	return polymer, rules, nil
}
