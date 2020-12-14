package main

import (
	"aoc/2020/day13/crt"
	"bufio"
	"errors"
	"io"
	"log"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func readNotes(reader io.Reader) (earliestTime int, busIds []int) {
	scanner := bufio.NewScanner(reader)
	busIds = make([]int, 0, 1000)
	if !scanner.Scan() {
		panic(errors.New("input too short"))
	}
	var err error
	earliestTime, err = strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	if !scanner.Scan() {
		panic(errors.New("input too short"))
	}
	idStrs := strings.Split(scanner.Text(), ",")
	for _, idStr := range idStrs {
		if idStr == "x" {
			busIds = append(busIds, -1)
		} else {
			busId, err := strconv.Atoi(idStr)
			if err != nil {
				panic(err)
			}
			busIds = append(busIds, busId)
		}
	}
	return
}

func main() {
	earliestTime, busIds := readNotes(os.Stdin)
	minTimeLeft := math.MaxInt32
	bestBusId := 0
	a := make([]*big.Int,0,1000)
	n := make([]*big.Int,0,1000)
	for idx, busId := range busIds {
		if busId > 0 {
			timeLeft := busId - earliestTime%busId
			if timeLeft < minTimeLeft {
				minTimeLeft = timeLeft
				bestBusId = busId
			}
			a = append(a, big.NewInt(int64(-(idx+1))))
			n = append(n, big.NewInt(int64(busId)))
		}
	}
	// cn (n>0) = t + n
	log.Printf("best bus is %d and requires %d minutes to wait. busId * mod = %d", bestBusId, minTimeLeft, bestBusId*minTimeLeft)
	answer, err := crt.Crt(a,n)
	if err != nil {
		panic(err)
	}
	answer.Add(answer,big.NewInt(1))
	log.Printf("part2 answer %s", answer.Text(10))
}
