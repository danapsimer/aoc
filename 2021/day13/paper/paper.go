package paper

import (
	"bytes"
	"fmt"
	"github.com/danapsimer/aoc/2021/utils"
	"math"
	"strconv"
	"strings"
)

type Direction int

const (
	Horizontal Direction = iota
	Vertical
)

type Fold struct {
	n   int
	dir Direction
}

type Paper map[int]map[int]bool

func (p Paper) Set(x, y int) {
	row, ok := p[y]
	if !ok {
		row = make(map[int]bool)
		p[y] = row
	}
	row[x] = true
}

func (p Paper) Unset(x, y int) {
	row, ok := p[y]
	if ok {
		delete(row, x)
	}
}

func (p Paper) IsSet(x, y int) bool {
	row, ok := p[y]
	if ok {
		_, ok := row[x]
		return ok
	}
	return false
}

func (p Paper) HorizontalFold(n int) {
	for srcRowIdx, srcRow := range p {
		if srcRowIdx > n {
			targetRowIdx := 2*n - srcRowIdx
			if targetRowIdx < 0 {
				panic(fmt.Errorf("horizontal folding at row %d resulted in negative target row: %d", n, targetRowIdx))
			}
			for x, _ := range srcRow {
				p.Set(x, targetRowIdx)
			}
			delete(p, srcRowIdx)
		}
	}
}

func (p Paper) VerticalFold(n int) {
	for _, row := range p {
		for x, _ := range row {
			if x > n {
				targetX := 2*n - x
				if targetX < 0 {
					panic(fmt.Errorf("vertical folding at %d resulted in negative target column: %d", n, targetX))
				}
				row[targetX] = true
				delete(row, x)
			}
		}
	}
}

func (p Paper) Fold(f *Fold) {
	if f.dir == Vertical {
		p.VerticalFold(f.n)
	} else {
		p.HorizontalFold(f.n)
	}
}

func (p Paper) CountDots() int {
	count := 0
	for _, row := range p {
		count += len(row)
	}
	return count
}

func (p Paper) XRange() (int, int) {
	minX := math.MaxInt
	maxX := 0
	for _, row := range p {
		for x, _ := range row {
			maxX = utils.IMax(maxX, x)
			minX = utils.IMin(minX, x)
		}
	}
	return minX, maxX
}

func (p Paper) YRange() (int, int) {
	minY := math.MaxInt
	maxY := 0
	for y, _ := range p {
		maxY = utils.IMax(maxY, y)
		minY = utils.IMin(minY, y)
	}
	return minY, maxY
}

func (p Paper) String() string {
	buf := &bytes.Buffer{}
	minY, maxY := p.YRange()
	minX, maxX := p.XRange()
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if p.IsSet(x, y) {
				fmt.Fprint(buf, "#")
			} else {
				fmt.Fprint(buf, " ")
			}
		}
		fmt.Fprintln(buf)
	}
	return buf.String()
}

func LoadPaperAndInstructions(lines <-chan string) (Paper, []*Fold, error) {
	readingDots := true
	paper := make(Paper)
	instructions := make([]*Fold, 0, 100)
	for line := range lines {
		if line == "" {
			readingDots = false
			continue
		}
		if readingDots {
			parts := strings.Split(line, ",")
			if len(parts) != 2 {
				return nil, nil, fmt.Errorf("unable to parse dot line: %s", line)
			}
			x, err := strconv.Atoi(parts[0])
			if err != nil {
				return nil, nil, err
			}
			y, err := strconv.Atoi(parts[1])
			if err != nil {
				return nil, nil, err
			}
			paper.Set(x, y)
		} else {
			parts := strings.Split(line, "=")
			n, err := strconv.Atoi(parts[1])
			if err != nil {
				return nil, nil, err
			}
			var dir Direction
			if parts[0][len(parts[0])-1] == 'y' {
				dir = Horizontal
			} else if parts[0][len(parts[0])-1] == 'x' {
				dir = Vertical
			} else {
				return nil, nil, fmt.Errorf("cannot parse instruction line: %s", line)
			}
			instructions = append(instructions, &Fold{n, dir})
		}
	}
	return paper, instructions, nil
}
