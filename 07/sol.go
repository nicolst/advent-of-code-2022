package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type File struct {
	Name string
	Size int
}

type Folder struct {
	Parent   *Folder
	Children map[string]*Folder
	Files    []*File
	size     *int
}

func (f *Folder) Size() int {
	if f.size != nil {
		return *f.size
	}

	size := 0

	for _, file := range f.Files {
		size += file.Size
	}

	for _, folder := range f.Children {
		size += folder.Size()
	}

	f.size = &size
	return size
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	dat, err := os.ReadFile("input.txt")
	check(err)

	lines := strings.Split(string(dat), "\n")

	root := Folder{Parent: nil, Children: map[string]*Folder{}, Files: []*File{}}

	allFolders := []*Folder{&root}

	currDir := &root

	for _, line := range lines {
		if line == "" {
			break
		}
		args := strings.Split(line, " ")
		if args[0] == "$" {
			if args[1] == "cd" {
				if args[2] == ".." {
					currDir = currDir.Parent
				} else if args[2] == "/" {
					currDir = &root
				} else {
					currDir = currDir.Children[args[2]]
				}
			}
		} else {
			if args[0] == "dir" {
				newDir := Folder{Parent: currDir, Children: map[string]*Folder{}, Files: []*File{}}
				currDir.Children[args[1]] = &newDir
				allFolders = append(allFolders, &newDir)
			} else {
				fileSize, err := strconv.Atoi(args[0])
				check(err)
				currDir.Files = append(currDir.Files, &File{Name: args[1], Size: fileSize})
			}
		}

	}

	sum := 0
	for _, folder := range allFolders {
		if folder.Size() <= 100000 {
			sum += folder.Size()
		}

	}
	fmt.Printf("Total size for all folders<100000: %d\n", sum)

	sort.Slice(allFolders, func(i, j int) bool {
		return allFolders[i].Size() < allFolders[j].Size()
	})

	requiredSpace := 30000000 - (70000000 - root.Size())
	for _, folder := range allFolders {
		if folder.Size() >= requiredSpace {
			fmt.Printf("Smallest folder freeing up enough space: %d\n", folder.Size())
			break
		}
	}

}
