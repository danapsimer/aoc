package hexgrid

import "strings"

type Direction string

const (
	East      Direction = "e"
	West      Direction = "w"
	SouthEast Direction = "se"
	SouthWest Direction = "sw"
	NorthEast Direction = "ne"
	NorthWest Direction = "nw"
)

var directions = []Direction{East, West, SouthEast, SouthWest, NorthEast, NorthWest}

func NextDirection(s string) (Direction, string) {
	for _, d := range directions {
		if strings.HasPrefix(s, string(d)) {
			return d, s[len(d):]
		}
	}
	panic("invalid direction at start of string: " + s)
}

func (d Direction) Opposite() Direction {
	switch d {
	case East:
		return West
	case West:
		return East
	case SouthEast:
		return NorthWest
	case SouthWest:
		return NorthEast
	case NorthEast:
		return SouthWest
	case NorthWest:
		return SouthEast
	}
	panic("invalid direction: " + d)
}

func (d Direction) Clockwise() Direction {
	switch d {
	case East:
		return SouthEast
	case West:
		return NorthWest
	case SouthEast:
		return SouthWest
	case SouthWest:
		return West
	case NorthEast:
		return East
	case NorthWest:
		return NorthEast
	}
	panic("invalid direction: " + d)
}

func (d Direction) CounterClockwise() Direction {
	switch d {
	case East:
		return NorthEast
	case West:
		return SouthWest
	case SouthEast:
		return East
	case SouthWest:
		return SouthEast
	case NorthEast:
		return NorthWest
	case NorthWest:
		return West
	}
	panic("invalid direction: " + d)
}

type Node struct {
	Grid      *HexGrid
	Neighbors map[Direction]*Node
	Black     bool
}

type HexGrid struct {
	NodeCount      int
	BlackNodeCount int
	Origin         *Node
}

func NewHexGrid() *HexGrid {
	grid := &HexGrid{1, 0, &Node{nil, make(map[Direction]*Node), false}}
	grid.Origin.Grid = grid
	return grid
}

func (g *HexGrid) FindNode(instruction string) *Node {
	node := g.Origin
	for len(instruction) > 0 {
		nextDir, remainingInstruction := NextDirection(instruction)
		node = node.Neighbor(nextDir)
		instruction = remainingInstruction
	}
	return node
}

func (n *Node) NewNeighbor(direction Direction) *Node {
	neighbor := &Node{n.Grid, make(map[Direction]*Node), false}
	neighbor.Neighbors[direction.Opposite()] = n
	neighbor.Neighbors[direction.Clockwise().Opposite()] = n.Neighbors[direction.Clockwise()]
	neighbor.Neighbors[direction.CounterClockwise().Opposite()] = n.Neighbors[direction.CounterClockwise()]
	n.Neighbors[direction] = neighbor
	n.Grid.NodeCount += 1
	return neighbor
}

func (n *Node) Neighbor(direction Direction) *Node {
	neighbor, ok := n.Neighbors[direction]
	if !ok {
		neighbor = n.NewNeighbor(direction)
	}
	return neighbor
}

func (n *Node) Flip() {
	n.Black = !n.Black
	if n.Black {
		n.Grid.BlackNodeCount += 1
	} else {
		n.Grid.BlackNodeCount -= 1
	}
}
