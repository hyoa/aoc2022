package day

import (
	"fmt"
	"hyoa/aoc2022/internal/utils"
	"strings"
)

type Day7 struct {
	inputs []string
}

type Node struct {
	size      int
	directory bool
	name      string
	children  map[string]*Node
	parent    *Node
	isRoot    bool
}

type Tree struct {
	root *Node
}

func (d *Day7) Init(path string) error {
	d.inputs = utils.ReadTextFileLinesAsString(path)
	return nil
}

func createTree(inputs []string) Tree {
	root := &Node{name: "/", directory: true, children: make(map[string]*Node), isRoot: true}
	tree := Tree{root: root}
	curr := root
	dirSize := make(map[string]int)

	for _, l := range inputs[1:] {
		if l == "$ cd .." {
			curr = curr.parent
		} else if strings.Contains(l, "$ cd") {
			var dir string
			fmt.Sscanf(l, "$ cd %s", &dir)

			curr = curr.children[dir]
		} else if l != "$ ls" {
			if strings.Contains(l, "dir") {
				var dir string
				fmt.Sscanf(l, "dir %s", &dir)
				curr.children[dir] = &Node{name: dir, parent: curr, directory: true, children: make(map[string]*Node)}
				dirSize[dir] = 0
			} else {
				var file string
				var size int
				fmt.Sscanf(l, "%d %s", &size, &file)
				nodeFile := &Node{name: file, parent: curr, size: size}
				curr.children[file] = nodeFile
				updateSizeFromFile(nodeFile, size, dirSize)
			}
		}
	}

	// fmt.Printf("%+v", tree.root)
	return tree
}

func updateSizeFromFile(node *Node, size int, dirSize map[string]int) {
	if node.isRoot {
		return
	}

	node.parent.size += size

	updateSizeFromFile(node.parent, size, dirSize)
}

func (d *Day7) Step1() (Result, error) {
	tree := createTree(d.inputs)
	size := walkForDirSizeLower100000(tree.root)

	return Result{Value: size}, nil
}

type sizeNode struct {
	v int
}

func (d *Day7) Step2() (Result, error) {
	tree := createTree(d.inputs)

	totalSize := 70000000
	spaceLeft := totalSize - tree.root.size
	sq := &sizeNode{v: 9999999999999999}

	walkForDirDeleteEnoughSpace(tree.root, spaceLeft, sq)

	return Result{Value: sq.v}, nil
}

func walkForDirSizeLower100000(curr *Node) int {
	if !curr.directory {
		return 0
	}

	totalSize := 0
	if curr.size <= 100000 {
		totalSize += curr.size
	}

	if curr.directory {
		for _, c := range curr.children {
			totalSize += walkForDirSizeLower100000(c)
		}
	}

	return totalSize
}

func walkForDirDeleteEnoughSpace(curr *Node, spaceLeft int, sq *sizeNode) bool {
	if !curr.directory || curr.size+spaceLeft < 30000000 {
		return false
	}

	for _, c := range curr.children {
		if walkForDirDeleteEnoughSpace(c, spaceLeft, sq) {
			if c.size < sq.v {
				sq.v = c.size
			}
		}
	}

	return true
}
