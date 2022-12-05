package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func pairRange(r string) []int {
	s := strings.Split(r, "-")
	first := getInt(s[0])
	second := getInt(s[1])
	var pair []int
	for i := first; i <= second; i++ {
		pair = append(pair, i)
	}
	return pair
}

func getInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func overlaps(s, d []int) bool {
	var found bool = false
	for _, num := range s {
		for _, n := range d {
			if num == n {
				found = true
			}
		}
	}
	return found
}

func contains(s, d []int) bool {
	var found bool = true
	m := make(map[int]bool)
	for _, num := range d {
		m[num] = true
	}
	for _, n := range s {
		if _, ok := m[n]; !ok {
			found = false
		}
	}
	return found
}

func main() {
	f, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var pairCount int
	var partialCount int
	s := bufio.NewScanner(f)
	for s.Scan() {
		pairs := strings.Split(s.Text(), ",")
		set1 := pairRange(pairs[0])
		set2 := pairRange(pairs[1])
		if contains(set1, set2) || contains(set2, set1) {
			pairCount++
		}
		if overlaps(set1, set2) || overlaps(set2, set1) {
			partialCount++
		}
		
	}
	fmt.Printf("Part 1: %d\n", pairCount)
	fmt.Printf("Part 2: %d\n", partialCount)
}