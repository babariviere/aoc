package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"

	mapset "github.com/deckarep/golang-set/v2"
)

func computePriority(items []byte) (totalPriority int) {
	for _, item := range items {
		var priority int
		if unicode.IsUpper(rune(item)) {
			priority = int(item - 'A' + 27)
		} else {
			priority = int(item - 'a' + 1)
		}
		fmt.Println(string(item), priority)
		totalPriority += priority
	}
	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var totalPriority, totalGlobalPriority int
	var groupSet mapset.Set[byte]

	var lineIdx int
	for scanner.Scan() {
		line := scanner.Text()
		mid := len(line) / 2
		fst, snd := line[:mid], line[mid:]

		fstSet, sndSet := mapset.NewSet[byte](), mapset.NewSet[byte]()

		for _, c := range []byte(fst) {
			fstSet.Add(c)
		}
		for _, c := range []byte(snd) {
			sndSet.Add(c)
		}

		{
			set := mapset.NewSet[byte]()
			for _, c := range []byte(line) {
				set.Add(c)
			}
			if (lineIdx % 3) == 0 {
				groupSet = set
			} else {
				groupSet = groupSet.Intersect(set)
			}
			if (lineIdx % 3) == 2 {
				fmt.Println("=====")
				fmt.Println("Group compute")
				items := groupSet.ToSlice()
				totalGlobalPriority += computePriority(items)
			}
		}

		fmt.Println("=====")
		fmt.Println("Normal compute")
		items := fstSet.Intersect(sndSet).ToSlice()
		totalPriority += computePriority(items)
		lineIdx += 1
	}

	fmt.Println("Total priority", totalPriority)
	fmt.Println("Total global priority", totalGlobalPriority)
}
