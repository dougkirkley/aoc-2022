package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	Rock     = 1
	Paper    = 2
	Scissors = 3
)

// objectMap has the relation of letters to object
var objectMap = map[string][]string{
	"rock": {
		"A", "X",
	},
	"paper": {
		"B", "Y",
	},
	"scissors": {
		"C", "Z",
	},
}

// loserMap shows which objects the key can lose to
var loserMap = map[string]string{
	"rock":     "paper",
	"paper":    "scissors",
	"scissors": "rock",
}

type Object string

func (o Object) String() string {
	for k, values := range objectMap {
		for _, value := range values {
			if value == string(o) {
				return k
			}
		}
	}
	return ""
}

// Round describes a round of the column 1 and column2 for the rock, paper or scissors
type Round struct {
	column1 Object
	column2 Object
}

func NewRound(line string) *Round {
	objects := strings.Split(line, " ")
	return &Round{
		column1: Object(objects[0]),
		column2: Object(objects[1]),
	}
}

func (r Round) Result() int {
	if r.column1.String() == r.column2.String() {
		return 3
	}
	// check if it's in the objects loser values
	if v, ok := loserMap[r.column1.String()]; ok {
		if r.column2.String() == v {
			return 0
		}
	} else {
		panic(fmt.Sprintf("%s not a valid choice", r.column1.String()))
	}
	return 6
}

func main() {
	f, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	var total int
	for s.Scan() {
		if s.Text() != "" {
			round := NewRound(s.Text())
			object := round.column2.String()
			r := round.Result()
			switch object {
			case "rock":
				result := r + Rock
				total += result
			case "paper":
				result := r + Paper
				total += result
			case "scissors":
				result := r + Scissors
				total += result
			}
		}
	}
	fmt.Printf("part 1 total is: %d\n", total)
}
