package main

import (
	"bufio"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

const BorderWidth = 10

type Light struct {
	image.Point
	dx, dy int
}

type Lights struct {
	image.Rectangle
	lights []*Light
}

func NewLights() *Lights {
	return &Lights{image.Rectangle{image.Pt(math.MaxInt32, math.MaxInt32), image.Pt(0, 0)}, make([]*Light, 0, 100)}
}

func (lights *Lights) Add(light *Light) {
	lights.lights = append(lights.lights, light)
	if light.X > lights.Max.X {
		lights.Max.X = light.X
	}
	if light.X < lights.Min.X {
		lights.Min.X = light.X
	}
	if light.Y > lights.Max.Y {
		lights.Max.Y = light.Y
	}
	if light.Y < lights.Min.Y {
		lights.Min.Y = light.Y
	}
}

func (lights *Lights) Shift(delta image.Point) *Lights {
	newLights := NewLights()
	for _, light := range lights.lights {
		newLight := &Light{image.Pt(light.X+delta.X, light.Y+delta.Y), light.dx, light.dy}
		newLights.Add(newLight)
	}
	return newLights
}

func (lights *Lights) Iteration(n int) *Lights {
	newLights := NewLights()
	for _, light := range lights.lights {
		newLight := &Light{image.Pt(light.X+light.dx*n, light.Y+light.dy*n), light.dx, light.dy}
		newLights.Add(newLight)
	}
	return newLights
}

func (lights *Lights) Render() *image.Paletted {
	shifted := lights.Shift(image.Pt(BorderWidth-lights.Min.X, BorderWidth-lights.Min.Y))
	pallet := color.Palette{color.Black, color.White}
	img := image.NewPaletted(image.Rect(0, 0, shifted.Dx()+BorderWidth*2, shifted.Dy()+BorderWidth*2), pallet)
	for _, l := range shifted.lights {
		img.SetColorIndex(l.X, l.Y, 1)
	}
	return img
}

var (
	startIteration = 0
	iterations     = 100
	outputFile     = "output.png"
)

func init() {
	flag.IntVar(&startIteration, "-start", 0, "the iteration to start at")
	flag.IntVar(&startIteration, "s", 0, "the iteration to start at (shorthand)")
	flag.IntVar(&iterations, "-iterations", 100, "the number of iterations to search")
	flag.IntVar(&iterations, "i", 100, "the number of iterations to search (shorthand)")
	flag.StringVar(&outputFile, "-output", "output.png", "The name of the output file")
	flag.StringVar(&outputFile, "o", "output.png", "The name of the output file (shorthand)")
}

func main() {
	flag.Parse()
	fmt.Printf("Search second %d to %d and rendering smallest image.\n", startIteration+1, startIteration+iterations)
	scanner := bufio.NewScanner(os.Stdin)
	lights := NewLights()
	for scanner.Scan() {
		line := scanner.Text()
		var x, y, dx, dy int
		fmt.Sscanf(line, "position=<%d, %d> velocity=<%d, %d>", &x, &y, &dx, &dy)
		lights.Add(&Light{image.Pt(x, y), dx, dy})
	}
	smallest := lights
	smallestIdx := -1
	smallestArea := math.MaxInt32
	for i := startIteration; i < startIteration+iterations; i++ {
		iterated := lights.Iteration(i)
		area := iterated.Dx() * iterated.Dy()
		if smallest == nil || smallestArea > area {
			smallest = iterated
			smallestArea = area
			smallestIdx = i
		}
	}
	fmt.Printf("%d seconds have passed: generating image: %dx%d", smallestIdx, smallest.Dx(), smallest.Dy())
	image := smallest.Render()
	WriteImage(image, outputFile)
}

func WriteImage(image *image.Paletted, file string) {
	f, err := os.Create(file)
	if err != nil {
		fmt.Printf("ERROR: creating output file: %s\n", err.Error())
		os.Exit(-1)
	}
	defer f.Close()
	png.Encode(f, image)
}
