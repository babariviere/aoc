package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	Direction byte
	Steps     int
}

type Pos struct {
	X, Y int
}

func (p *Pos) Follow(o Pos) {
	if p.X == o.X && p.Y == o.Y {
		return
	}

	diffX := o.X - p.X
	diffY := o.Y - p.Y
	dirX, dirY := 1, 1
	if diffX < 0 {
		dirX = -1
	}
	if diffY < 0 {
		dirY = -1
	}

	posX, posY := diffX*dirX, diffY*dirY
	if posX+posY >= 3 {
		p.X += 1 * dirX
		p.Y += 1 * dirY
		return
	}

	if posX >= 2 {
		p.X += 1 * dirX
	} else if posY >= 2 {
		p.Y += 1 * dirY
	}
}

func (p *Pos) Move(dir byte) {
	switch dir {
	case 'L':
		p.X -= 1
	case 'R':
		p.X += 1
	case 'U':
		p.Y -= 1
	case 'D':
		p.Y += 1
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var instructions []Instruction
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		dir := line[0][0]
		steps, _ := strconv.Atoi(line[1])
		instructions = append(instructions, Instruction{Direction: dir, Steps: steps})
	}

	visited := make(map[Pos]bool)
	visited2 := make(map[Pos]bool)
	head, tail := Pos{}, Pos{}
	rope := make([]Pos, 9)
	visited[tail] = true
	for _, instruction := range instructions {
		for step := 0; step < instruction.Steps; step++ {
			head.Move(instruction.Direction)
			for i := 0; i < len(rope); i++ {
				if i == 0 {
					rope[i].Follow(head)
				} else {
					rope[i].Follow(rope[i-1])
				}
			}
			visited2[rope[8]] = true
			tail.Follow(head)
			visited[tail] = true
		}
	}

	var count, count2 int
	for _, vis := range visited {
		if vis {
			count += 1
		}
	}
	for _, vis := range visited2 {
		if vis {
			count2 += 1
		}
	}
	fmt.Println("Visited", count)
	fmt.Println("Visited2", count2)
}
