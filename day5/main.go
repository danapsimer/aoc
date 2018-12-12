package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strings"
)

func isUpper(s string) bool {
	return 'A' <= s[0] && s[0] <= 'Z'
}

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Printf("ERROR reading input: %s\n", err.Error())
		os.Exit(-1)
	}
	smallestLen := math.MaxInt64
	s := string(bytes)
	fmt.Printf("input length = %d\n", len(s))
	for c := 'a'; c <= 'z'; c += 1 {
		fmt.Printf("evaluating '%s'\n", string(c))
		e := eliminate(s, c)
		fmt.Printf("eliminated length = %d\n", len(e))
		r := react(e)
		fmt.Printf("reacted length = %d\n", len(r))
		if len(r) < smallestLen {
			smallestLen = len(r)
		}
	}
	fmt.Printf("%d\n", smallestLen)
}

func eliminate(s string, c rune) string {
	for i := 0; i < len(s)-1; {
		c1 := s[i : i+1]
		if strings.ToLower(c1) == string(c) {
			s = s[0:i] + s[i+1:]
		} else {
			i += 1
		}
	}
	return s
}

func react(s string) string {
	s = strings.Trim(s, "\n\t ")
	for {
		var reactions int
		s, reactions = eliminateReactions(s)
		if reactions == 0 {
			break
		}
	}
	return s
}

func eliminateReactions(s string) (string, int) {
	reactions := 0
	for i := 0; i < len(s)-1; {
		c1 := s[i : i+1]
		c2 := s[i+1 : i+2]
		if strings.ToLower(c1) == strings.ToLower(c2) && isUpper(c1) != isUpper(c2) {
			s = s[0:i] + s[i+2:]
			reactions += 1
		} else {
			i += 1
		}
	}
	return s, reactions
}
