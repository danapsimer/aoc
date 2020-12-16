package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	Start, End int
}

func (r Range) In(v int) bool {
	return r.Start <= v && v <= r.End
}

type Field struct {
	Name   string
	Ranges []Range
}

func (f *Field) Valid(v int) bool {
	for _, r := range f.Ranges {
		if r.In(v) {
			return true
		}
	}
	return false
}

type Ticket []int

type Notes struct {
	Fields        map[string]*Field
	MyTicket      Ticket
	NearbyTickets []Ticket
}

var fieldRegexp = regexp.MustCompile("^([a-zA-Z ]+):\\s+(\\d+)-(\\d+)\\s+or\\s+(\\d+)-(\\d+)$")

func ReadNotes(reader io.Reader) *Notes {
	scanner := bufio.NewScanner(reader)
	stage := 0
	notes := Notes{make(map[string]*Field), nil, make([]Ticket, 0, 1000)}
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		switch stage {
		case 0: // reading Fields
			if line == "your ticket:" {
				stage = 1
			} else {
				match := fieldRegexp.FindStringSubmatch(line)
				r1s, err := strconv.Atoi(match[2])
				if err != nil {
					panic(fmt.Errorf("field field rule: %s - %s", line, err.Error()))
				}
				r1e, err := strconv.Atoi(match[3])
				if err != nil {
					panic(fmt.Errorf("field field rule: %s - %s", line, err.Error()))
				}
				r2s, err := strconv.Atoi(match[4])
				if err != nil {
					panic(fmt.Errorf("field field rule: %s - %s", line, err.Error()))
				}
				r2e, err := strconv.Atoi(match[5])
				if err != nil {
					panic(fmt.Errorf("field field rule: %s - %s", line, err.Error()))
				}
				notes.Fields[match[1]] = &Field{match[1], []Range{Range{r1s, r1e}, Range{r2s, r2e}}}
			}
		case 1: // Reading My Ticket
			if line == "nearby tickets:" {
				stage = 2
			} else {
				notes.MyTicket = parseTicket(line)
			}
		case 2: // Reading Nearby Tickets
			notes.NearbyTickets = append(notes.NearbyTickets, parseTicket(line))
		}
	}
	return &notes
}

func parseTicket(line string) Ticket {
	values := strings.Split(line, ",")
	ticket := make(Ticket, len(values))
	var err error
	for idx, value := range values {
		ticket[idx], err = strconv.Atoi(value)
		if err != nil {
			panic(fmt.Errorf("my ticket: %s - %s", line, err.Error()))
		}
	}
	return ticket
}

func main() {
	notes := ReadNotes(os.Stdin)
	sumOfInvalidValues, validTickets := SumInvalidValues(notes)
	log.Printf("Sum of Invalid Values: %d", sumOfInvalidValues)
	var fieldOrder []string = CalculateFieldOrder(notes, validTickets)
	answer := 1
	for idx, fieldName := range fieldOrder {
		if strings.HasPrefix(fieldName, "departure") {
			a := answer
			answer *= notes.MyTicket[idx]
			log.Printf("multiplying %s by %d = %d", fieldName, a, answer)
		}
	}
	log.Printf("Multiplication of departure fields: %d", answer)
}

func contains(a []string, s string) bool {
	for _, as := range a {
		if as == s {
			return true
		}
	}
	return false
}

func CalculateFieldOrder(notes *Notes, tickets []Ticket) []string {
	fieldCount := len(notes.Fields)
	fieldOrder := make([]string, fieldCount)
	candidates := make(map[int][]string)
	for fn := 0; fn < fieldCount; fn++ {
		for _, field := range notes.Fields {
			tCount := 0
			for _, ticket := range tickets {
				if field.Valid(ticket[fn]) {
					tCount += 1
				}
			}
			if tCount == len(tickets) {
				if _, present := candidates[fn]; !present {
					candidates[fn] = make([]string,0,10)
				}
				candidates[fn] = append(candidates[fn], field.Name)
			}
		}
	}
	order := make([]int,fieldCount)
	for i := 0; i < fieldCount; i++ {
		order[i] = i
	}
	sort.Slice(order, func(i, j int) bool {
		return len(candidates[order[i]]) < len(candidates[order[j]])
	})
	for len(candidates) > 0 {
		for _, fn := range order {
			possibles := candidates[fn]
			for p := 0; p < len(possibles) && len(possibles) > 0; {
				if contains(fieldOrder, possibles[p]) {
					possibles = append(possibles[0:p], possibles[p+1:]...)
				} else {
					p += 1
				}
			}
			if len(possibles) == 0 {
				panic(fmt.Errorf("%d has no possible fields", fn))
			} else if len(possibles) == 1 {
				fieldOrder[fn] = possibles[0]
				delete(candidates, fn)
			} else {
				candidates[fn] = possibles
			}
		}
	}
	if len(fieldOrder) != fieldCount {
		panic("could not find all field positions")
	}
	return fieldOrder
}

func SumInvalidValues(notes *Notes) (int, []Ticket) {
	sumOfInvalidValues := 0
	validTickets := make([]Ticket, 0, len(notes.NearbyTickets))
	for _, t := range notes.NearbyTickets {
		invalid := false
		for _, v := range t {
			valid := false
			for _, f := range notes.Fields {
				if f.Valid(v) {
					valid = true
					break
				}
			}
			if !valid {
				sumOfInvalidValues += v
				invalid = true
			}
		}
		if !invalid {
			validTickets = append(validTickets, t)
		}
	}
	return sumOfInvalidValues, validTickets
}
