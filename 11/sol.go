package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	Inspected  int
	Items      []int
	Op         func(old int) int
	Test       func(item int) bool
	Next       map[bool]*Monkey
	LessStress bool
}

func (m *Monkey) turn(stress bool) {
	for _, item := range m.Items {
		item = m.Op(item)
		if !stress {
			item /= 3
		}
		test := m.Test(item)
		m.Inspected++
		m.Next[test].Items = append(m.Next[test].Items, item)
		m.Items = m.Items[1:]
	}

}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func attrVal(attr string) string {
	return strings.Split(attr, ": ")[1]
}

var Ops = map[string]func(i, j int) int{
	"+": func(i, j int) int { return i + j },
	"*": func(i, j int) int { return i * j },
}

func getLast(line string) int {
	arr := strings.Split(line, " ")
	val, err := strconv.Atoi(arr[len(arr)-1])
	check(err)
	return val
}

func getMonkeys(input string) []Monkey {
	sections := strings.Split(input, "\n\n")

	monkeys := make([]Monkey, len(sections))

	for i, section := range sections {
		monkey := &monkeys[i]
		monkeyAttrs := strings.Split(section, "\n")

		items := strings.Split(attrVal(monkeyAttrs[1]), ", ")
		for _, item := range items {
			iitem, err := strconv.Atoi(item)
			check(err)
			monkey.Items = append(monkey.Items, iitem)
		}

		ops := strings.Split(strings.Split(attrVal(monkeyAttrs[2]), "= ")[1], " ")
		if ops[2] == "old" {
			monkey.Op = func(item int) int { return Ops[ops[1]](item, item) }
		} else {
			val, err := strconv.Atoi(ops[2])
			check(err)
			monkey.Op = func(item int) int { return Ops[ops[1]](item, val) }
		}

		testVal := getLast(monkeyAttrs[3])
		monkey.Test = func(item int) bool { return item%testVal == 0 }

		monkey.Next = map[bool]*Monkey{}
		monkey.Next[true] = &monkeys[getLast(monkeyAttrs[4])]
		monkey.Next[false] = &monkeys[getLast(monkeyAttrs[5])]
	}

	return monkeys
}

func monkeyBusiness(input string, rounds int, stress bool) int {
	monkeys := getMonkeys(input)

	for round := 0; round < rounds; round++ {
		for i, _ := range monkeys {
			monkeys[i].turn(stress)
		}
	}

	sort.Slice(monkeys, func(i, j int) bool { return monkeys[i].Inspected > monkeys[j].Inspected })

	return monkeys[0].Inspected * monkeys[1].Inspected

}

func main() {
	fb, err := os.ReadFile("input.txt")
	check(err)

	input := string(fb)

	// Part 1
    fmt.Printf("Part 1 monkey business: %d\n", monkeyBusiness(input, 20, false))

	// Part 2
    fmt.Printf("Part 2 monkey business: %d\n", monkeyBusiness(input, 10000, true))

}
