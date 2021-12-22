package decoder

import (
	"fmt"
	"strings"
)

var digits = []string{
	"abcefg",  //0
	"cf",      //1
	"acdeg",   //2
	"acdfg",   //3
	"bcdf",    //4
	"abdfg",   //5
	"abdefg",  //6
	"acf",     //7
	"abcdefg", //8
	"abcdfg",  //9
}
var uniqueDigits = []int{1, 4, 7, 8}
var segmentFrequencies = calcSegmentFrequencies(digits)
var digitScores = calcDigitScores(digits, segmentFrequencies)

func calcSegmentFrequencies(inputs []string) []int {
	frequencies := make([]int, 7)
	for _, input := range inputs {
		for _, c := range input {
			frequencies[int(c-'a')] += 1
		}
	}
	return frequencies
}

func calcDigitScore(digit string, segmentFrequencies []int) int {
	score := 0
	for _, c := range digit {
		score += segmentFrequencies[int(c-'a')]
	}
	return score
}

func calcDigitScores(digits []string, segmentFrequencies []int) []int {
	scores := make([]int, len(digits))
	for i, digit := range digits {
		scores[i] = calcDigitScore(digit, segmentFrequencies)
	}
	return scores
}

func decodeLine(dl *decoderLine) int {
	inFrequencies := calcSegmentFrequencies(dl.in)
	outScores := calcDigitScores(dl.out, inFrequencies)
	value := 0
	for _, score := range outScores {
		value *= 10
		var digit int
		for dx, digitScore := range digitScores {
			if digitScore == score {
				digit = dx
				break
			}
		}
		value += digit
	}
	return value
}

type decoderLine struct {
	in  []string
	out []string
}

func LoadDecoderLines(in <-chan string) <-chan *decoderLine {
	decoderLineCh := make(chan *decoderLine)
	go func() {
		defer close(decoderLineCh)
		for decoderLineStr := range in {
			decoderLine, err := ParseDecoderLine(decoderLineStr)
			if err == nil {
				decoderLineCh <- decoderLine
			}
		}
	}()
	return decoderLineCh
}

func ParseDecoderLine(decoderLineStr string) (*decoderLine, error) {
	parts := strings.Split(decoderLineStr, "|")
	for i, part := range parts {
		parts[i] = strings.TrimSpace(part)
	}
	if len(parts) != 2 {
		return nil, fmt.Errorf("there should be 2 parts seperated by '|': %s", decoderLineStr)
	}
	return &decoderLine{
		in:  strings.Split(parts[0], " "),
		out: strings.Split(parts[1], " "),
	}, nil
}

func CountUniqueOutputs(in <-chan *decoderLine) int {
	count := 0
	for dl := range in {
		for _, out := range dl.out {
			for _, ud := range uniqueDigits {
				if len(out) == len(digits[ud]) {
					count += 1
				}
			}
		}
	}
	return count
}

func SumOutputValues(in <-chan *decoderLine) int {
	sum := 0
	for dl := range in {
		sum += decodeLine(dl)
	}
	return sum
}
