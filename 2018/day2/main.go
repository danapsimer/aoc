package main

import (
	"bufio"
	"fmt"
	"os"
)

func hasOneDiff(l1, l2 string) (hasOnlyOneDiff bool, common string) {
	hasOnlyOneDiff = false
	common = ""
	diffCount := 0
	for i := 0; i < len(l1); i++ {
		if l1[i] != l2[i] {
			diffCount += 1
		} else {
			common = common + string(l1[i])
		}
	}
	if diffCount == 1 {
		hasOnlyOneDiff = true
	}
	return
}

func main() {
	lines := make([]string,0,250)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines,scanner.Text())
	}

	for i1, l1 := range lines {
		for i2, l2 := range lines {
			if i1 == i2 {
				continue
			}
			if ok, common := hasOneDiff(l1,l2); ok {
				fmt.Println(common)
				os.Exit(0)
			}
		}
	}
	os.Exit(-1)
}
