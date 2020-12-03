package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"strconv"
)

func readExpenceReport(r io.Reader) (report []int, err error) {
	scanner := bufio.NewScanner(r)
	report = make([]int, 0, 1000)
	for scanner.Scan() {
		v := scanner.Text()
		var vi int
		vi, err = strconv.Atoi(v)
		if err != nil {
			return
		}
		report = append(report, vi)
	}
	return
}

func findSumsTo(sum int, report []int) (int, int, error) {
	for i := 0; i < len(report); i++ {
		diff := sum - report[i]
		if diff < 0 {
			continue
		}
		for j := i; j < len(report); j++ {
			if report[j] == diff {
				return i, j, nil
			}
		}
	}
	return 0, 0, errors.New("no 2 elements sum up to 2020")
}
func find3SumsTo(sum int, report []int) (int, int, int, error) {
	for i := 0; i < len(report); i++ {
		diff := sum - report[i]
		if diff <= 0 {
			continue
		}
		for j := i + 1; j < len(report); j++ {
			diff2 := diff - report[j]
			if diff2 <= 0 {
				continue
			}
			for k := j + 1; k < len(report); k++ {
				if report[k] == diff2 {
					return i, j, k, nil
				}
			}
		}
	}
	return 0, 0, 0, errors.New("no 3 elements sum up to 2020")
}

func main() {
	report, err := readExpenceReport(os.Stdin)
	if err != nil {
		log.Fatal("Error reading expence report", err.Error())
	}
	i, j, err := findSumsTo(2020, report)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("report[%d] * report[%d] = %d", i, j, report[i]*report[j])
	i, j, k, err := find3SumsTo(2020, report)
	log.Printf("report[%d] * report[%d] * report[%d] = %d", i, j, k, report[i]*report[j]*report[k])
}
