package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

const testInput = `light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`

func TestReadContainerRules(t *testing.T) {

	rules := readContainerRules(strings.NewReader(testInput))
	assert.NotNil(t, rules)
	assert.Equal(t, 1, rules["bright white"].CanContainBag["shiny gold"])
	assert.Equal(t, 2, rules["muted yellow"].CanContainBag["shiny gold"])
	assert.Equal(t, 3, rules["dark orange"].CanContainBag["bright white"])
	assert.Equal(t, 4, rules["dark orange"].CanContainBag["muted yellow"])
	assert.Equal(t, 2, rules["light red"].CanContainBag["muted yellow"])
	assert.Equal(t, 1, rules["light red"].CanContainBag["bright white"])
}
