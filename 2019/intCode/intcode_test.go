package intCode

import (
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
	"strings"
	"testing"
)

var (
	inputPrograms = [][]int{
		{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50},
		{1, 0, 0, 0, 99},
		{2, 3, 0, 3, 99},
		{2, 4, 4, 5, 99, 0},
		{1, 1, 1, 4, 99, 5, 6, 0, 99},
		{1002, 4, 3, 4, 33},
		{1101, 100, -1, 4, 0},
		{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
		{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
		{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
		{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
		{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
		{3, 3, 1108, -1, 8, 3, 4, 3, 99},
		{3, 3, 1108, -1, 8, 3, 4, 3, 99},
		{3, 3, 1107, -1, 8, 3, 4, 3, 99},
		{3, 3, 1107, -1, 8, 3, 4, 3, 99},
		{3, 3, 1107, -1, 8, 3, 4, 3, 99},
		{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
		{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
		{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1},
		{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1},
		{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
			1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
			999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
	}
	inputs = [][]int{
		{},
		{},
		{},
		{},
		{},
		{},
		{},
		{8},
		{9},
		{8},
		{7},
		{9},
		{8},
		{9},
		{8},
		{7},
		{9},
		{5},
		{0},
		{5},
		{0},
		{7},
		{8},
		{9},
	}
	outputPrograms = [][]int{
		{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50},
		{2, 0, 0, 0, 99},
		{2, 3, 0, 6, 99},
		{2, 4, 4, 5, 99, 9801},
		{30, 1, 1, 4, 2, 5, 6, 0, 99},
		{1002, 4, 3, 4, 99},
		{1101, 100, -1, 4, 99},
		{3, 9, 8, 9, 10, 9, 4, 9, 99, 1, 8},
		{3, 9, 8, 9, 10, 9, 4, 9, 99, 0, 8},
		{3, 9, 7, 9, 10, 9, 4, 9, 99, 0, 8},
		{3, 9, 7, 9, 10, 9, 4, 9, 99, 1, 8},
		{3, 9, 7, 9, 10, 9, 4, 9, 99, 0, 8},
		{3, 3, 1108, 1, 8, 3, 4, 3, 99},
		{3, 3, 1108, 0, 8, 3, 4, 3, 99},
		{3, 3, 1107, 0, 8, 3, 4, 3, 99},
		{3, 3, 1107, 1, 8, 3, 4, 3, 99},
		{3, 3, 1107, 0, 8, 3, 4, 3, 99},
		{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, 5, 1, 1, 9},
		{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, 0, 0, 1, 9},
		{3, 3, 1105, 5, 9, 1101, 0, 0, 12, 4, 12, 99, 1},
		{3, 3, 1105, 0, 9, 1101, 0, 0, 12, 4, 12, 99, 0},
		{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
			1106, 0, 36, 98, 0, 7, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
			999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
		{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
			1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
			999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
		{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
			1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
			999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
	}
	expectedOutputs = [][]int{
		{},
		{},
		{},
		{},
		{},
		{},
		{},
		{1},
		{0},
		{0},
		{1},
		{0},
		{1},
		{0},
		{0},
		{1},
		{0},
		{1},
		{0},
		{1},
		{0},
		{999},
		{1000},
		{1001},
	}
)

func TestIntCodeProgram_RunProgram(t *testing.T) {
	for i, inputProgram := range inputPrograms {
		prg := NewIntCodeProgram(inputProgram)
		go prg.RunProgram()
		outputs := make([]int, 0, 100)
		go func() {
			for _, in := range inputs[i] {
				prg.GetInput() <- in
			}
		}()
		for o := range prg.GetOutput() {
			outputs = append(outputs, o)
		}
		if !assert.ObjectsAreEqualValues(expectedOutputs[i], outputs) {
			t.Errorf("outputs are not equal: expected %v and got %v", expectedOutputs[i], outputs)
		}
		if !assert.ObjectsAreEqualValues(outputPrograms[i], prg.GetProgram()) {
			t.Errorf("programs are not equal: expected \n      %v\nand got %v", outputPrograms[i], prg.GetProgram())
		}
	}
}

func TestReadIntCodeProgramWithInput(t *testing.T) {
	for i, inputProgram := range inputPrograms {
		inputProgramStr := ""
		for p, ip := range inputProgram {
			if p%10 == 9 {
				inputProgramStr += "\n\t"
			} else if p > 0 {
				inputProgramStr += ", "
			}
			inputProgramStr += strconv.Itoa(ip)
		}
		log.Printf("inputProgramStr = %s", inputProgramStr)
		prg := ReadIntCodeProgram(strings.NewReader(inputProgramStr))
		go prg.RunProgram()
		outputs := make([]int, 0, 100)
		go func() {
			for _, in := range inputs[i] {
				prg.GetInput() <- in
			}
		}()
		for o := range prg.GetOutput() {
			outputs = append(outputs, o)
		}
		if !assert.ObjectsAreEqualValues(expectedOutputs[i], outputs) {
			t.Errorf("outputs are not equal: expected %v and got %v", expectedOutputs[i], outputs)
		}
		if !assert.ObjectsAreEqualValues(outputPrograms[i], prg.GetProgram()) {
			t.Errorf("programs are not equal: expected \n      %v\nand got %v", outputPrograms[i], prg.GetProgram())
		}
	}
}