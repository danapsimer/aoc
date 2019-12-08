package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type claim struct {
	id, x, y, h, w int
}

func parseClaim(s string) (*claim, error) {
	var theClaim claim
	i, err := fmt.Sscanf(s, "#%d @ %d,%d: %dx%d", &theClaim.id, &theClaim.x, &theClaim.y, &theClaim.w, &theClaim.h)
	if err != nil {
		return nil, err
	}
	if i != 5 {
		return nil, fmt.Errorf("malformed line: %s", s)
	}
	return &theClaim, nil
}

type claimCountMap [1016][1012]int

func (ccm *claimCountMap) updateMap(clm *claim) {
	right := clm.x + clm.w
	bottom := clm.y + clm.h
	for x, y := clm.x, clm.y; y < bottom; {
		(*ccm)[y][x] += 1
		x += 1
		if x >= right {
			x = clm.x
			y += 1
		}
	}
}

func (intKeyMap *claimCountMap) sortedKeys() []int {
	keys := make([]int, 0, len(*intKeyMap))
	for k, _ := range *intKeyMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}

func sortedKeys(intKeyMap map[int]int) []int {
	keys := make([]int, 0, len(intKeyMap))
	for k, _ := range intKeyMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}

func (ccm *claimCountMap) print() {
	x, y := 0, 0
	for _, row := range *ccm {
		fmt.Print("\n")
		for _, count := range row {
			if count == 0 {
				fmt.Print(" ")
			} else if count > 9 {
				fmt.Print("*")
			} else {
				fmt.Print(count)
			}
			x += 1
		}
		y += 1
	}
	fmt.Print("\n")
}

func (clm *claim) hasOverlap(ccm *claimCountMap) bool {
	right := clm.x + clm.w
	bottom := clm.y + clm.h
	for x, y := clm.x, clm.y; y < bottom; {
		if (*ccm)[y][x] > 1 {
			return true
		}
		x += 1
		if x >= right {
			x = clm.x
			y += 1
		}
	}
	return false
}

func main() {
	ccm := new(claimCountMap)
	claims := make([]*claim, 0, 1337)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		theClaim, err := parseClaim(scanner.Text())
		if err != nil {
			fmt.Printf("ERROR: %s\n", err.Error())
			os.Exit(-1)
		}
		claims = append(claims, theClaim)
		ccm.updateMap(theClaim)
	}
	fmt.Println("The following claims have no overlap:")
	for _, claim := range claims {
		if !claim.hasOverlap(ccm) {
			fmt.Printf("%d\n", claim.id)
		}
	}
}
