package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Material int

const (
	Air        Material = 0
	Rock                = 1
	Sand                = 2
	SandSource          = 3
)

type Pos struct {
	X, Y int
}

type Path []Pos

type Map map[Pos]Material

func (m Map) DrawLine(a, b Pos, mat Material) {
	dirX, dirY := 1, 1
	if a.X > b.X {
		dirX = -1
	}
	if a.Y > b.Y {
		dirY = -1
	}

	x, y := a.X, a.Y
	for {
		m[Pos{X: x, Y: y}] = mat

		if x == b.X && y == b.Y {
			break
		}
		if x != b.X {
			x += dirX
		}
		if y != b.Y {
			y += dirY
		}
	}
}

func (m Map) DrawPath(path Path, mat Material) {
	for i := 0; i < len(path)-1; i++ {
		m.DrawLine(path[i], path[i+1], mat)
	}
}

func (m Map) CanFall(sand Pos, floor int) (Pos, bool) {
	if sand.Y + 1 >= floor && floor > 0 {
		return sand, false
	}
	bot := Pos{X: sand.X, Y: sand.Y + 1}
	left := Pos{X: sand.X - 1, Y: sand.Y + 1}
	right := Pos{X: sand.X + 1, Y: sand.Y + 1}
	if mat, ok := m[bot]; !ok || mat == Air {
		return bot, true
	}
	if mat, ok := m[left]; !ok || mat == Air {
		return left, true
	}
	if mat, ok := m[right]; !ok || mat == Air {
		return right, true
	}
	return sand, false
}

// Returns true if sand has hit a "floor"
func (m Map) GenerateSand(source Pos, floor int) bool {
	pos := Pos{X: source.X, Y: source.Y}
	var canFall bool
	for {
		pos, canFall = m.CanFall(pos, floor)
		if !canFall {
			m[pos] = Sand
			return pos != source
		}
		if pos.Y > 500 {
			return false
		}
	}
}

func (m Map) Print(floor int) {
	var topLeft, bottomRight Pos
	for key := range m {
		if topLeft.X == 0 && topLeft.Y == 0 {
			topLeft = key
		}
		if bottomRight.X == 0 && bottomRight.Y == 0 {
			bottomRight = key
		}
		if key.X <= topLeft.X {
			topLeft.X = key.X
		}
		if key.Y <= topLeft.Y {
			topLeft.Y = key.Y
		}
		if key.X >= bottomRight.X {
			bottomRight.X = key.X
		}
		if key.Y >= bottomRight.Y {
			bottomRight.Y = key.Y
		}
	}

	rowSize := bottomRight.Y - topLeft.Y + 1
	colSize := bottomRight.X - topLeft.X + 1
	if floor > 0 {
		rowSize = floor + 1
	}
	l := make([][]Material, rowSize)
	for i := 0; i < rowSize; i++ {
		l[i] = make([]Material, colSize)
	}

	for pos, mat := range m {
		l[pos.Y - topLeft.Y][pos.X - topLeft.X] = mat
	}
	if floor > 0 {
	for x := 0; x < colSize; x++ {
		l[floor][x] = Rock
	}
	}

	for _, row := range l {
		for _, mat := range row {
			switch mat {
			case Air:
				fmt.Print(".")
			case Rock:
				fmt.Print("#")
			case Sand:
				fmt.Print("o")
			case SandSource:
				fmt.Print("+")
			}
		}
		fmt.Println()
	}
}

func main() {
	var paths []Path

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " -> ")
		var path Path
		for _, p := range line {
			list := strings.Split(p, ",")
			x, _ := strconv.Atoi(list[0])
			y, _ := strconv.Atoi(list[1])
			path = append(path, Pos{X: x, Y: y})
		}
		paths = append(paths, path)
	}

	m := make(Map)
	source :=Pos{X: 500, Y: 0}
	m[source] = SandSource
	for _, path := range paths {
		m.DrawPath(path, Rock)
	}

	m.Print(-1)
	for m.GenerateSand(source, -1) {
		m.Print(-1)
		fmt.Println()
	}

	m2 := make(Map)
	m2[source] = SandSource

	for _, path := range paths {
		m2.DrawPath(path, Rock)
	}
	var floor int
	for key := range m2 {
		if key.Y > floor {
			floor = key.Y
		}
	}
	floor += 2
	m2.Print(floor)
	for m2.GenerateSand(source, floor) {
		m2.Print(floor)
		fmt.Println()
	}

	var sands int
	for _, mat := range m {
		if mat == Sand {
			sands += 1
		}
	}
	var sands2 int
	for _, mat := range m2 {
		if mat == Sand {
			sands2 += 1
		}
	}
	fmt.Println("Unit of sands:", sands)
	fmt.Println("Unit of sands p2:", sands2)
}
