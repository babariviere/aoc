package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Op struct {
	amount int
	from   int
	to     int
}

func parseOps(rawOps []string) []Op {
	var ops []Op
	for _, rawOp := range rawOps {
		list := strings.Split(rawOp, " ")
		amount, _ := strconv.Atoi(list[1])
		from, _ := strconv.Atoi(list[3])
		to, _ := strconv.Atoi(list[5])
		ops = append(ops, Op{
			amount: amount,
			from:   from - 1,
			to:     to - 1,
		})
	}
	return ops
}

func (o Op) Exec(stacks []Stack) {
	fmt.Println(stacks[o.from])
	for i := 0; i < o.amount; i++ {
		stacks[o.to].Push(stacks[o.from].Pop())
	}
}

func (o Op) Exec9001(stacks []Stack) {
	fmt.Println(stacks[o.from])
	bs := make([]byte, o.amount)
	for i := 0; i < o.amount; i++ {
		bs[o.amount-i-1] = stacks[o.from].Pop()
	}
	stacks[o.to].Extend(bs)
}

type Stack struct {
	list []byte
}

func (s *Stack) Push(b byte) {
	s.list = append(s.list, b)
}

func (s *Stack) Extend(bs []byte) {
	s.list = append(s.list, bs...)
}

func (s *Stack) Pop() byte {
	val := s.list[len(s.list)-1]
	s.list = s.list[:len(s.list)-1]
	return val
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func chunkString(text string, size int) [][]string {
	lines := strings.Split(text, "\n")

	var chunks [][]string
	for _, line := range lines {
		var chunk []string
		for i := 0; i < len(line); i += size {
			c := line[i:min(i+size, len(line))]
			chunk = append(chunk, c)
		}
		chunks = append(chunks, chunk)
	}
	return chunks
}

func parseStacks(text string) []Stack {
	chunks := chunkString(text, 4)
	size := len(chunks[0])

	stacks := make([]Stack, size)
	for a := len(chunks) - 1; a >= 0; a-- {
		chunk := chunks[a]
		for i, c := range chunk {
			if c[0] == '[' {
				stacks[i].Push(c[1])
			}
		}
	}
	return stacks
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var rawStacks, rawOps []string
	var isOps bool
	for scanner.Scan() {
		if isOps {
			rawOps = append(rawOps, scanner.Text())
			continue
		}

		if scanner.Text() == "" {
			isOps = true
			continue
		}
		rawStacks = append(rawStacks, scanner.Text())
	}

	stacks := parseStacks(strings.Join(rawStacks, "\n"))
	stacks9001 := parseStacks(strings.Join(rawStacks, "\n"))
	ops := parseOps(rawOps)
	for _, op := range ops {
		fmt.Println(op)
		op.Exec(stacks)
		op.Exec9001(stacks9001)
	}

	fmt.Print("Op 9000: ")
	for _, stack := range stacks {
		fmt.Print(string(stack.Pop()))
	}
	fmt.Println()

	fmt.Print("Op 9001: ")
	for _, stack := range stacks9001 {
		fmt.Print(string(stack.Pop()))
	}
	fmt.Println()
}
