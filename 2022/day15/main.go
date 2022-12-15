package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Pos struct {
	X, Y int
}

func (p Pos) Distance(o Pos) int {
	return abs(p.X-o.X) + abs(p.Y-o.Y)
}

type Line struct {
	A, B Pos
}

type Scan struct {
	Scanner  Pos
	Beacon   Pos
	Distance int
}

func (s Scan) CollideY(y int) *Line {
	middle := Pos{X: s.Scanner.X, Y: y}
	distance := s.Scanner.Distance(middle)

	if distance > s.Distance {
		return nil
	}

	diff := s.Distance - distance
	left := Pos{X: s.Scanner.X - diff, Y: y}
	right := Pos{X: s.Scanner.X + diff, Y: y}
	return &Line{A: left, B: right}
}

func (s Scan) Hit(p Pos) bool {
	distance := s.Scanner.Distance(p)
	return distance <= s.Distance
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func parseScan(line string) Scan {
	re := regexp.MustCompile(`Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`)
	matches := re.FindStringSubmatch(line)
	scannerX, _ := strconv.Atoi(matches[1])
	scannerY, _ := strconv.Atoi(matches[2])
	beaconX, _ := strconv.Atoi(matches[3])
	beaconY, _ := strconv.Atoi(matches[4])

	scanner := Pos{
		X: scannerX,
		Y: scannerY,
	}
	beacon := Pos{
		X: beaconX,
		Y: beaconY,
	}
	return Scan{
		Scanner:  scanner,
		Beacon:   beacon,
		Distance: scanner.Distance(beacon),
	}
}

func main() {
	scanY := 10

	if len(os.Args) < 3 {
		fmt.Println("Usage: day15 <y> <max>")
		os.Exit(1)
	}

	scanY, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	max, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)
	var scans []Scan
	for scanner.Scan() {
		scan := parseScan(scanner.Text())
		scans = append(scans, scan)
		// top left, bottom right
		fmt.Println(scan)
	}

	collisions := make(map[Pos]bool)
	for _, scan := range scans {
		fmt.Println("For scanner", scan.Scanner)
		if line := scan.CollideY(scanY); line != nil {
			fmt.Println("Collision at", line)
			x := line.A.X
			for {
				col :=Pos{X: x, Y: line.A.Y}
				if col != scan.Beacon {
					collisions[col] = true
				}
				x += 1
				if x > line.B.X {
					break
				}
			}
		}
	}
	fmt.Println(collisions)
	fmt.Println(len(collisions))

	outsidePoints := make(map[Pos]bool)
	for y := 0; y < max; y++ {
		for _, scan := range scans {
			line := scan.CollideY(y)
			if line != nil {
				left := Pos{X: line.A.X - 1, Y: line.A.Y}
				right := Pos{X: line.B.X + 1, Y: line.B.Y}
				outsidePoints[left] = true
				outsidePoints[right] = true
			}
		}
	}

	for point := range outsidePoints {
		var hit bool
		for _, scan := range scans {
			hit = scan.Hit(point)
			if hit {
				break
			}
		}
		if !hit {
			if 0 <= point.X && point.X <= max && 0 <= point.Y && point.Y <= max {
				tuning := point.X*4000000 + point.Y
				fmt.Println("Tuning frequency", tuning)
				break
			}
		}
	}
}
