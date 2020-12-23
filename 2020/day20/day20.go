package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/bits"
	"os"
	"strconv"
	"strings"
)

type Boarder int

const (
	North Boarder = iota
	East
	South
	West
)

func (b Boarder) String() string {
	switch b {
	case North:
		return "North"
	case East:
		return "East"
	case South:
		return "South"
	case West:
		return "West"
	}
	return ""
}

type Tile struct {
	Id      int
	Img     []uint16
	Boarder []uint16
	Edges   []*Tile
}

func NewTile(id int, img []uint16) *Tile {
	tile := &Tile{Id: id, Img: img, Boarder: make([]uint16, 4), Edges: make([]*Tile, 4)}
	for x := 0; x < 10; x++ {
		if x > 0 {
			for i := North; i <= West; i++ {
				tile.Boarder[i] <<= 1
			}
		}
		if img[0]&(0x1<<x) != 0 {
			tile.Boarder[North] |= 0x1
		}
		if img[x]&(0x1<<9) != 0 {
			tile.Boarder[East] |= 0x1
		}
		if img[9]&(0x1<<x) != 0 {
			tile.Boarder[South] |= 0x1
		}
		if img[x]&0x1 != 0 {
			tile.Boarder[West] |= 0x1
		}
	}
	return tile
}

func (t *Tile) Rotate() {
	newImg := make([]uint16, 10)
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			nx := 10 - y - 1
			ny := x
			mask := (t.Img[y] & (0x1 << x))
			if nx > x {
				newImg[ny] |= mask << (nx - x)
			} else {
				newImg[ny] |= mask >> (x - nx)
			}
		}
	}
	t.Img = newImg
	westEdge := t.Edges[West]
	westBoarder := t.Boarder[West]
	for i := West; i > North; i-- {
		t.Boarder[i] = t.Boarder[i-1]
		t.Edges[i] = t.Edges[i-1]
	}
	t.Boarder[North] = westBoarder
	t.Edges[North] = westEdge
}

func (t *Tile) VerticalFlip() {
	newImg := make([]uint16, 10)
	for y := range t.Img {
		newImg[y] = t.Img[10-y-1]
	}
	southEdge := t.Edges[South]
	t.Edges[South] = t.Edges[North]
	t.Edges[North] = southEdge

	south := t.Boarder[South]
	t.Boarder[South] = t.Boarder[North]
	t.Boarder[North] = south

	t.Boarder[East] = bits.Reverse16(t.Boarder[East]) >> 6
	t.Boarder[West] = bits.Reverse16(t.Boarder[West]) >> 6
}

func (t *Tile) HorizontalFlip() {
	newImg := make([]uint16, 10)
	for y := range t.Img {
		newImg[y] = bits.Reverse16(t.Img[y]) >> 6
	}
	eastEdge := t.Edges[East]
	t.Edges[East] = t.Edges[West]
	t.Edges[West] = eastEdge

	east := t.Boarder[East]
	t.Boarder[East] = t.Boarder[West]
	t.Boarder[West] = east

	t.Boarder[North] = bits.Reverse16(t.Boarder[North]) >> 6
	t.Boarder[South] = bits.Reverse16(t.Boarder[South]) >> 6
}

func CompareBoarder(i, j uint16) bool {
	ri := bits.Reverse16(i) >> 6
	rj := bits.Reverse16(j) >> 6
	return i == j || ri == rj || i == rj || ri == j
}

func (t *Tile) Compare(ot *Tile) (bool, Boarder, Boarder) {
	for i := North; i <= West; i++ {
		for j := North; j <= West; j++ {
			if CompareBoarder(t.Boarder[i], ot.Boarder[j]) {
				return true, i, j
			}
		}
	}
	return false, -1, -1
}

func (t *Tile) EdgeCount() int {
	count := 0
	for i := North; i <= West; i++ {
		if t.Edges[i] != nil {
			count += 1
		}
	}
	return count
}

func FindEdges(tiles []*Tile) {
	for i, t1 := range tiles {
		for j, t2 := range tiles {
			if i <= j {
				continue
			}
			if match, b1, b2 := t1.Compare(t2); match {
				//log.Printf("tile #%d's %s edge matches tile #%d's %s edge", t1.Id, b1.String(), t2.Id, b2.String())
				if t1.Edges[b1] != nil || t2.Edges[b2] != nil {
					panic("found multiple match for edges")
				}
				t1.Edges[b1] = t2
				t2.Edges[b2] = t1
			}
		}
	}
}

func FindCorners(tiles []*Tile) []*Tile {
	corners := make([]*Tile, 0, 4)
	for _, t := range tiles {
		if t.EdgeCount() == 2 {
			corners = append(corners, t)
		}
	}
	return corners
}

func ReadTiles(reader io.Reader) []*Tile {
	scanner := bufio.NewScanner(reader)
	tiles := make([]*Tile, 0, 144)
	img := make([]uint16, 0, 10)
	id := 0
	var err error
	for scanner.Scan() {
		line := scanner.Text()
		if len(strings.Trim(line, " \n\r\t")) == 0 {
			if len(img) == 10 {
				tiles = append(tiles, NewTile(id, img))
			}
			img = make([]uint16, 0, 10)
			id = 0
		} else {
			if strings.HasPrefix(line, "Tile ") {
				id, err = strconv.Atoi(line[5 : len(line)-1])
				if err != nil {
					panic(fmt.Errorf("invalid tile header: %s", line))
				}
			} else {
				bb := uint16(0)
				for bidx, r := range line {
					if r == '#' {
						bb |= 0x1 << bidx
					}
				}
				img = append(img, bb)
			}
		}
	}
	if len(img) == 10 {
		tiles = append(tiles, NewTile(id, img))
	}
	return tiles
}

func (t *Tile) OrientTile(pred func(*Tile) bool) {
	r := 0
	f := 0
	for r < 4 && !pred(t) {
		if f == 0 {
			t.HorizontalFlip()
			f += 1
		} else if f == 1 {
			t.HorizontalFlip()
			t.VerticalFlip()
			f += 1
		} else {
			t.VerticalFlip()
			t.Rotate()
			r += 1
			f = 0
		}
	}
	if !pred(t) {
		panic("failed to find way to orient tile")
	}
}

func (t *Tile) OrientTileWithSide(side Boarder, mask uint16) {
	t.OrientTile(func(tile *Tile) bool {
		return tile.Boarder[side] == mask
	})
}

func (t *Tile) OrientTileULCorner() {
	t.OrientTile(func(tile *Tile) bool {
		return tile.Boarder[West] == 0 && tile.Boarder[North] == 0
	})
}

func (t *Tile) OrientRow() {
	curr := t
	for curr.Edges[East] != nil {
		curr.Edges[East].OrientTileWithSide(West, t.Boarder[East])
		curr = curr.Edges[East]
	}
}

func OrientTiles(tiles []*Tile) *Tile {
	corners := FindCorners(tiles)
	ulc := corners[0]
	ulc.OrientTileULCorner()
	curr := ulc
	for {
		curr.OrientRow()
		if curr.Edges[South] == nil {
			break
		}
		curr = ulc.Edges[South]
		curr.OrientTileWithSide(West, 0)
	}
	return ulc
}

func StitchImage(ulc *Tile) [][]uint32 {
	height := 96
	image := make([][]uint32,0,height)

	y := 0
	row := ulc
	for row != nil {
		for i := 0; i < 8; i++ {
			image[y+i] = make([]uint32,3)
		}
		curr := row
		x := 0
		for curr != nil {
			for yy := 0; yy < 8; yy++ {
				mask := (curr.Img[yy+1] & 0b111111110) >> 1
				shift := 24 - x % 32
				image[y+yy][2-x/32] |= uint32(mask) << shift
			}
			x += 8
			curr = curr.Edges[East]
		}

		y += 8
		row = row.Edges[South]
	}
	return image
}

func main() {

	tiles := ReadTiles(os.Stdin)
	log.Printf("len(tiles) = %d", len(tiles))
	FindEdges(tiles)
	corners := FindCorners(tiles)
	if len(corners) != 4 {
		panic(fmt.Errorf("there should be 4 corners. Found %d", len(corners)))
	}
	m := 1
	for _, corner := range corners {
		m *= corner.Id
	}
	log.Printf("m = %d", m)
	ulc := OrientTiles(tiles)
	image := StitchImage(ulc)
	bldr := strings.Builder{}
  for y := range image {
  	if y > 0 {
  		bldr.WriteString("\n")
		}
  	for x := 0; x < 96; x++ {
  		if image[y][2-x/32] & 0x1 << (24-x%32) != 0 {
  			bldr.WriteString("#")
			} else {
				bldr.WriteString(" ")
			}
		}
	}
	log.Printf("Image = %s", bldr.String())
}
