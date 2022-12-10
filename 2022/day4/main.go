package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	from int
	to   int
}

func NewRange(text string) Range {
	num := strings.Split(text, "-")
	from, _ := strconv.Atoi(num[0])
	to, _ := strconv.Atoi(num[1])
	return Range{
		from, to,
	}
}

func (r Range) FullyContains(b Range) bool {
	return r.from <= b.from && r.to >= b.to
}

func (r Range) Overlap(b Range) bool {
	return r.to >= b.from && r.from <= b.from
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var fullyContained, overlaps int
	for scanner.Scan() {
		ranges := strings.Split(scanner.Text(), ",")
		r1 := NewRange(ranges[0])
		r2 := NewRange(ranges[1])
		if r1.FullyContains(r2) || r2.FullyContains(r1) {
			fullyContained += 1
		}
		if r1.Overlap(r2) || r2.Overlap(r1) {
			overlaps += 1
		}
	}

	fmt.Println("fullyContained", fullyContained)
	fmt.Println("overlaps", overlaps)
}
