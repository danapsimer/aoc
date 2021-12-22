package chunks

import (
	"container/list"
	"fmt"
	"sort"
)

type chunkDelimiter struct {
	Open, Close     int32
	Score           int
	CompletionScore int
}

var chunkDelimiters = []*chunkDelimiter{
	{'(', ')', 3, 1},
	{'[', ']', 57, 2},
	{'{', '}', 1197, 3},
	{'<', '>', 25137, 4},
}

func isOpenChar(c int32) *chunkDelimiter {
	for _, cd := range chunkDelimiters {
		if cd.Open == c {
			return cd
		}
	}
	return nil
}

func isCloseChar(c int32) *chunkDelimiter {
	for _, cd := range chunkDelimiters {
		if cd.Close == c {
			return cd
		}
	}
	return nil
}

func LineChecker(line string) (int, int) {
	stack := list.New()
	for _, c := range line {
		if stack.Len() > 0 && c == stack.Front().Value.(*chunkDelimiter).Close {
			stack.Remove(stack.Front())
		} else if cd := isOpenChar(c); cd != nil {
			stack.PushFront(cd)
		} else if cd := isCloseChar(c); cd != nil {
			fmt.Printf("corrupted line found: %s\n", line)
			return cd.Score, 0
		} else {
			panic(fmt.Errorf("unexpected character: %c", c))
		}
	}
	completionScore := 0
	for stack.Len() > 0 {
		front := stack.Front()
		stack.Remove(front)
		completionScore *= 5
		completionScore += front.Value.(*chunkDelimiter).CompletionScore
	}
	return 0, completionScore
}

func CheckLines(lines <-chan string) (int, int) {
	completionScores := make([]int, 0, 100)
	score := 0
	for line := range lines {
		s, cs := LineChecker(line)
		score += s
		if cs > 0 {
			completionScores = append(completionScores, cs)
		}
	}
	sort.Ints(completionScores)
	return score, completionScores[len(completionScores)/2]
}
