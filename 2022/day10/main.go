package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CPU struct {
	regX         int
	cycle        int
	instructions []ScheduledInstruction
}

func NewCPU() CPU {
	return CPU{
		regX:  1,
		cycle: 0,
	}
}

func (c *CPU) Register(i Instruction) {
	lastCycle := 0
	if len(c.instructions) > 0 {
		lastCycle = c.instructions[len(c.instructions)-1].cycle
	}
	scheduled := ScheduledInstruction{
		instruction: i,
		cycle:       lastCycle + i.Cycles(),
	}
	c.instructions = append(c.instructions, scheduled)
}

func (c *CPU) StartCycle() int {
	if len(c.instructions) == 0 {
		c.cycle += 1
		return c.cycle
	}
	c.cycle += 1
	return c.cycle
}

func (c *CPU) FinishCycle() {
	if len(c.instructions) == 0 {
		return
	}
	op := c.instructions[0]
	if c.cycle == op.cycle {
		fmt.Println("Exec", op.instruction, "in cycle", op.cycle)
		op.instruction.Exec(c)
		c.instructions = c.instructions[1:]
	}
}

func (c CPU) SignalStrength() int {
	fmt.Println(c.regX, c.cycle)
	return c.regX * c.cycle
}

func (c CPU) HasInstruction() bool {
	return len(c.instructions) > 0
}

type CRT struct {
	screen [6][40]bool
}

func NewCRT() CRT {
	return CRT{}
}

func (c *CRT) Draw(cpu CPU) {
	cycle := cpu.cycle - 1
	pos := cycle % 40
	spritePos := cpu.regX
	if pos >= spritePos-1 && pos <= spritePos+1 {
		row := cycle / 40
		c.screen[row][pos] = true
	}
}

func (c CRT) Print() {
	for _, row := range c.screen {
		for _, col := range row {
			if col {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

type ScheduledInstruction struct {
	instruction Instruction
	cycle       int
}

type Instruction interface {
	Exec(c *CPU)
	Cycles() int
}

type Noop struct{}

func (n Noop) Exec(c *CPU) {}
func (n Noop) Cycles() int { return 1 }

type Addx struct {
	num int
}

func (a Addx) Exec(c *CPU) {
	c.regX += a.num
}
func (a Addx) Cycles() int { return 2 }

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var instructions []Instruction
	for scanner.Scan() {
		instr := strings.Split(scanner.Text(), " ")
		switch instr[0] {
		case "noop":
			instructions = append(instructions, Noop{})
		case "addx":
			num, _ := strconv.Atoi(instr[1])
			instructions = append(instructions, Addx{
				num,
			})
		}
	}

	cpu := NewCPU()
	crt := NewCRT()
	for _, instruction := range instructions {
		cpu.Register(instruction)
	}

	var signalStrength int
	for cpu.HasInstruction() {
		cycle := cpu.StartCycle()
		if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
			signalStrength += cpu.SignalStrength()
		}
		crt.Draw(cpu)
		cpu.FinishCycle()
	}
	fmt.Println("Signal strength", signalStrength)
	crt.Print()
}
