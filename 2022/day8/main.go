package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Grid [][]int

type Pos struct {
	X, Y int
}

func visibleFromOutside(grid Grid, pos Pos) bool {
	// y == 0 -> y
	// len(grid) -> y
	// x == 0 -> x
	// len(grid[0]) -> x
	val := grid[pos.Y][pos.X]
	for y := 0; y <= pos.Y; y++ {
		if y == pos.Y {
			return true
		}
		if grid[y][pos.X] >= val {
			break
		}
	}
	for y := len(grid) - 1; y >= pos.Y; y-- {
		if y == pos.Y {
			return true
		}
		if grid[y][pos.X] >= val {
			break
		}
	}

	for x := 0; x <= pos.X; x++ {
		if x == pos.X {
			return true
		}
		if grid[pos.Y][x] >= val {
			break
		}
	}
	for x := len(grid[0]) - 1; x >= pos.X; x-- {
		if x == pos.X {
			return true
		}
		if grid[pos.Y][x] >= val {
			break
		}
	}
	return false
}

func calculateScenicScore(grid Grid, pos Pos) int {
	fmt.Print("Scenic score for ", pos.X, " ", pos.Y, ": ")
	var count int
	scenicScore := 1
	val := grid[pos.Y][pos.X]
	for y := pos.Y - 1; y >= 0; y-- {
		count += 1
		if grid[y][pos.X] >= val {
			break
		}
	}
	scenicScore *= count
	fmt.Print(count, " * ")
	count = 0
	for y := pos.Y + 1; y < len(grid); y++ {
		count += 1
		if grid[y][pos.X] >= val {
			break
		}
	}
	scenicScore *= count
	fmt.Print(count, " * ")
	count = 0
	for x := pos.X - 1; x >= 0; x-- {
		count += 1
		if grid[pos.Y][x] >= val {
			break
		}
	}
	scenicScore *= count
	fmt.Print(count, " * ")
	count = 0
	for x := pos.X + 1; x < len(grid[0]); x++ {
		count += 1
		if grid[pos.Y][x] >= val {
			break
		}
	}
	scenicScore *= count
	fmt.Print(count)
	count = 0
	fmt.Println(" =", scenicScore)
	return scenicScore
}

func main() {
	var grid Grid
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		nums := strings.Split(scanner.Text(), "")
		trees := make([]int, len(nums))
		for i, num := range nums {
			trees[i], _ = strconv.Atoi(num)
		}
		grid = append(grid, trees)
	}

	visibleTrees := 0
	maxScenicScore := 0
	for y, row := range grid {
		for x := range row {
			pos := Pos{X: x, Y: y}
			if visibleFromOutside(grid, pos) {
				visibleTrees += 1
			}
			scenicScore := calculateScenicScore(grid, pos)
			if scenicScore > maxScenicScore {
				maxScenicScore = scenicScore
			}
		}
	}

	fmt.Println("Visible", visibleTrees)
	fmt.Println("Scenic score", maxScenicScore)
}
