package utils

import (
	"bufio"
	"context"
	"io"
	"log"
	"os"
	"strconv"
)

func ReadIntegersFromReader(ctx context.Context, r io.Reader) <-chan int {
	intCh := make(chan int)
	go func() {
		defer close(intCh)
		strCh := ReadLinesFromReader(r)
		for vstr := range strCh {
			v, err := strconv.Atoi(vstr)
			if err != nil {
				log.Default().Printf("error converting input %s to integer: %s", vstr, err.Error())
				return
			}
			intCh <- v
		}
	}()
	return intCh
}

func ReadLinesFromReader(f io.Reader) <-chan string {
	lines := make(chan string)
	go func() {
		defer close(lines)
		s := bufio.NewScanner(f)
		for s.Scan() {
			lines <- s.Text()
		}
	}()
	return lines
}

func ReadLinesFromFile(filename string) (<-chan string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return ReadLinesFromReader(f), nil
}

func ReadIntegersFromFile(filename string) (<-chan int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return ReadIntegersFromReader(context.TODO(), f), nil
}

func StringArrayToChannel(arr []string) <-chan string {
	intCh := make(chan string)
	go func() {
		defer close(intCh)
		for _, v := range arr {
			intCh <- v
		}
	}()
	return intCh
}
