package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Direction int

type Pos struct {
	x int
	y int
}

var (
	Cardinals map[string]Pos = map[string]Pos{
		"U": Pos{0, 1},
		"D": Pos{0, -1},
		"R": Pos{1, 0},
		"L": Pos{-1, 0},
	}
)

func (pos *Pos) Add(move Pos) {
	pos.x += move.x
	pos.y += move.y
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func Min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func Sgn(i int) int {
    if i < 0 {
        return -1
    } 
    return 1
}


func (p1 Pos) GenMove(p2 Pos) Pos {
	xmove := p2.x - p1.x
	ymove := p2.y - p1.y

	if Abs(xmove)+Abs(ymove) <= 1 {
		return Pos{}
	}

	if Abs(xmove) == 1 && Abs(ymove) == 1 {
		return Pos{}
	}

    if xmove != 0 {
        xmove /= Abs(xmove)
    }
    if ymove != 0 {
        ymove /= Abs(ymove)
    }

	return Pos{xmove, ymove}
}

type RopeEnd struct {
	Pos
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func Simulate(knots int) int {
	f, err := os.ReadFile("input.txt")
	check(err)

    lines := strings.Split(string(f), "\n")

    ropeKnots := make([]Pos, knots)

	tailPositions := map[int]map[int]bool{}
    visited := 0

    head := &ropeKnots[0]
    tail := &ropeKnots[knots - 1]
    for _, line := range lines[:len(lines)-1] {
		move := strings.Split(line, " ")
		cardinal := Cardinals[move[0]]
		count, err := strconv.Atoi(move[1])
		check(err)
		for i := 0; i < count; i++ {
			head.Add(cardinal)
            for i := 1; i < knots; i++{
                m := ropeKnots[i].GenMove(ropeKnots[i-1])
                ropeKnots[i].Add(m)
            }
			if _, ok := tailPositions[tail.x]; !ok {
				tailPositions[tail.x] = map[int]bool{}
			}
            if _, ok := tailPositions[tail.x][tail.y]; !ok {
                visited++
            }
			tailPositions[tail.x][tail.y] = true
		}
	}
    return visited
}

func main() {
	fmt.Printf("Part 1: The tail visited %d places\n", Simulate(2))
	fmt.Printf("Part 2: The tail visited %d places", Simulate(10))
}
