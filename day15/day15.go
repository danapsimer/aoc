package main

import (
	"bufio"
	"bytes"
	"container/list"
	"fmt"
	"image"
	"io"
	"io/ioutil"
	"math"
	"os"
	"sort"
	"strings"
)

type CellType byte
type NPCType byte
type Action byte

const (
	Wall CellType = iota
	Open

	Elf NPCType = iota
	Goblin

	Up Action = iota
	Left
	Right
	Down
)

var (
	GoblinDamage = 3
	ElfDamage = 3
)

func (ct CellType) String() string {
	if ct == Wall {
		return "#"
	} else if ct == Open {
		return "."
	} else {
		return "!"
	}
}

func (nt NPCType) String() string {
	if nt == Elf {
		return "E"
	} else if nt == Goblin {
		return "G"
	} else {
		return "¡"
	}
}

type NPC struct {
	npcType NPCType
	hp      int
	cell    *Cell
	board   *Board
}

func NewNPC(npcType NPCType, cell *Cell) *NPC {
	return &NPC{npcType, 200, cell, cell.board}
}

func (npc *NPC) MoveTo(to *Cell) {
	npc.cell.npc = nil
	npc.cell = to
	to.npc = npc
}

func (npc *NPC) TakeDamage() {
	if npc.npcType == Elf {
		npc.hp -= GoblinDamage
	} else {
		npc.hp -= ElfDamage
	}
	if npc.hp <= 0 {
		npc.cell.npc = nil
		npc.board.RemoveNPC(npc)
	}
}

func (npc *NPC) TakeTurn() bool {
	target := npc.FindTarget()
	if target == nil {
		targets := npc.board.TargetsFor(npc)
		if len(targets) == 0 {
			return true
		}
		inrange := make([]*Cell, 0, len(targets)*4)
		for _, npc := range targets {
			for i := 0; i < len(dx); i++ {
				down := npc.board.Cell(npc.cell.X+dx[i], npc.cell.Y+dy[i])
				if down.isOpen() {
					inrange = append(inrange, down)
				}
			}
		}
		// sort so the first cell is the one that breaks ties.
		sort.Slice(inrange, ReadingOrder(inrange))
		shortestPathLength := int(math.MaxInt32)
		shortestPathCell := ([]*Cell)(nil)
		shortestPath := ([][]*Cell)(nil)
		for _, cell := range inrange {
			length, paths := npc.board.ShortestPathLength(npc.cell, cell)
			if length >= 0 && length < shortestPathLength {
				shortestPathLength = length
				shortestPathCell = make([]*Cell, 0, 10)
				shortestPath = make([][]*Cell, 0, 10)
				shortestPathCell = append(shortestPathCell, cell)
				shortestPath = append(shortestPath, paths...)
			} else if length == shortestPathLength {
				shortestPathCell = append(shortestPathCell, cell)
				shortestPath = append(shortestPath, paths...)
			}
		}
		if shortestPathCell != nil {
			firsts := make([]*Cell, 0, len(shortestPath))
			for _, path := range shortestPath {
				if len(path) > 0 && path[0].npc == nil {
					firsts = append(firsts, path[0])
				}
			}
			sort.SliceStable(firsts, ReadingOrder(firsts))
			if len(firsts) > 0 {
				npc.MoveTo(firsts[0])
			}
		}
		target = npc.FindTarget()
	}
	if target != nil {
		target.TakeDamage()
	}
	return false
}

func (npc *NPC) FindTarget() *NPC {
	target := (*NPC)(nil)
	for i := 0; i < len(dx); i++ {
		x, y := npc.cell.X+dx[i], npc.cell.Y+dy[i]
		cell := npc.board.Cell(x, y)
		if cell.npc != nil && cell.npc.npcType != npc.npcType && (target == nil || cell.npc.hp < target.hp) {
			target = cell.npc
		}
	}
	return target
}

func (npc *NPC) String() string {
	return fmt.Sprintf("%s(%d)",npc.npcType.String(), npc.hp)
}

type Cell struct {
	image.Point
	board    *Board
	cellType CellType
	npc      *NPC
}

func NewCell(x, y int, board *Board, cellType CellType) *Cell {
	return &Cell{image.Pt(x, y), board, cellType, nil}
}

func (cell *Cell) isOpen() bool {
	return cell.cellType == Open && cell.npc == nil
}

type Board struct {
	width int
	cells [][]*Cell
	npcs  []*NPC
}

func NewBoard(width int) *Board {
	return &Board{width, make([][]*Cell, 0, 100), make([]*NPC, 0, 100)}
}

func (b *Board) AddRow() {
	b.cells = append(b.cells, make([]*Cell, b.width))
}

func (b *Board) AddNPC(npc *NPC) {
	b.npcs = append(b.npcs, npc)
}

func (b *Board) TargetsFor(npc *NPC) []*NPC {
	targets := make([]*NPC, 0, len(b.npcs))
	for _, t := range b.npcs {
		if t != nil && t.npcType != npc.npcType {
			targets = append(targets, t)
		}
	}
	return targets
}

func ReadingOrder(cells []*Cell) func(i, j int) bool {
	return func(i, j int) bool {
		c1, c2 := cells[i], cells[j]
		if c1.Y < c2.Y {
			return true
		} else if c1.Y == c2.Y && c1.X < c2.X {
			return true
		}
		return false
	}
}

func (b *Board) DetermineTurnOrder() {
	// eliminate dead npcs
	for i := 0; i < len(b.npcs); {
		if b.npcs[i] == nil {
			b.npcs = append(b.npcs[:i], b.npcs[i+1:]...)
		} else {
			i += 1
		}
	}
	// Sort by reading order
	sort.Slice(b.npcs, func(i, j int) bool {
		n1, n2 := b.npcs[i], b.npcs[j]
		if n1.cell.Y < n2.cell.Y {
			return true
		} else if n1.cell.Y == n2.cell.Y && n1.cell.X < n2.cell.X {
			return true
		}
		return false
	})
}

var (
	dx = []int{0, -1, 1, 0}
	dy = []int{-1, 0, 0, 1}
)

func (b *Board) findPaths(depth int, depths map[*Cell]int, start, current *Cell, path []*Cell) [][]*Cell {
	if depth <= 0 {
		return nil
	}
	cpath := make([]*Cell, len(path), len(path)+1)
	copy(cpath, path)
	path = append(cpath, current)
	parents := b.GetAdjacent(current)
	paths := make([][]*Cell, 0, len(parents))
	for _, p := range parents {
		if p == start {
			for i := 0; i < len(path)/2; i++ {
				path[i], path[len(path)-i-1] = path[len(path)-i-1], path[i]
			}
			return [][]*Cell{path}
		} else if p.isOpen() && depths[p] == depth - 1 {
			rpaths := b.findPaths(depth-1, depths, start, p, path)
			if rpaths != nil {
				paths = append(paths, rpaths...)
			}
		}
	}
	return paths
}

func (b *Board) GetAdjacent(c *Cell) []*Cell {
	cells := make([]*Cell,4)
	for i := 0; i < len(dx); i++ {
		cells[i] = b.Cell(c.X+dx[i],c.Y+dy[i])
	}
	return cells
}

func (b *Board) ShortestPathLength(c1, c2 *Cell) (int, [][]*Cell) {
	visited := make(map[image.Point]bool)
	q := list.New()
	q.PushBack(c1)
	nodesLeftInLayer := 1
	nodesInNextLayer := 0
	moves := 0
	meta := make(map[*Cell]*Cell)
	depth := make(map[*Cell]int)
	depth[c1] = 0
	for q.Len() > 0 {
		p := q.Remove(q.Front()).(*Cell)
		if p == c2 {
			return moves, b.findPaths(moves, depth, c1, p, []*Cell{})
		} else {
			for _, cell := range b.GetAdjacent(p) {
				if !visited[cell.Point] && cell.isOpen() {
					depth[cell] = moves + 1
					meta[cell] = p
					q.PushBack(cell)
					visited[cell.Point] = true
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
	return -1, nil
}

func (b *Board) Cell(x, y int) *Cell {
	return b.cells[y][x]
}

func (b *Board) SetCell(x, y int, cell *Cell) {
	b.cells[y][x] = cell
}

func (b *Board) String() string {
	var buf bytes.Buffer
	for _, row := range b.cells {
		npcs := make([]*NPC,0,len(row))
		for _, c := range row {
			if c.npc != nil {
				npcs = append(npcs,c.npc)
			}
			if c.npc != nil {
				buf.WriteString(c.npc.npcType.String())
			} else {
				buf.WriteString(c.cellType.String())
			}
		}
		for _, npc := range npcs {
			buf.WriteString(fmt.Sprintf(" %s", npc.String()))
		}
		buf.WriteString("\n")
	}
	return string(buf.Bytes())
}

func (b *Board) RemoveNPC(npc *NPC) {
	for i, p := range b.npcs {
		if p == npc {
			b.npcs[i] = nil
		}
	}
}

func (b *Board) Turn() bool {
	b.DetermineTurnOrder()
	for _, npc := range b.npcs {
		if npc == nil {
			continue
		}
		if npc.TakeTurn() {
			return true
		}
	}
	return false
}

func (b *Board) SumHp() int {
	sum := 0
	for _, npc := range b.npcs {
		if npc != nil {
			sum += npc.hp
		}
	}
	return sum
}

func (b *Board) RunBattle() int {
	t := 1
	fmt.Printf("Initial:\n%s\n", b.String())
	for !b.Turn() {
		fmt.Printf("after turn #%d:\n%s\n", t, b.String())
		t += 1
	}
	fmt.Printf("final: #%d:\n%s\n", t, b.String())
	return (t-1) * b.SumHp()
}

func LoadBoard(reader io.Reader) *Board {
	board := (*Board)(nil)
	scanner := bufio.NewScanner(reader)
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		if board == nil {
			board = NewBoard(len(line))
		}
		board.AddRow()
		for x, c := range line {
			cell := (*Cell)(nil)
			switch {
			case c == '#':
				cell = NewCell(x, y, board, Wall)
				board.SetCell(x, y, cell)
			case c == 'E' || c == 'G' || c == '.':
				cell = NewCell(x, y, board, Open)
				board.SetCell(x, y, cell)
			default:
				panic(fmt.Sprintf("unknown cell type: %v\n", c))
			}
			npc := (*NPC)(nil)
			if c == 'E' {
				npc = NewNPC(Elf, cell)
			} else if c == 'G' {
				npc = NewNPC(Goblin, cell)
			}
			cell.npc = npc
			if npc != nil {
				board.AddNPC(npc)
			}
			x += 1
		}
		y += 1
	}
	return board
}

func (b* Board) ElfCount() int {
	sum := 0
	for _, npc := range b.npcs {
		if npc != nil && npc.npcType == Elf {
			sum += 1
		}
	}
	return sum
}

func main() {
	input := readReaderToString(os.Stdin)
	outcome := FindAttachPowerToWin(input)
	fmt.Printf("%d: %d\n", ElfDamage, outcome)
}

func readFileToString(fileName string) string {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()
	return readReaderToString(f)
}

func readReaderToString(reader io.Reader) string {
	inputBytes, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Printf("ERROR: reading input: %s", err.Error())
		os.Exit(-1)
	}
	input := string(inputBytes)
	return input
}

func FindAttachPowerToWin(input string) int {
	var outcome int
	for {
		board := LoadBoard(strings.NewReader(input))
		startingElfCount := board.ElfCount()
		fmt.Printf("Running battle with elf attack power at %d\n", ElfDamage)
		outcome = board.RunBattle()
		fmt.Printf("%d: start = %d, end = %d, outcome = %d\n",ElfDamage, startingElfCount, board.ElfCount(), outcome)
		if board.ElfCount() == startingElfCount {
			break
		}
		ElfDamage += 1
	}
	return outcome
}
