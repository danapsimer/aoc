package diagnostic

type node struct {
	Count    int
	Value    uint
	Branches []*node
}

func newNode(b uint) *node {
	return &node{0, b, make([]*node, 2)}
}

func newTree() *node {
	return newNode(0)
}

func (n *node) add(value uint, width int) {
	n.Count += 1
	if width > 0 {
		b := (value & (1 << (width - 1))) >> (width - 1)
		if n.Branches[b] == nil {
			n.Branches[b] = newNode(b)
		}
		n.Branches[b].add(value, width-1)
	}
}

func (n *node) traverse(width int, rule func(*node) uint) uint {
	b := rule(n)
	tail := uint(0)
	if width > 1 && n.Branches[b] != nil {
		tail = n.Branches[b].traverse(width-1, rule)
	}
	return (b << (width - 1)) | tail
}

func (n *node) branchCounts() []int {
	bc := make([]int, len(n.Branches))
	for i, b := range n.Branches {
		bc[i] = 0
		if b != nil {
			bc[i] = b.Count
		}
	}
	return bc
}

func (n *node) oxyGenRating(width int) uint {
	return n.traverse(width, func(nn *node) uint {
		branchCounts := nn.branchCounts()
		if branchCounts[0] > branchCounts[1] {
			return 0
		} else {
			return 1
		}
	})
}

func (n *node) co2ScrubRating(width int) uint {
	return n.traverse(width, func(nn *node) uint {
		branchCounts := nn.branchCounts()
		if branchCounts[0] <= branchCounts[1] && nn.Branches[0] != nil {
			return 0
		} else if nn.Branches[1] != nil {
			return 1
		} else {
			return 0
		}
	})
}

func loadTree(values <-chan uint, width int) *node {
	root := newTree()
	for v := range values {
		root.add(v, width)
	}
	return root
}

func CalculateRatings(values <-chan uint, width int) (uint, uint) {
	tree := loadTree(values, width)
	return tree.oxyGenRating(width), tree.co2ScrubRating(width)
}
