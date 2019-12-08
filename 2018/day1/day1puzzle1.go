package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func contains(s []int, v int) bool{
	for _, e := range s {
		if e == v {
			return true
		}
	}
	return false
}

type repeatingStream struct {
	reader io.Reader
	pastEOF bool
	values []int
	pos int
}

func newRepeatingStream() *repeatingStream {
	return &repeatingStream{bufio.NewReader(os.Stdin), false, make([]int,0,1000), 0}
}

func (rs *repeatingStream) nextValue() (int, error) {
	if rs.pastEOF {
		if rs.pos >= len(rs.values) {
			rs.pos = 0
		}
		v := rs.values[rs.pos]
		rs.pos += 1
		return v, nil
	} else {
		v := 0
		for {
			i, err := fmt.Fscanf(rs.reader, "%d\n", &v)
			if err == nil && i == 0 {
				continue
			}
			if err == io.EOF {
				rs.pastEOF = true
				return rs.nextValue()
			}
			rs.values = append(rs.values,v)
			return v, err;
		}
	}
}

func main() {
	stream := newRepeatingStream()
	sum := 0
	sums := make([]int,0,1000)
	for {
		v, err := stream.nextValue()
		if err != nil {
			fmt.Printf("ERROR: %s\n",err.Error())
			break
		}
		sum += v
		if contains(sums,sum) {
			break
		}
		sums = append(sums,sum)
		if err == io.EOF {
			break
		}
	}
	fmt.Printf("%d\n", sum)
	os.Exit(0)
}
