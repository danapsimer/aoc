package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"regexp"
	"strings"
)

var (
	reactionRegex = regexp.MustCompile("^(([0-9]+ [A-Z]+)(,? ))+=> ([0-9]+ [A-Z]+)$")
)

type Reaction struct {
	inputs    map[string]int
	output    string
	outputQty int
}

func (reaction *Reaction) String() string {
	buf := &bytes.Buffer{}
	for n, c := range reaction.inputs {
		if buf.Len() > 0 {
			_, _ = fmt.Fprint(buf, ", ")
		}
		_, _ = fmt.Fprintf(buf, "%d %s", c, n)
	}
	_, _ = fmt.Fprint(buf, " => ")
	_, _ = fmt.Fprintf(buf, "%d %s", reaction.outputQty, reaction.output)
	return buf.String()
}

type Reactions map[string]*Reaction

func (reactions Reactions) String() string {
	buf := &bytes.Buffer{}
	for _, reaction := range reactions {
		_, _ = fmt.Fprintln(buf, reaction.String())
	}
	return buf.String()
}

func ReadReactions(in io.Reader) Reactions {
	scanner := bufio.NewScanner(in)
	reactions := Reactions(make(map[string]*Reaction))
	for scanner.Scan() {
		line := scanner.Text()
		arrowIdx := strings.Index(line, "=>")
		inputsStr := line[:arrowIdx]
		outputStr := line[arrowIdx+2:]
		inputStrs := strings.Split(inputsStr, ",")
		var outName string
		var outNum int
		_, err := fmt.Sscanf(outputStr, "%d %s", &outNum, &outName)
		if err != nil {
			panic(err)
		}
		inputs := make(map[string]int)
		for _, inputStr := range inputStrs {
			var inName string
			var inNum int
			_, err := fmt.Sscanf(strings.Trim(inputStr, " "), "%d %s", &inNum, &inName)
			if err != nil {
				panic(err)
			}
			inputs[inName] = inNum
		}
		reaction := &Reaction{inputs, outName, outNum}
		reactions[outName] = reaction
	}
	log.Println(reactions.String())
	return reactions
}

func (reactions Reactions) React(target string, needed int) int {
	reaction := reactions[target]
	oreNeeded := 0
	reactionCount := needed / reaction.outputQty
	if needed % reaction.outputQty > 0 {
		reactionCount += 1
	}
	for inName, inCount := range reaction.inputs {
		if inName == "ORE" {
			oreNeeded += inCount * reactionCount
		} else {
			oreNeeded += reactions.React(inName, inCount * reactionCount)
		}
	}
	return oreNeeded
}

func Day14Part1(reactions Reactions) int {
	return reactions.React("FUEL", 1)
}

func main() {

}
