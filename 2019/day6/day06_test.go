package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var (
	in = `COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L`
	inPart02 = `COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L
K)YOU
I)SAN`
)

func TestDay06(t *testing.T) {
	assert.Equal(t, 42, Day06(strings.NewReader(in)))
}

func TestDay06Part2(t *testing.T) {
	assert.Equal(t, 4, Day06Part2(strings.NewReader(inPart02)))
}
