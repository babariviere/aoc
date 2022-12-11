package main

import (
	"bufio"
	"fmt"
	// "math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Expr interface {
	Value(input int) int
}

type Number int

func (n Number) Value(input int) int { return int(n) }

type InputRef struct{}

func (i InputRef) Value(input int) int { return input }

func parseExpr(in string) Expr {
	if in == "old" {
		return InputRef{}
	}
	val, _ := strconv.Atoi(in)
	return Number(val)
}

type Operation struct {
	operator byte
	left     Expr
	right    Expr
}

func (o Operation) Exec(in int) int {
	switch o.operator {
	case '+':
		return o.left.Value(in) + o.right.Value(in)
	case '*':
		return o.left.Value(in) * o.right.Value(in)
	default:
		panic(fmt.Sprint("Unexpected operator", string(o.operator)))
	}
}

type Condition struct {
	divisibleBy   int
	onTrueTarget  int
	onFalseTarget int
}

func (c Condition) TestItem(item int, monkeys []*Monkey) {
	if item%c.divisibleBy == 0 {
		monkeys[c.onTrueTarget].Catch(item)
	} else {
		monkeys[c.onFalseTarget].Catch(item)
	}
}

type Monkey struct {
	name      string
	items     []int
	op        Operation
	condition Condition
	inspected int
}

func (m *Monkey) InspectItems(monkeys []*Monkey, denominator int, doMod bool) {
	for _, item := range m.items {
		m.inspected += 1
		worryLevel := m.op.Exec(item)
		if doMod {
			worryLevel %= denominator
		} else {
			worryLevel /= denominator
		}
		m.condition.TestItem(worryLevel, monkeys)
	}
	m.items = []int{}
}

func (m *Monkey) Catch(item int) {
	m.items = append(m.items, item)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var monkeys []*Monkey
	var idx int
	var currMonkey *Monkey
	for scanner.Scan() {
		if scanner.Text() == "" {
			if currMonkey != nil {
				monkeys = append(monkeys, currMonkey)
			}
			idx += 1
			currMonkey = nil
			continue
		}
		if currMonkey == nil {
			currMonkey = &Monkey{}
		}
		line := strings.Split(strings.TrimSpace(scanner.Text()), " ")
		switch line[0] {
		case "Monkey":
			currMonkey.name = strings.TrimSuffix(scanner.Text(), ":")
		case "Starting":
			items := strings.Split(strings.Join(line[2:], ""), ",")
			for _, item := range items {
				i, _ := strconv.Atoi(item)
				currMonkey.items = append(currMonkey.items, i)
			}
		case "Operation:":
			currMonkey.op.left = parseExpr(line[3])
			currMonkey.op.operator = line[4][0]
			currMonkey.op.right = parseExpr(line[5])

		case "Test:":
			div, _ := strconv.Atoi(line[3])
			currMonkey.condition.divisibleBy = div

		case "If":
			target, _ := strconv.Atoi(line[5])
			if line[1] == "true:" {
				currMonkey.condition.onTrueTarget = target
			} else {
				currMonkey.condition.onFalseTarget = target
			}

		default:
			panic(line[0])
		}
	}
	if currMonkey != nil {
		monkeys = append(monkeys, currMonkey)
	}
	// Part 1
	denominator := 3
	maxRound := 20
	doMod := false

	// Part 2
	// maxRound := 10000
	// denominator := 1
	// doMod := true
	// for _, monkey := range monkeys {
	// 	denominator *= monkey.condition.divisibleBy
	// }

	for round := 1; round <= maxRound; round++ {
		for _, monkey := range monkeys {
			monkey.InspectItems(monkeys, denominator, doMod)
		}
		fmt.Println("Round", round, "state:")
		for _, monkey := range monkeys {
			fmt.Println(monkey.name, monkey.items)
		}
	}

	inspected := make([]int, len(monkeys))
	for i, monkey := range monkeys {
		fmt.Println(monkey.name, "inspected", monkey.inspected, "items")
		inspected[i] = monkey.inspected
	}
	sort.Slice(inspected, func(i, j int) bool {
		return inspected[i] > inspected[j]
	})
	fmt.Println("monkey business", inspected[0]*inspected[1])
}
