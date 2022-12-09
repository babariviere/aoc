package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	content, _ := ioutil.ReadAll(reader)

	var maxCal int

	leaderboard := []int{0, 0, 0, 0}
	for _, elf := range strings.Split(string(content), "\n\n") {
		var cal int
		for _, parsedCal := range strings.Split(elf, "\n") {
			i, _ := strconv.Atoi(parsedCal)
			cal += i
		}
		if cal > maxCal {
			maxCal = cal
		}
		leaderboard[3] = cal
		sort.Slice(leaderboard, func(i, j int) bool {
			return leaderboard[i] > leaderboard[j]
		})
	}

	fmt.Println("Result", maxCal)
	fmt.Println("Leaderboard", leaderboard[:3])
	fmt.Println("Total", leaderboard[0]+leaderboard[1]+leaderboard[2])
}
