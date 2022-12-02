package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, err := os.Open("fixtures/day_1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	elfs := make(map[int][]int)
	var calories []int
	elfCount := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if scanner.Text() != "" {
			num, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Println(err)
			}
			elfs[elfCount] = append(elfs[elfCount], num)
		} else {
			calories = append(calories, sum(elfs[elfCount]))
			elfCount++
		}
	}
	sort.Ints(calories)
	// part 1
	fmt.Println(calories[len(calories)-1])
	// part 2
	fmt.Println(sum(calories[len(calories)-3:]))

}

func sum(items []int) int {
	var calories int
	for _, item := range items {
		calories += item
	}
	return calories
}