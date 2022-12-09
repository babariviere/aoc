package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const LOST = 0
const DRAW = 3
const WON = 6

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var totalScore, totalScoreStrategy int
	for scanner.Scan() {
		commands := strings.Split(scanner.Text(), " ")
		var opponent, me int

		var score, strategyScore int
		// value between 0 and 2
		opponent = int(commands[0][0] - 'A')
		me = int(commands[1][0] - 'X')

		score = me + 1

		winMap := map[int]int{
			0: 2,
			1: 0,
			2: 1,
		}

		if opponent == me {
			score += DRAW
		} else if winMap[me] == opponent {
			score += WON
		} else {
			score += LOST
		}

		strat := me * 3
		choice := (me-1+3+opponent)%3 + 1
		strategyScore = strat + choice

		totalScoreStrategy += strategyScore
		totalScore += score
	}

	fmt.Println("Score", totalScore)
	fmt.Println("Strat score", totalScoreStrategy)
}
