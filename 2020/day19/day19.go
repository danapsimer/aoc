package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Rule struct {
	Match   string
	OrRules [][]int
}

func (r *Rule) Matches(idx int, s string, rules Rules) (bool, []int) {
	if idx >= len(s) {
		return false, nil
	}
	if len(r.Match) > 0 {
		return strings.HasPrefix(s[idx:], r.Match), []int{idx + len(r.Match)}
	}
	allMatchIdxs := make([]int,0,10)
	for _, or := range r.OrRules {
		matches := true
		matchIdxs := make([]int,1,10)
		matchIdxs[0] = idx
		for _, ruleIdx := range or {
			newMatchIdxs := make([]int,0,10)
			for _, matchIdx := range matchIdxs {
				var nextMatchIdxs []int
				matches, nextMatchIdxs = rules[ruleIdx].Matches(matchIdx, s, rules)
				if !matches {
					break
				} else {
					newMatchIdxs = append(newMatchIdxs, nextMatchIdxs...)
				}
			}
			matchIdxs = newMatchIdxs
		}
		if matches {
			allMatchIdxs = append(allMatchIdxs, matchIdxs...)
		}
	}
	if len(allMatchIdxs) > 0 {
		return true, allMatchIdxs
	}
	return false, nil
}

func (r *Rule) MakeRegexStr(rules Rules) string {
	if len(r.Match) > 0 {
		return r.Match
	}
	builder := strings.Builder{}
	builder.WriteString("(")
	for i, orRules := range r.OrRules {
		if i > 0 {
			builder.WriteString("|")
		}
		for _, rule := range orRules {
			builder.WriteString(rules[rule].MakeRegexStr(rules))
		}
	}
	builder.WriteString(")")
	return builder.String()
}

func (r *Rule) MakeRegexp(rules Rules) (*regexp.Regexp, error) {
	expr := "^" + r.MakeRegexStr(rules) + "$"
	log.Printf("expr = %s", expr)
	return regexp.Compile(expr)
}

type Rules map[int]*Rule

func ReadRules(reader io.Reader) Rules {
	scanner := bufio.NewScanner(reader)
	rules := make(Rules)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		parts := strings.Split(line, ":")
		if parts == nil || len(parts) != 2 {
			panic(fmt.Errorf("malformed rule line: %s", line))
		}
		id, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		var rule *Rule
		ruleStr := strings.Trim(parts[1], " ")
		if ruleStr[0] == '"' {
			rule = &Rule{Match: ruleStr[1 : len(ruleStr)-1]}
		} else {
			parts = strings.Split(ruleStr, "|")
			ors := make([][]int, 0, 10)
			for _, part := range parts {
				part = strings.Trim(part, " ")
				orParts := strings.Split(part, " ")
				or := make([]int, 0, 10)
				for _, orPart := range orParts {
					idx, err := strconv.Atoi(orPart)
					if err != nil {
						panic(err)
					}
					or = append(or, idx)
				}
				if len(or) == 0 {
					panic("empty rule list")
				}
				ors = append(ors, or)
			}
			if len(ors) == 0 {
				panic("empty or list")
			}
			rule = &Rule{OrRules: ors}
		}
		rules[id] = rule
	}
	return rules
}

func ReadMessages(reader io.Reader) []string {
	messages := make([]string, 0, 1000)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			messages = append(messages, line)
		}
	}
	return messages
}

func CountMatches(rules Rules, messages []string) int {
	rule0 := rules[0]
	count := 0
	//rule0Regex, err := rule0.MakeRegexp(rules)
	//if err != nil {
	//	panic(err)
	//}
	for _, message := range messages {
		matches, matchIdxs := rule0.Matches(0, message, rules)
		if matches {
			fullyConsumed := false
			for _, idx := range matchIdxs {
				if len(message) == idx {
					fullyConsumed = true
				}
			}
			if fullyConsumed {
				count += 1
			}
		}
	}
	return count
}

func main() {
	rules := ReadRules(os.Stdin)
	messages := ReadMessages(os.Stdin)
	count := CountMatches(rules, messages)
	log.Printf("valid message count = %d", count)
	rules[8] = &Rule{OrRules: [][]int{{42}, {42, 8}}}
	rules[11] = &Rule{OrRules: [][]int{{42, 31}, {42, 11, 31}}}

}
