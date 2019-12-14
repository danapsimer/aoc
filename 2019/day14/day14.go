package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
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

type requirement struct {
	name string
	qty  int
}

func (reactions Reactions) React(qty int) int {
	requiredOre := 0
	requirementQueue := make([]requirement, 0, 1000)
	requirementQueue = append(requirementQueue, requirement{"FUEL", qty})
	leftOvers := make(map[string]int)
	for len(requirementQueue) > 0 {
		req := requirementQueue[0]
		requirementQueue = requirementQueue[1:]
		if req.name == "ORE" {
			requiredOre += req.qty
			continue
		}

		leftOversAvailable, ok := leftOvers[req.name]
		if ok {
			delete(leftOvers, req.name)
		}
		qtyNeeded := req.qty - leftOversAvailable
		if qtyNeeded < 0 {
			leftOvers[req.name] = -qtyNeeded
		} else if qtyNeeded > 0 {
			reaction, ok := reactions[req.name]
			if !ok {
				panic(fmt.Errorf("couldn't find requirement: %s", req.name))
			}
			reactionsNeeded := qtyNeeded / reaction.outputQty
			if qtyNeeded%reaction.outputQty > 0 {
				reactionsNeeded += 1
			}
			for inName, inQty := range reactions[req.name].inputs {
				requirementQueue = append(requirementQueue, requirement{inName, inQty * reactionsNeeded})
			}
			if qtyNeeded < reaction.outputQty*reactionsNeeded {
				leftOvers[req.name] = reaction.outputQty*reactionsNeeded - qtyNeeded
			}
		}
	}
	return requiredOre
}
const oreAvailable = 1000000000000
func Day14Part2(reactions Reactions) int {
	lowerLimit :=  oreAvailable / reactions.React(1)
	dir := 1
	step := 1000000
	qty := lowerLimit
	for {
		ore := reactions.React(qty)
		if dir == 1 {
			if ore > oreAvailable {
				dir = -1
				if step > 1 {
					step = step / 10
				}
			} else {
				qty += step
			}
		} else {
			if ore < oreAvailable {
				dir = 1
				if step > 1 {
					step = step / 10
				} else {
					return qty
				}
			} else {
				qty -= step
			}
		}
	}
}

func Day14Part1(reactions Reactions) int {
	return reactions.React(1)
}

func main() {
	reactions := ReadReactions(os.Stdin)
	log.Printf("part1 = %d", Day14Part1(reactions))
	log.Printf("part2 = %d", Day14Part2(reactions))
}
