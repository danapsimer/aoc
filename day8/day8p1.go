package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	children []*Node
	metaData []int
}

func ScanInt(scanner *bufio.Scanner) (int, error) {
	var i int
	if scanner.Scan() {
		_, err := fmt.Sscanf(scanner.Text(),"%d", &i)
		if err != nil {
			return 0, fmt.Errorf("expected an integer, found %s: %s", scanner.Text(), err.Error())
		}
	} else {
		return 0, fmt.Errorf("expected input, found none")
	}
	return i, nil
}

func ReadNode(scanner *bufio.Scanner) (*Node,error) {
	childCount, err := ScanInt(scanner)
	if err != nil {
		return nil, fmt.Errorf("exepected a child count: %s", err.Error())
	}
	metaCount, err := ScanInt(scanner)
	if err != nil {
		return nil, fmt.Errorf("expected a meta count: %s", err.Error())
	}
	children := make([]*Node,0,childCount)
	for childCount > 0 {
		child, err := ReadNode(scanner)
		if err != nil {
			return nil, err
		}
		children = append(children,child)
		childCount -= 1
	}
	metadata := make([]int,metaCount)
	for metaCount > 0 {
		meta, err := ScanInt(scanner)
		if err != nil {
			return nil, fmt.Errorf("expected meta data: %s", err.Error())
		}
		metadata = append(metadata,meta)
		metaCount -= 1
	}
	return &Node{children,metadata}, nil
}

func (n *Node) sumMetadata() int {
	sum := 0
	for _, c := range n.children {
		sum += c.sumMetadata()
	}
	for _, m := range n.metaData {
		sum += m
	}
	return sum
}

func (n *Node) value() int {
	sum := 0
	if len(n.children) > 0 {
		for _, m := range n.metaData {
			if 0 < m && m <= len(n.children) {
				sum += n.children[m-1].value()
			}
		}
	} else {
		for _, m := range n.metaData {
			sum += m
		}
	}
	return sum
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	rootNode, err := ReadNode(scanner)
	if err != nil {
		fmt.Printf("ERROR: reading input: %s\n", err.Error())
		os.Exit(-1)
	}
	fmt.Printf("sum of root = %d\n",rootNode.sumMetadata())
	fmt.Printf("value of root = %d\n",rootNode.value())
}
