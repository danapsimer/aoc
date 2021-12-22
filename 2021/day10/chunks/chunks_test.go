package chunks

import (
	"fmt"
	"github.com/danapsimer/aoc/2021/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

var CheckLinesScenarios = []struct {
	Lines         []string
	ExpectedScore int
}{
	{
		Lines: []string{
			"[({(<(())[]>[[{[]{<()<>>",
			"[(()[<>])]({[<{<<[]>>(",
			"{([(<{}[<>[]}>{[]{[(<()>",
			"(((({<>}<{<{<>}{[]{[]{}",
			"[[<[([]))<([[{}[[()]]]",
			"[{[{({}]{}}([{[{{{}}([]",
			"{<[[]]>}<{[{[{[]{()[[[]",
			"[<(<(<(<{}))><([]([]()",
			"<{([([[(<>()){}]>(<<{{",
			"<{([{{}}[<[[[<>{}]]]>[]]",
		},
		ExpectedScore: 26397,
	},
}

func TestCheckLines(t *testing.T) {
	for d, scenario := range CheckLinesScenarios {
		t.Run(fmt.Sprintf("CheckLines(%d)", d), func(t *testing.T) {
			assert.Equal(t, scenario.ExpectedScore, CheckLines(utils.StringArrayToChannel(scenario.Lines)))
		})
	}
}
