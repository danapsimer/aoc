package location

import (
	"context"
	"github.com/danapsimer/aoc/2021/utils"
	"io"
	"math"
	"os"
)

const (
	defaultConcurrency = 5
	defaultWindowSize  = 3
)

func concurrency(ctx context.Context) int {
	concurrency, ok := ctx.Value("concurrency").(int)
	if ok {
		return concurrency
	}
	return defaultConcurrency
}

func windowSize(ctx context.Context) int {
	windowSize, ok := ctx.Value("windowSize").(int)
	if ok {
		return windowSize
	}
	return defaultWindowSize
}

func CountIncreasesFromPreviousInReader(ctx context.Context, r io.Reader) (int, error) {
	return CountIncreasesFromPreviousInChannel(ctx, utils.ReadIntegersFromReader(ctx, r)), nil
}

func CountIncreasesFromPreviousInChannel(ctx context.Context, integers <-chan int) int {
	increaseCount := 0
	lastV := math.MaxInt
	for v := range integers {
		if v > lastV {
			increaseCount += 1
		}
		lastV = v
	}
	return increaseCount
}

func CountIncreasesFromPreviousInFile(ctx context.Context, filename string) (int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	return CountIncreasesFromPreviousInReader(ctx, f)
}

func GroupByWindow(ctx context.Context, integers <-chan int) <-chan int {
	runningTotals := make(chan int)
	windowSize := windowSize(ctx)
	windowSums := make([]int, windowSize, windowSize)
	go func() {
		defer close(runningTotals)
		p := 0
		for v := range integers {
			for pp, _ := range windowSums {
				if p >= windowSize && p%windowSize == pp {
					runningTotals <- windowSums[pp]
					windowSums[pp] = 0
				}
				if p >= pp {
					windowSums[pp] += v
				}
			}
			p += 1
		}
		if p >= windowSize {
			runningTotals <- windowSums[p%windowSize]
		}
	}()
	return runningTotals
}

func CountIncreasesFromPreviousInReaderPart2(ctx context.Context, r io.Reader) int {
	return CountIncreasesFromPreviousInChannel(ctx, GroupByWindow(ctx, utils.ReadIntegersFromReader(ctx, r)))
}

func CountIncreasesFromPreviousInFilePart2(ctx context.Context, filename string) (int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	return CountIncreasesFromPreviousInReaderPart2(ctx, f), nil
}
