package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type CacheVal struct {
	Pressure int
	Path     []string
}

type Cache map[string]CacheVal

func (c Cache) key(curr string, opened []string, remainingTime int) string {
	sort.Strings(opened)

	return curr + strings.Join(opened, "") + strconv.Itoa(remainingTime)
}

func (c Cache) Add(curr string, opened []string, remainingTime int, val CacheVal) {
	c[c.key(curr, opened, remainingTime)] = val
}

func (c Cache) Get(curr string, opened []string, remainingTime int) (CacheVal, bool) {
	val, found := c[c.key(curr, opened, remainingTime)]
	return val, found
}

type Valve struct {
	Name         string
	Rate         int
	ConntectedTo []string
	Open         bool
}

func (v Valve) SimulatePressure(remainingTime int) int {
	return v.Rate * remainingTime
}

func (v *Valve) CalculateMaxPressure(valves map[string]*Valve, remainingTime int, cache Cache) (int, []string) {
	if remainingTime <= 0 {
		return 0, nil
	}

	var opened []string
	for _, valve := range valves {
		if valve.Open && valve.Rate > 0 {
			opened = append(opened, valve.Name)
		}
	}

	fmt.Println("Opened", opened)

	if best, ok := cache.Get(v.Name, opened, remainingTime); ok {
		fmt.Println("Cache hit!", v.Name, opened, remainingTime, "with value", best)
		return best.Pressure, best.Path
	}

	var best int
	var bestPath []string
	// Best value if we skip current valve
	for _, key := range v.ConntectedTo {
		valve := valves[key]
		b, path := valve.CalculateMaxPressure(valves, remainingTime-1, cache)
		if b > best {
			best = b
			bestPath = path
		}
	}

	if !v.Open && v.Rate > 0 && remainingTime > 0 {
		v.Open = true
		fmt.Println("Open", v.Name)

		pressure := v.SimulatePressure(remainingTime - 1)

		for _, key := range v.ConntectedTo {
			valve := valves[key]
			b, path := valve.CalculateMaxPressure(valves, remainingTime-2, cache)
			b += pressure
			if b > best {
				best = b
				bestPath = append([]string{v.Name}, path...)
			}
		}

		v.Open = false
		fmt.Println("Close", v.Name)
	}

	fmt.Println(v.Name, remainingTime, best)
	cache.Add(v.Name, opened, remainingTime, CacheVal{Pressure: best, Path: bestPath})
	return best, bestPath
}

func main() {
	valves := make(map[string]*Valve)

	scanner := bufio.NewScanner(os.Stdin)
	re := regexp.MustCompile(`Valve (\w+) has flow rate=(\d+); tunnels? leads? to valves? (.*)`)
	for scanner.Scan() {
		matches := re.FindStringSubmatch(scanner.Text())
		if len(matches) < 4 {
			continue
		}
		rate, _ := strconv.Atoi(matches[2])
		connectedTo := strings.Split(matches[3], ", ")
		valves[matches[1]] = &Valve{
			Name:         matches[1],
			Rate:         rate,
			ConntectedTo: connectedTo,
		}
	}

	currentValve := "AA"
	totalPressure, path := valves[currentValve].CalculateMaxPressure(valves, 30, make(Cache))
	fmt.Println("Total pressure", totalPressure, "for path", path)
}
