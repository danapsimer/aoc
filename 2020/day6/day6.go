package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

type answers [26]bool
type group []answers

func (a answers) countYes() int {
	c := 0
	for _, answer := range a {
		if answer {
			c += 1
		}
	}
	return c
}

func (g group) mergeAnswers() answers {
	var merged answers
	for _, a := range g {
		for idx, answer := range a {
			merged[idx] = merged[idx] || answer
		}
	}
	return merged
}

func (g group) mergeAnswersPart2() answers {
	var merged answers
	for x := range merged {
		merged[x] = true
	}
	for _, a := range g {
		for idx, answer := range a {
			merged[idx] = merged[idx] && answer
		}
	}
	return merged
}

func main() {
	groups := readGroups(os.Stdin)
	count := 0
	countPart2 := 0
	for _, g := range groups {
		a := g.mergeAnswers()
		for _, answer := range a {
			if answer {
				count += 1
			}
		}
		a = g.mergeAnswersPart2()
		for _, answer := range a {
			if answer {
				countPart2 += 1
			}
		}
	}
	log.Printf("sum of any yes counts = %d", count)
	log.Printf("sum of all yes counts = %d", countPart2)
}

func readGroups(r io.Reader) []group {
	scanner := bufio.NewScanner(r)
	groups := make([]group, 0, 1000)
	for {
		group := readGroup(scanner)
		if group == nil {
			return groups
		}
		groups = append(groups, group)
	}
}

func readGroup(scanner *bufio.Scanner) group {
	g := make([]answers, 0, 20)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			return g
		}
		a := parseAnswer(text)
		g = append(g, a)
	}
	if len(g) == 0 {
		return nil
	}
	return g
}

func parseAnswer(text string) answers {
	var a answers
	for _, c := range text {
		a[c-'a'] = true
	}
	return a
}
