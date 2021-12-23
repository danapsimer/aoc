package octo

type octopus struct {
	value   int
	flashed bool
}

func LoadOctopi(lines <-chan string) [][]*octopus {
	octopi := make([][]*octopus, 10)
	r := 0
	for line := range lines {
		octopi[r] = make([]*octopus, 10)
		for o, c := range line {
			octopi[r][o] = &octopus{int(c - '0'), false}
		}
		r += 1
	}
	return octopi
}

func RoundStep1(octopi [][]*octopus) {
	for _, row := range octopi {
		for _, octopus := range row {
			octopus.value += 1
		}
	}
}

func incAndFlashIf(octopi [][]*octopus, x, y int) {
	if x >= 0 && y >= 0 && y < len(octopi) && x < len(octopi[y]) && !octopi[y][x].flashed {
		octopi[y][x].value += 1
		flashIf(octopi, x, y)
	}
}

func flashIf(octopi [][]*octopus, x, y int) {
	if x >= 0 && y >= 0 && y < len(octopi) && x < len(octopi[y]) && !octopi[y][x].flashed && octopi[y][x].value > 9 {
		octopi[y][x].flashed = true
		incAndFlashIf(octopi, x+1, y+1)
		incAndFlashIf(octopi, x+1, y)
		incAndFlashIf(octopi, x+1, y-1)
		incAndFlashIf(octopi, x-1, y+1)
		incAndFlashIf(octopi, x-1, y)
		incAndFlashIf(octopi, x-1, y-1)
		incAndFlashIf(octopi, x, y-1)
		incAndFlashIf(octopi, x, y+1)
	}
}

func RoundStep2(octopi [][]*octopus) {
	for y, row := range octopi {
		for x, _ := range row {
			flashIf(octopi, x, y)
		}
	}
}

func RoundStep3(octopi [][]*octopus) int {
	count := 0
	for _, row := range octopi {
		for _, octopus := range row {
			if octopus.flashed {
				count += 1
				octopus.value = 0
				octopus.flashed = false
			}
		}
	}
	return count
}

func Round(octopi [][]*octopus) int {
	RoundStep1(octopi)
	RoundStep2(octopi)
	return RoundStep3(octopi)
}

func NRounds(n int, octopi [][]*octopus) int {
	count := 0
	for n > 0 {
		count += Round(octopi)
		n -= 1
	}
	return count
}
