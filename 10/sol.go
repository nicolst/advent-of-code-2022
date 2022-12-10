package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Register []int

func addx(reg *Register, x int) {
    prev := (*reg)[len(*reg)-1]
    *reg = append(*reg, prev, prev+x)
}

func noop(reg *Register) {
    prev := (*reg)[len(*reg)-1]
    *reg = append(*reg, prev)
}

func check(err error) {
    if err != nil {
        panic(err)
    }
}

func Abs(i int) int {
    if i < 0 {
        return -i
    }
    return i
}

func main() {
    fb, err := os.ReadFile("input.txt")
    check(err)

    lines := strings.Split(string(fb), "\n")

    reg := make(Register, 1, len(lines))
    reg[0] = 1

    for _, line := range lines[:len(lines)-1] {
        cmd := strings.Split(line, " ")
        if cmd[0] == "noop" {
            noop(&reg)
        } else if cmd[0] == "addx" {
            x, err := strconv.Atoi(cmd[1])
            check(err)
            addx(&reg, x)
        }
    }

    tot := 0
    for i := 19; i < 220; i += 40 {
        tot += reg[i]*(i+1)
    }
    fmt.Printf("Total signal strength: %d\n", tot)

    fmt.Println("CRT screen")
    for row := 0; row < 6; row++ {
        for col := 0; col < 40; col++ {
            if Abs(col - reg[row * 40 + col]) <= 1 {
                fmt.Print("#")
            } else {
                fmt.Print(".")
            }
        }
        fmt.Println()
    }
}
