package location

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var scenarios = []struct {
	Input    []int
	Expected []int
}{
	{
		Input:    []int{199, 200, 208, 210, 200, 207, 240, 269, 260},
		Expected: []int{607, 618, 618, 617, 647, 716, 769},
	},
	{
		Input:    []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
		Expected: []int{607, 618, 618, 617, 647, 716, 769, 792},
	},
	{
		Input:    []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263, 201},
		Expected: []int{607, 618, 618, 617, 647, 716, 769, 792, 724},
	},
	{
		Input:    []int{199, 200},
		Expected: []int{},
	},
}

func arrayToChannel(arr []int) <-chan int {
	intCh := make(chan int)
	go func() {
		defer close(intCh)
		for _, v := range arr {
			intCh <- v
		}
	}()
	return intCh
}

func channelToArray(ch <-chan int) []int {
	arr := make([]int, 0, 50)
	for v := range ch {
		arr = append(arr, v)
	}
	return arr
}

func TestGroupByWindow(t *testing.T) {
	for _, scenario := range scenarios {
		t.Run(fmt.Sprintf("Input = %v", scenario.Input), func(t *testing.T) {
			actual := channelToArray(GroupByWindow(context.TODO(), arrayToChannel(scenario.Input)))
			assert.ElementsMatch(t, scenario.Expected, actual)
		})
	}
}
