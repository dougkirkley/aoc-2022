package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

type Group []*Rucksack

func (g Group) String() string {
	s := strings.Builder{}
	for _, ruck := range g {
		s.WriteString(ruck.uniq())
	}
	return s.String()
}

func (g Group) Priority() int {
	for _, v := range g.String() {
		if strings.Count(g.String(), string(v)) == 3 {
			return strings.Index(alphabet, string(v)) + 1
		}
	}
	return 0
}

func rucksackToMap(rucksack string) map[string]bool {
	m := make(map[string]bool)
	for _, item := range rucksack {
		m[string(item)] = true
	}
	return m
}

type Rucksack struct {
	compartment1 Compartment
	compartment2 Compartment
}

func NewRucksack(items string) *Rucksack {
	// split rucksack in half
	compartmentCount := (len(items) / 2)

	return &Rucksack{
		compartment1: Compartment(items[0:compartmentCount]),
		compartment2: Compartment(items[compartmentCount:]),
	}
}

func (r Rucksack) String() string {
	var items []string
	items = append(items, r.compartment1.Items()...)
	items = append(items, r.compartment2.Items()...)
	return strings.Join(items, "")
}

// uniq returns only uniq characters found in string.
func (r Rucksack) uniq() string {
	s := &strings.Builder{}
	for _, sub := range r.String() {
		if !strings.Contains(s.String(), string(sub)) {
			s.WriteRune(sub)
		}
	}
	return s.String()
}

// Comparables returns items from both
func (r Rucksack) Comparables() []Priority {
	var both Priorities
	for _, i := range r.compartment1.Items() {
		for _, c := range r.compartment2.Items() {
			if i == c {
				if !strings.Contains(both.Join(), i) {
					p := NewPriority(i)
					both = append(both, p)
				}
			}
		}
	}
	return both
}

type Compartment string

func (c Compartment) Items() []string {
	var runes []string
	for _, r := range c {

		runes = append(runes, string(r))
	}
	return runes
}

type Priorities []Priority

func (p Priorities) Join() string {
	var joined []string
	for _, priority := range p {
		joined = append(joined, string(priority))
	}
	return strings.Join(joined, "")
}

type Priority string

func NewPriority(item string) Priority {
	return Priority(item)
}

func (p Priority) Int() int {
	return strings.Index(alphabet, string(p)) + 1
}

func main() {
	f, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	var sum int
	var groupSum int
	var groupCount int
	group := make(Group, 0)

	for s.Scan() {

		groupCount++
		r := NewRucksack(s.Text())
		group = append(group, r)

		if (groupCount % 3) == 0 {
			groupSum += group.Priority()
			group = Group{}
		}

		for _, item := range r.Comparables() {
			sum += item.Int()
		}
	}
	fmt.Printf("Part 1: %d\n", sum)
	fmt.Printf("Part 2: %d\n", groupSum)
}
