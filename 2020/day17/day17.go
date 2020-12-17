package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type Space1D struct {
	x        map[int]bool
	min, max int
}

func NewSpace1D() *Space1D {
	return &Space1D{make(map[int]bool), 0, 0}
}

func (s Space1D) Get(x int) bool {
	return s.x[x]
}

func (s Space1D) Set(x int, active bool) {
	if active {
		s.x[x] = active
	} else {
		delete(s.x, x)
	}
}

func (s *Space1D) Range() (int, int) {
	return s.min, s.max
}

func (s *Space1D) ActiveCount() int {
	return len(s.x)
}

type Space2D struct {
	y                      map[int]*Space1D
	minX, maxX, minY, maxY int
}

func NewSpace2D() *Space2D {
	return &Space2D{make(map[int]*Space1D), 0, 0, 0, 0}
}

func (s *Space2D) Get(x, y int) bool {
	s1d, ok := s.y[y]
	if ok {
		return s1d.Get(x)
	}
	return false
}

func (s *Space2D) Set(x, y int, active bool) {
	s1d, ok := s.y[y]
	if !ok {
		s1d = NewSpace1D()
		s.y[y] = s1d

	}
	s1d.Set(x, active)
	if y < s.minY {
		s.minY = y
	} else if y > s.maxY {
		s.maxY = y
	}
	if x < s.minX {
		s.minX = x
	} else if x > s.maxX {
		s.maxX = x
	}
}

func (s *Space2D) Range() (int, int, int, int) {
	return s.minX, s.maxX, s.minY, s.maxY
}

func (s *Space2D) SubSpace(y int) *Space1D {
	return s.y[y]
}

func (s *Space2D) ActiveCount() int {
	activeCount := 0
	for _, xs := range s.y {
		activeCount += xs.ActiveCount()
	}
	return activeCount
}

type Space3D struct {
	z                                  map[int]*Space2D
	minZ, maxZ, minY, maxY, minX, maxX int
}

func NewSpace3D(s2d *Space2D) *Space3D {
	s := &Space3D{make(map[int]*Space2D), 0, 0, 0, 0, 0, 0}
	if s2d != nil {
		s.z[0] = s2d
		s.minX, s.maxX, s.minY, s.maxY = s2d.Range()
	}
	return s
}

func (s *Space3D) Get(x, y, z int) bool {
	s2d, ok := s.z[z]
	if ok {
		return s2d.Get(x, y)
	}
	return false
}

func (s *Space3D) Set(x, y, z int, active bool) {
	s2d, ok := s.z[z]
	if !ok {
		s2d = NewSpace2D()
		s.z[z] = s2d

	}
	s2d.Set(x, y, active)
	if z < s.minZ {
		s.minZ = z
	} else if z > s.maxZ {
		s.maxZ = z
	}
	if y < s.minY {
		s.minY = y
	} else if y > s.maxY {
		s.maxY = y
	}
	if x < s.minX {
		s.minX = x
	} else if x > s.maxX {
		s.maxX = x
	}
}

func (s *Space3D) SubSpace(z int) *Space2D {
	return s.z[z]
}

func (s *Space3D) Range() (int, int, int, int, int, int) {
	return s.minX, s.maxX, s.minY, s.maxY, s.minZ, s.maxZ
}

func (s *Space3D) ActiveCount() int {
	activeCount := 0
	for _, ys := range s.z {
		activeCount += ys.ActiveCount()
	}
	return activeCount
}

func (s *Space3D) String() string {
	sw := new(bytes.Buffer)
	for z := s.minZ; z <= s.maxZ; z++ {
		fmt.Fprintf(sw, "z = %d\n", z)
		for y := s.minY; y <= s.maxY; y++ {
			fmt.Fprintf(sw, "% 5d: ", y)
			for x := s.minX; x <= s.maxX; x++ {
				if s.Get(x, y, z) {
					sw.WriteString("#")
				} else {
					sw.WriteString(".")
				}
			}
			sw.WriteString("\n")
		}
		sw.WriteString("\n")
	}
	return sw.String()
}

func (s *Space3D) ExecuteCycle() *Space3D {
	out := NewSpace3D(nil)
	minX, maxX, minY, maxY, minZ, maxZ := s.minX-1, s.maxX+1, s.minY-1, s.maxY+1, s.minZ-1, s.maxZ+1
	for z := minZ; z <= maxZ; z++ {
		for y := minY; y <= maxY; y++ {
			for x := minX; x <= maxX; x++ {
				activeCount := 0
				for dx := -1; dx <= 1; dx++ {
					for dy := -1; dy <= 1; dy++ {
						for dz := -1; dz <= 1; dz++ {
							if dx == 0 && dy == 0 && dz == 0 {
								continue
							}
							if s.Get(x+dx, y+dy, z+dz) {
								activeCount += 1
							}
						}
					}
				}
				active := s.Get(x, y, z)
				if active {
					if activeCount == 2 || activeCount == 3 {
						out.Set(x, y, z, true)
					}
				} else {
					if activeCount == 3 {
						out.Set(x, y, z, true)
					}
				}
			}
		}
	}
	return out
}

type Space4D struct {
	w                                              map[int]*Space3D
	minW, maxW, minZ, maxZ, minY, maxY, minX, maxX int
}

func NewSpace4D(s3d *Space3D) *Space4D {
	s := &Space4D{make(map[int]*Space3D), 0, 0, 0, 0, 0, 0, 0, 0}
	if s3d != nil {
		s.w[0] = s3d
		s.minX, s.maxX, s.minY, s.maxY, s.minZ, s.maxZ = s3d.Range()
	}
	return s
}

func (s *Space4D) Get(x, y, z, w int) bool {
	s3d, ok := s.w[w]
	if ok {
		return s3d.Get(x, y, z)
	}
	return false
}

func (s *Space4D) Set(x, y, z, w int, active bool) {
	s3d, ok := s.w[w]
	if !ok {
		s3d = NewSpace3D(nil)
		s.w[w] = s3d
	}
	s3d.Set(x, y, z, active)
	if w < s.minW {
		s.minW = w
	} else if w > s.maxW {
		s.maxW = w
	}
	if z < s.minZ {
		s.minZ = z
	} else if z > s.maxZ {
		s.maxZ = z
	}
	if y < s.minY {
		s.minY = y
	} else if y > s.maxY {
		s.maxY = y
	}
	if x < s.minX {
		s.minX = x
	} else if x > s.maxX {
		s.maxX = x
	}
}

func (s *Space4D) SubSpace(w int) *Space3D {
	return s.w[w]
}

func (s *Space4D) Range() (int, int, int, int, int, int, int, int) {
	return s.minX, s.maxX, s.minY, s.maxY, s.minZ, s.maxZ, s.minW, s.maxW
}

func (s *Space4D) ActiveCount() int {
	activeCount := 0
	for _, zs := range s.w {
		activeCount += zs.ActiveCount()
	}
	return activeCount
}

func (s *Space4D) String() string {
	sw := new(bytes.Buffer)
	for w := s.minW; w <= s.maxW; w++ {
		for z := s.minZ; z <= s.maxZ; z++ {
			fmt.Fprintf(sw, "z = %d, w = %d\n", z, w)
			for y := s.minY; y <= s.maxY; y++ {
				fmt.Fprintf(sw, "% 5d: ", y)
				for x := s.minX; x <= s.maxX; x++ {
					if s.Get(x, y, z, w) {
						sw.WriteString("#")
					} else {
						sw.WriteString(".")
					}
				}
				sw.WriteString("\n")
			}
			sw.WriteString("\n")
		}
	}
	return sw.String()
}

func (s *Space4D) ExecuteCycle() *Space4D {
	out := NewSpace4D(nil)
	minX, maxX, minY, maxY, minZ, maxZ, minW, maxW := s.minX-1, s.maxX+1, s.minY-1, s.maxY+1, s.minZ-1, s.maxZ+1, s.minW-1, s.maxW+1
	for w := minW; w <= maxW; w++ {
		for z := minZ; z <= maxZ; z++ {
			for y := minY; y <= maxY; y++ {
				for x := minX; x <= maxX; x++ {
					activeCount := 0
					for dx := -1; dx <= 1; dx++ {
						for dy := -1; dy <= 1; dy++ {
							for dz := -1; dz <= 1; dz++ {
								for dw := -1; dw <= 1; dw++ {
									if dx == 0 && dy == 0 && dz == 0 && dw == 0 {
										continue
									}
									if s.Get(x+dx, y+dy, z+dz, w+dw) {
										activeCount += 1
									}
								}
							}
						}
					}
					active := s.Get(x, y, z, w)
					if active {
						if activeCount == 2 || activeCount == 3 {
							out.Set(x, y, z, w, true)
						}
					} else if activeCount == 3 {
						out.Set(x, y, z, w, true)
					}
				}
			}
		}
	}
	return out
}

func ReadInput(reader io.Reader) *Space2D {
	s := NewSpace2D()
	scanner := bufio.NewScanner(reader)
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		for x, r := range line {
			s.Set(x, y, r == '#')
		}
		y += 1
	}
	return s
}

func main() {
	inputGrid := ReadInput(os.Stdin)
	current3D := NewSpace3D(inputGrid)
	current4D := NewSpace4D(current3D)
	//log.Printf("part1 initial:\n%s", current3D.String())
	for c := 0; c < 6; c++ {
		current3D = current3D.ExecuteCycle()
		//log.Printf("part1 cycle %d:\n%s", c+1, current3D.String())
	}
	log.Printf("part1 active count after 6 cyles: %d", current3D.ActiveCount())

	//log.Printf("part2 initial:\n%s", current4D.String())
	for c := 0; c < 6; c++ {
		current4D = current4D.ExecuteCycle()
		//log.Printf("part2 cycle %d:\n%s", c+1, current4D.String())
	}
	log.Printf("part2 active count after 6 cyles: %d", current4D.ActiveCount())
}
