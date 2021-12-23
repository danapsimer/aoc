package caves

import (
	"bytes"
	"container/list"
	"fmt"
	"strings"
)

type cave struct {
	name        string
	connections caves
}

type caves map[string]*cave

type path []*cave

func (p path) String() string {
	buf := &bytes.Buffer{}
	for i, pe := range p {
		if i > 0 {
			fmt.Print("->")
		}
		fmt.Print(pe.name)
	}
	return buf.String()
}

func (p path) Equal(other path) bool {
	if len(p) == len(other) {
		for i := range p {
			if p[i].name != other[i].name {
				return false
			}
		}
		return true
	}
	return false
}

func (cs caves) AddConnection(c1, c2 string) {
	cave1, ok := cs[c1]
	if !ok {
		cave1 = &cave{c1, make(caves)}
		cs[c1] = cave1
	}
	cave2, ok := cs[c2]
	if !ok {
		cave2 = &cave{c2, make(caves)}
		cs[c2] = cave2
	}
	cave1.connections[c2] = cave2
	cave2.connections[c1] = cave1
}

func (cs caves) Start() *cave {
	return cs["start"]
}

func (p path) countInPath(name string) int {
	count := 0
	for _, pe := range p {
		if pe.name == name {
			count += 1
		}
	}
	return count
}

func isUnique(queue *list.List, p path) bool {
	for e := queue.Front(); e != nil; e = e.Next() {
		if e.Value.(path).Equal(p) {
			return false
		}
	}
	return true
}

func (cs caves) SmallCaves() []string {
	sc := make([]string, 0, len(cs))
	for k, _ := range cs {
		if k[0] >= 'a' && k != "start" && k != "end" {
			sc = append(sc, k)
		}
	}
	return sc
}

func (cs caves) FindPaths(mayVisitTwice string) []path {
	completePaths := make([]path, 0, 100)
	queue := list.New()
	queue.PushBack(path{cs.Start()})
	for queue.Len() > 0 {
		front := queue.Front()
		queue.Remove(front)
		currPath := front.Value.(path)
		for _, connection := range currPath[len(currPath)-1].connections {
			newPath := make(path, len(currPath)+1)
			copy(newPath, currPath)
			newPath[len(currPath)] = connection
			if connection.name == "end" {
				completePaths = append(completePaths, newPath)
			} else {
				if connection.name[0] >= 'a' {
					count := currPath.countInPath(connection.name)
					if (connection.name == mayVisitTwice && count < 2) || (connection.name != mayVisitTwice && count < 1) {
						queue.PushBack(newPath)
					}
				} else {
					queue.PushBack(newPath)
				}
			}
		}
	}
	return completePaths
}

func (cs caves) FindPathWith1DuplicateSmallCave() []path {
	paths := make([]path, 0, 1000)
	for _, sc := range cs.SmallCaves() {
		for _, uniquePath := range cs.FindPaths(sc) {
			unique := true
			for _, pp := range paths {
				if pp.Equal(uniquePath) {
					unique = false
				}
			}
			if unique {
				paths = append(paths, uniquePath)
			}
		}
	}
	return paths
}

func LoadCaves(lines <-chan string) (caves, error) {
	cs := make(caves)
	for line := range lines {
		names := strings.Split(line, "-")
		if len(names) != 2 {
			return nil, fmt.Errorf("read %d elements instead of 2", len(names))
		}
		cs.AddConnection(names[0], names[1])
	}
	return cs, nil
}
