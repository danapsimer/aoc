package main

import (
	"bufio"
	"bytes"
	"container/list"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strings"
	"unicode"
)

var (
	failed = &trie{0, make(map[byte]*trie)}
)

type trie struct {
	v        byte
	children map[byte]*trie
}

func (t *trie) add(s string) {
	if len(s) == 0 {
		return
	}
	v := byte(s[0])
	c, ok := t.children[v]
	if !ok {
		c = &trie{v, make(map[byte]*trie)}
		t.children[v] = c
	}
	c.add(s[1:])
}

func (t *trie) containsToDepth(s string, depth int) int {
	if len(s) == 0 {
		return 0
	}
	v := byte(s[0])
	c, ok := t.children[v]
	if !ok {
		return depth
	}
	return c.containsToDepth(s[1:], depth+1)
}

type point struct {
	x, y int
}

type Grid struct {
	grid          []string
	allKeys       string
	distanceCache map[string]int
}

func keysString(keys map[byte]point) string {
	b := bytes.Buffer{}
	for k, _ := range keys {
		b.WriteByte(k)
	}
	return b.String()
}

func (grid *Grid) reachableKeys(currentPt point, keys map[byte]point, keysReached string) string {
	reachable := ""
	emptyCells := ".@" + grid.allKeys + strings.ToUpper(keysReached)
	for k, kpt := range keys {
		if ShortestPathLength(grid.grid, emptyCells, currentPt, kpt) > 0 {
			reachable += string(k)
		}
	}
}

func (grid *Grid) distanceToCollectKeys(currentKey byte, currentPt point, keys map[byte]point, keysReached string, cache map[string]int) int {
	if len(keys) == 0 {
		return 0
	}
	cacheKey := string(currentKey) + "|" + keysString(keys)
	result := 0
	ok := false
	if result, ok = cache[cacheKey]; ok {
		return result
	}

	result = math.MaxInt64
	reachable := grid.reachableKeys(currentPt, keys, keysReached)
	for _, k := range reachable {

	}

}

func Day18Part1(grid []string) (int, string) {
	p, ok := findPosition(grid, '@')
	if !ok {
		panic(fmt.Errorf("could not find position in %s", strings.Join(grid, "\n")))
	}
	allKeys := ""
	keys := make(map[byte]point)
	for c := byte('a'); c <= 'z'; c += 1 {
		k, ok := findPosition(grid, c)
		if !ok {
			break
		}
		keys[c] = k
		allKeys += string(c)
	}

	reached := ""
	perms := make([]string, 0, 100000)
	unreachedKeys := allKeys
	firstGroup := true
	for len(reached) < len(allKeys) {
		// Find reachable keys
		emptySquares := ".@" + allKeys + strings.ToUpper(reached)
		reachable := ""
		for i, k := range unreachedKeys {
			if ShortestPathLength(grid, emptySquares, p, keys[byte(k)]) != -1 {
				reachable = reachable + string(k)
				unreachedKeys = unreachedKeys[:i] + unreachedKeys[i+1:]
				reached += string(k)
			}
		}
		permChan := make(chan string)
		go permutations(reachable, 0, len(reachable)-1, permChan)
		if firstGroup {
			for perm := range permChan {
				perms = append(perms, perm)
			}
			firstGroup = false
		} else {
			permsLen := len(perms)
			for perm := range permChan {

				for i := 0; i < permsLen; i++ {
					perms = append(perms, perms[i]+perm)
				}
			}
			copy(perms, perms[permsLen:])
			perms = perms[:len(perms)-permsLen]
		}
	}
	log.Printf("trying permutations = %v", perms)
	shortestLength := math.MaxInt32
	shortestPerm := ""
	for _, perm := range perms {
		// skip permutations where a prefix has already failed
		if failed.containsToDepth(perm, 0) > 0 {
			continue
		}
		sum := 0
		curr := p
		emptySquares := ".@" + allKeys
		valid := true
		for i, c := range perm {
			kpt := keys[byte(c)]
			distance := ShortestPathLength(grid, emptySquares, curr, kpt)
			if distance < 0 {
				failed.add(perm[i:])
				valid = false
				break
			}
			sum += distance
			curr = kpt
			emptySquares += string(unicode.ToUpper(c))
		}
		if valid && shortestLength > sum {
			shortestLength = sum
			shortestPerm = perm
		}
	}
	return shortestLength, shortestPerm
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func swap(s string, p1, p2 int) string {
	if p1 == p2 {
		panic("swap called with equal values")
	}
	s1 := s[p1 : p1+1]
	s2 := s[p2 : p2+1]
	if p2 == len(s) {
		return s[:p1] + s2 + s[p1+1:p2] + s1
	}
	return s[:p1] + s2 + s[p1+1:p2] + s1 + s[p2+1:]
}

func permutations(s string, l, r int, out chan<- string) {
	if len(s) > 0 && failed.containsToDepth(s[:l], 0) > 0 {
		return
	}
	if l == r {
		out <- s
	} else {
		for i := l; i <= r; i++ {
			if l != i {
				s = swap(s, l, i)
			}
			permutations(s, l+1, r, out)
			if l != i {
				s = swap(s, l, i)
			}
		}
	}
	if l == 0 && r == len(s)-1 {
		close(out)
	}
}

type Move int

const (
	North Move = 1
	South Move = 2
	West  Move = 3
	East  Move = 4
)

func move(pos point, dir Move) point {
	switch dir {
	case North:
		pos.y -= 1
	case South:
		pos.y += 1
	case West:
		pos.x -= 1
	case East:
		pos.x += 1
	}
	return pos
}

func adjacent(grid []string, emptyCells string, p point) []point {
	ret := make([]point, 0, 4)
	for dir := North; dir <= East; dir += 1 {
		c := move(p, dir)
		if strings.Contains(emptyCells, string(grid[c.y][c.x])) {
			ret = append(ret, c)
		}
	}
	return ret
}

func ShortestPathLength(grid []string, distanceCache map[string]int, emptyCells string, c1, c2 point) int {
	cacheKey := fmt.Printf("(%d,%d)->(%d->%d)")
	visited := make(map[point]bool)
	q := list.New()
	q.PushBack(c1)
	nodesLeftInLayer := 1
	nodesInNextLayer := 0
	moves := 0
	meta := make(map[point]point)
	depth := make(map[point]int)
	depth[c1] = 0
	for q.Len() > 0 {
		p := q.Remove(q.Front()).(point)
		if p == c2 {
			return moves
		} else {
			for _, cell := range adjacent(grid, emptyCells, p) {
				if !visited[cell] {
					depth[cell] = moves + 1
					meta[cell] = p
					q.PushBack(cell)
					visited[cell] = true
					nodesInNextLayer += 1
				}
			}
		}
		nodesLeftInLayer -= 1
		if nodesLeftInLayer == 0 {
			nodesLeftInLayer = nodesInNextLayer
			nodesInNextLayer = 0
			moves += 1
		}
	}
	return -1
}

func findPosition(grid []string, c byte) (point, bool) {
	for y := 0; y < len(grid); y += 1 {
		for x := 0; x < len(grid[y]); x += 1 {
			if grid[y][x] == c {
				return point{x, y}, true
			}
		}
	}
	return point{}, false
}

func ReadGrid(in io.Reader) []string {
	scanner := bufio.NewScanner(in)
	grid := make([]string, 0, 1000)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}
	return grid
}

func main() {
	grid := ReadGrid(os.Stdin)
	steps, order := Day18Part1(grid)
	log.Printf("part1 = %d, %s", steps, order)
}
