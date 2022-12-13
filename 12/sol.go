package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type Node struct {
	Height    int
	Visited   bool
	Distance  int
	Neighbors []*Node
}


type Map [][]*Node

func MakeMap(height, width int) Map {
	newMap := make(Map, height)
	for i := 0; i < height; i++ {
		newMap[i] = make([]*Node, width)
		for j := 0; j < width; j++ {
			newMap[i][j] = &Node{}
		}
	}
	return newMap
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fb, err := os.ReadFile("input.txt")
	check(err)

	rows := strings.Split(string(fb), "\n")

	height := len(rows) - 1
	width := len(rows[0])
    maxDistance := height*width

	nodeMap := MakeMap(height, width)
	unvisited := make([]*Node, height*width)

	var end *Node

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			nodeMap[i][j].Height = int(rows[i][j]) - int('a')
			nodeMap[i][j].Distance = maxDistance
			nodeMap[i][j].Neighbors = make([]*Node, 4)
			if i != 0 {
				nodeMap[i][j].Neighbors[0] = nodeMap[i-1][j]
			}
			if j != 0 {
				nodeMap[i][j].Neighbors[1] = nodeMap[i][j-1]
			}
			if i != height-1 {
				nodeMap[i][j].Neighbors[2] = nodeMap[i+1][j]
			}
			if j != width-1 {
				nodeMap[i][j].Neighbors[3] = nodeMap[i][j+1]
			}
			if rows[i][j] == 'S' {
				nodeMap[i][j].Height = 0
				nodeMap[i][j].Distance = 0
			} else if rows[i][j] == 'E' {
				nodeMap[i][j].Height = 25
				end = nodeMap[i][j]
			}
			unvisited[i*width+j] = nodeMap[i][j]
		}
	}

	for !end.Visited {
		sort.Slice(unvisited, func(i, j int) bool {
			return unvisited[i].Distance < unvisited[j].Distance
		})
		current := unvisited[0]

		for _, neighbor := range current.Neighbors {
			if neighbor == nil || neighbor.Visited || neighbor.Height > current.Height+1 {
				continue
			}
			if current.Distance+1 < neighbor.Distance {
				neighbor.Distance = current.Distance + 1
				if neighbor == end {
					end.Visited = true
					break
				}
			}
		}

		current.Visited = true
		unvisited = unvisited[1:]

	}

	fmt.Println("Shortest path: ", end.Distance)

}
