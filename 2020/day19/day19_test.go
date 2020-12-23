package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

var testRules = `0: 4 1 5
1: 2 3 | 3 2
2: 4 4 | 5 5
3: 4 5 | 5 4
4: "a"
5: "b"
`

var testInputs = []string{
	"ababbb",
	"bababa",
	"abbbab",
	"aaabbb",
	"aaaabbb",
}

type testResult struct {
	matches bool
	left    int
}

var testResults = []testResult{
	{true, 0},
	{false, 0},
	{true, 0},
	{false, 0},
	{true, 6},
}

func TestRule_Matches(t *testing.T) {
	rules := ReadRules(strings.NewReader(testRules))
	rule0 := rules[0]
	for tidx := 0; tidx < len(testInputs); tidx++ {
		t.Run(testInputs[tidx], func(t *testing.T) {
			matches, left := rule0.Matches(0, testInputs[tidx], rules)
			assert.Equal(t, testResults[tidx].matches, matches)
			if matches {
				assert.Equal(t, testResults[tidx].left, left)
			}
		})
	}
}

func TestCountMatches(t *testing.T) {
	f, err := os.Open("day19.in")
	if assert.NoError(t, err) {
		defer f.Close()
		rules := ReadRules(f)
		messages := ReadMessages(f)
		count := CountMatches(rules, messages)
		assert.Equal(t, 0, count)
	}
}

func TestRegex(t *testing.T) {
	s1 := "^(((b(a(b(a(a(b(ba)|a(ba|ab))|b(b(ba|aa)|a(ab|bb)))|b(a(b(ba|aa)|a(ba))|b((ba|ab)a|(ab|bb)b)))|a((((ba)a|((b|a)(b|a))b)b|((ab|bb)b|(b(b|a)|aa)a)a)b|(a(b(ab|aa)|a(ba|aa))|b((ba|bb)b|(ab|b(b|a))a))a))|b(b(b(b((aa|bb)a|(ba|bb)b)|a((ba)a|(ba|ab)b))|a(a((ba)b|(ba|aa)a)|b(a(b(b|a)|aa)|b(ab|b(b|a)))))|a(b((b(bb)|a(aa|(b|a)b))b|((ba|ab)b|(ba|bb)a)a)|a(a((ba)b|(aa|(b|a)b)a)|b(a(ab)|b(ba|aa))))))|a((((((ba|bb)b|(ab|b(b|a))a)b|((ab)b|(ab|aa)a)a)a|((b(ab|b(b|a))|a(ba|bb))a|((ba|ab)a|(ab|aa)b)b)b)a|((a(b(ab|b(b|a))|a((b|a)(b|a)))|b(a(ab|aa)|b(ba)))a|((a(ba|ab)|b(ab|b(b|a)))b|(b(ab|aa)|a(ba|bb))a)b)b)a|(((a(a(ba)|b(ba))|b(b(ba|ab)|a(ab)))b|(a(b(ba|ab)|a(ab))|b(b(ab|aa)|a(ba|aa)))a)a|((((ab)a|(ba)b)b|((ba|bb)a|(aa|(b|a)b)b)a)a|(b(a(ba|a(b|a))|b(ba))|a((aa|(b|a)b)a|(ab|aa)b))b)b)b)))((b(a(b(a(a(b(ba)|a(ba|ab))|b(b(ba|aa)|a(ab|bb)))|b(a(b(ba|aa)|a(ba))|b((ba|ab)a|(ab|bb)b)))|a((((ba)a|((b|a)(b|a))b)b|((ab|bb)b|(b(b|a)|aa)a)a)b|(a(b(ab|aa)|a(ba|aa))|b((ba|bb)b|(ab|b(b|a))a))a))|b(b(b(b((aa|bb)a|(ba|bb)b)|a((ba)a|(ba|ab)b))|a(a((ba)b|(ba|aa)a)|b(a(b(b|a)|aa)|b(ab|b(b|a)))))|a(b((b(bb)|a(aa|(b|a)b))b|((ba|ab)b|(ba|bb)a)a)|a(a((ba)b|(aa|(b|a)b)a)|b(a(ab)|b(ba|aa))))))|a((((((ba|bb)b|(ab|b(b|a))a)b|((ab)b|(ab|aa)a)a)a|((b(ab|b(b|a))|a(ba|bb))a|((ba|ab)a|(ab|aa)b)b)b)a|((a(b(ab|b(b|a))|a((b|a)(b|a)))|b(a(ab|aa)|b(ba)))a|((a(ba|ab)|b(ab|b(b|a)))b|(b(ab|aa)|a(ba|bb))a)b)b)a|(((a(a(ba)|b(ba))|b(b(ba|ab)|a(ab)))b|(a(b(ba|ab)|a(ab))|b(b(ab|aa)|a(ba|aa)))a)a|((((ab)a|(ba)b)b|((ba|bb)a|(aa|(b|a)b)b)a)a|(b(a(ba|a(b|a))|b(ba))|a((aa|(b|a)b)a|(ab|aa)b))b)b)b))(a((((b((b|a)(ba|aa))|a((aa|(b|a)b)b|(ab)a))a|((a(ab|aa)|b(aa))a|((aa|bb)a|(ba|bb)b)b)b)a|(a(b(a(aa)|b(bb))|a(a(ab|bb)|b(ab)))|b((a(ab)|b(ba|aa))b|((ab|b(b|a))b|(ba|ab)a)a))b)a|(a(a(a(b(ba|ab)|a(ba|a(b|a)))|b(b(ba|a(b|a))|a(aa)))|b(a((aa|(b|a)b)(b|a))|b((ba|bb)(b|a))))|b(b((a(ab|b(b|a))|b(ba|bb))b|((aa|bb)b|(ba|bb)a)a)|a((a(ba|ab)|b(ab|b(b|a)))a|((bb)b|(ba|ab)a)b)))b)|b((a(b(a(a(ab|b(b|a))|b(ba|bb))|b(b(ba|aa)|a((b|a)(b|a))))|a(a(b(ab|aa)|a(ba|ab))|b((ab)a|(ba)b)))|b((b((ba|aa)b|(aa|bb)a)|a((aa|(b|a)b)a|(ab|aa)b))b|(a(b(ab|b(b|a))|a(aa|bb))|b((b|a)(ba|aa)))a))b|(((((aa|bb)a|(b(b|a)|aa)b)a|((bb)b|(ba|ab)a)b)a|((a(aa|bb)|b(ab))b|((aa|bb)(b|a))a)b)b|((b(b(ba)|a(ba|ab))|a((ba|bb)b|(ab|b(b|a))a))a|((a(ba)|b(ab|aa))a|(b(ba|bb)|a(aa))b)b)a)a))))$"
	s2 := "^(((b(a(b(a(a(b(ba)|a(ba|ab))|b(b(ba|aa)|a(ab|bb)))|b(a(b(ba|aa)|a(ba))|b((ba|ab)a|(ab|bb)b)))|a((((ba)a|((b|a)(b|a))b)b|((ab|bb)b|(b(b|a)|aa)a)a)b|(a(b(ab|aa)|a(ba|aa))|b((ba|bb)b|(ab|b(b|a))a))a))|b(b(b(b((aa|bb)a|(ba|bb)b)|a((ba)a|(ba|ab)b))|a(a((ba)b|(ba|aa)a)|b(a(b(b|a)|aa)|b(ab|b(b|a)))))|a(b((b(bb)|a(aa|(b|a)b))b|((ba|ab)b|(ba|bb)a)a)|a(a((ba)b|(aa|(b|a)b)a)|b(a(ab)|b(ba|aa))))))|a((((((ba|bb)b|(ab|b(b|a))a)b|((ab)b|(ab|aa)a)a)a|((b(ab|b(b|a))|a(ba|bb))a|((ba|ab)a|(ab|aa)b)b)b)a|((a(b(ab|b(b|a))|a((b|a)(b|a)))|b(a(ab|aa)|b(ba)))a|((a(ba|ab)|b(ab|b(b|a)))b|(b(ab|aa)|a(ba|bb))a)b)b)a|(((a(a(ba)|b(ba))|b(b(ba|ab)|a(ab)))b|(a(b(ba|ab)|a(ab))|b(b(ab|aa)|a(ba|aa)))a)a|((((ab)a|(ba)b)b|((ba|bb)a|(aa|(b|a)b)b)a)a|(b(a(ba|a(b|a))|b(ba))|a((aa|(b|a)b)a|(ab|aa)b))b)b)b)))((b(a(b(a(a(b(ba)|a(ba|ab))|b(b(ba|aa)|a(ab|bb)))|b(a(b(ba|aa)|a(ba))|b((ba|ab)a|(ab|bb)b)))|a((((ba)a|((b|a)(b|a))b)b|((ab|bb)b|(b(b|a)|aa)a)a)b|(a(b(ab|aa)|a(ba|aa))|b((ba|bb)b|(ab|b(b|a))a))a))|b(b(b(b((aa|bb)a|(ba|bb)b)|a((ba)a|(ba|ab)b))|a(a((ba)b|(ba|aa)a)|b(a(b(b|a)|aa)|b(ab|b(b|a)))))|a(b((b(bb)|a(aa|(b|a)b))b|((ba|ab)b|(ba|bb)a)a)|a(a((ba)b|(aa|(b|a)b)a)|b(a(ab)|b(ba|aa))))))|a((((((ba|bb)b|(ab|b(b|a))a)b|((ab)b|(ab|aa)a)a)a|((b(ab|b(b|a))|a(ba|bb))a|((ba|ab)a|(ab|aa)b)b)b)a|((a(b(ab|b(b|a))|a((b|a)(b|a)))|b(a(ab|aa)|b(ba)))a|((a(ba|ab)|b(ab|b(b|a)))b|(b(ab|aa)|a(ba|bb))a)b)b)a|(((a(a(ba)|b(ba))|b(b(ba|ab)|a(ab)))b|(a(b(ba|ab)|a(ab))|b(b(ab|aa)|a(ba|aa)))a)a|((((ab)a|(ba)b)b|((ba|bb)a|(aa|(b|a)b)b)a)a|(b(a(ba|a(b|a))|b(ba))|a((aa|(b|a)b)a|(ab|aa)b))b)b)b))(a((((b((b|a)(ba|aa))|a((aa|(b|a)b)b|(ab)a))a|((a(ab|aa)|b(aa))a|((aa|bb)a|(ba|bb)b)b)b)a|(a(b(a(aa)|b(bb))|a(a(ab|bb)|b(ab)))|b((a(ab)|b(ba|aa))b|((ab|b(b|a))b|(ba|ab)a)a))b)a|(a(a(a(b(ba|ab)|a(ba|a(b|a)))|b(b(ba|a(b|a))|a(aa)))|b(a((aa|(b|a)b)(b|a))|b((ba|bb)(b|a))))|b(b((a(ab|b(b|a))|b(ba|bb))b|((aa|bb)b|(ba|bb)a)a)|a((a(ba|ab)|b(ab|b(b|a)))a|((bb)b|(ba|ab)a)b)))b)|b((a(b(a(a(ab|b(b|a))|b(ba|bb))|b(b(ba|aa)|a((b|a)(b|a))))|a(a(b(ab|aa)|a(ba|ab))|b((ab)a|(ba)b)))|b((b((ba|aa)b|(aa|bb)a)|a((aa|(b|a)b)a|(ab|aa)b))b|(a(b(ab|b(b|a))|a(aa|bb))|b((b|a)(ba|aa)))a))b|(((((aa|bb)a|(b(b|a)|aa)b)a|((bb)b|(ba|ab)a)b)a|((a(aa|bb)|b(ab))b|((aa|bb)(b|a))a)b)b|((b(b(ba)|a(ba|ab))|a((ba|bb)b|(ab|b(b|a))a))a|((a(ba)|b(ab|aa))a|(b(ba|bb)|a(aa))b)b)a)a))))$"
	assert.Equal(t, s1, s2)
}
