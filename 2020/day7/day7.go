package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type rule struct {
	Name       string
	CanContain map[string]int
}

type rules map[string]*rule

var Rules rules
var bagNameRegex = regexp.MustCompile("^([a-z]+\\s+[a-z]+)\\s+bags\\s+contain")
var bagCanContainRegex = regexp.MustCompile("(\\d+)\\s+([a-z]+\\s+[a-z]+)\\s+bags?")

func readContainerRules(reader io.Reader) map[string]*rule {
	scanner := bufio.NewScanner(reader)
	rules := make(map[string]*rule)
	for scanner.Scan() {
		text := scanner.Text()
		nameSubmatch := bagNameRegex.FindStringSubmatch(text)
		r := &rule{nameSubmatch[1], make(map[string]int)}
		submatches := bagCanContainRegex.FindAllStringSubmatch(text, -1)
		for _, submatch := range submatches {
			if len(submatch[1]) != 0 && len(submatch[2]) != 0 {
				var err error
				r.CanContain[submatch[2]], err = strconv.Atoi(submatch[1])
				if err != nil {
					panic(err)
				}
			}
			log.Print("["+strings.Join(submatch, ", ")+"]", r.String())
		}
		log.Printf("%s = %s", text, r.String())
		rules[r.Name] = r
	}
	return rules
}

func (r *rule) CanContainBag(name string) bool {
	if _, ok := r.CanContain[name]; ok {
		return true
	}
	for n, _ := range r.CanContain {
		if Rules[n].CanContainBag(name) {
			return true
		}
	}
	return false
}

func (r *rule) MustContain() int {
	count := 0
	for name, c := range r.CanContain {
		count += (1 + Rules[name].MustContain()) * c
	}
	return count
}

func (r *rule) String() string {
	sw := new(bytes.Buffer)
	encoder := json.NewEncoder(sw)
	encoder.SetIndent("", "  ")
	encoder.Encode(r)
	return sw.String()
}

func (rr rules) String() string {
	sw := new(bytes.Buffer)
	encoder := json.NewEncoder(sw)
	encoder.SetIndent("", "  ")
	encoder.Encode(rr)
	return sw.String()
}

func main() {

	Rules = readContainerRules(os.Stdin)
	//fmt.Printf("rules = %s\n", Rules.String())
	count := 0
	for _, r := range Rules {
		if r.CanContainBag("shiny gold") {
			count += 1
		}
	}
	fmt.Printf("number of bags that can contain 'shiny gold' bags is %d\n", count)

	fmt.Printf("number of bags a 'shiny gold' bag must contain is %d\n", Rules["shiny gold"].MustContain())
}
