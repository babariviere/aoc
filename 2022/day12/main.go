package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Pos struct {
	x, y int
}

type Path []Pos

type HeightMap map[Pos]int

func (h HeightMap) canClimb(curr, target Pos) bool {
	currVal := h[curr]
	targetVal := h[target]
	diff := targetVal - currVal
	return diff <= 1
}

func (h HeightMap) neighbour(p Pos) []Pos {
	var result []Pos

	for _, pos := range []Pos{{p.x + 1, p.y}, {p.x - 1, p.y}, {p.x, p.y + 1}, {p.x, p.y - 1}} {
		if _, ok := h[pos]; ok && h.canClimb(p, pos) {
			result = append(result, pos)
		}
	}
	return result
}

func (h HeightMap) shortestPath(start, end Pos) int {
	visited := make(map[Pos]bool)
	dist := make(map[Pos]int)
	pred := make(map[Pos]Pos)

	dist[start] = 0
	queue := []Pos{start}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		for _, neighbour := range h.neighbour(curr) {
			if seen := visited[neighbour]; !seen {
				fmt.Println("visiting", neighbour)
				visited[neighbour] = true
				dist[neighbour] = dist[curr] + 1
				pred[neighbour] = curr

				if neighbour == end {
					return dist[end]
				}

				queue = append(queue, neighbour)
			}
		}
	}
	return math.MaxInt
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var start, edge, end Pos
	elevationMap := make(HeightMap)
	var lineIdx int
	for scanner.Scan() {
		for i, val := range scanner.Text() {
			pos := Pos{x: i, y: lineIdx}
			if val == 'S' {
				start = pos
				elevationMap[pos] = 0
				continue
			}
			if val == 'E' {
				end = pos
				elevationMap[pos] = 'z' - 'a'
				continue
			}
			elevationMap[pos] = int(val - 'a')
		}

		edge.x = len(scanner.Text()) - 1
		edge.y = lineIdx
		lineIdx += 1
	}

	fmt.Println("start", start, "end", end)
	score := elevationMap.shortestPath(start, end)
	fmt.Println("Steps", score)
	minScore := math.MaxInt
	for pos, val := range elevationMap {
		if val == 0 {
			score = elevationMap.shortestPath(pos, end)
			if score < minScore {
				minScore = score
			}
		}
	}
	fmt.Println("Shortest path from a", minScore)
}
