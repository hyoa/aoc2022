package day

import (
	"fmt"
	"hyoa/aoc2022/internal/utils"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Day13 struct {
	nodes []NodeSignal
	lines []string
}

type NodeSignal struct {
	childs []NodeSignal
	value  int
	data   string
}

func (d *Day13) Init(path string) error {
	nodes := make([]NodeSignal, 0)
	inputs := utils.ReadTextFileLinesAsString(path)

	for _, input := range inputs {
		if len(input) == 0 {
			continue
		}

		nodes = append(nodes, buildNode(input))
		d.lines = append(d.lines, input)
	}

	d.nodes = nodes

	return nil
}

func buildNode(input string) NodeSignal {
	// rootInput := input
	reg := regexp.MustCompile(`(?m)(?P<input>\[[\d\w,]*\])`)
	dictNodeReplaced := make(map[string]NodeSignal)
	itReplacement := 1

	for {
		sub := reg.FindAllString(input, -1)

		if len(sub) == 0 {
			toReturn := dictNodeReplaced[fmt.Sprintf("rep%d", itReplacement-1)]
			return toReturn
		}

		for _, s := range sub {
			name := fmt.Sprintf("rep%d", itReplacement)
			input = strings.Replace(input, s, name, -1)
			node := NodeSignal{value: -1}

			if s != "[]" {
				s = strings.ReplaceAll(s, "[", "")
				s = strings.ReplaceAll(s, "]", "")

				numbers := strings.Split(s, ",")

				for _, n := range numbers {
					var nodeToAppend NodeSignal
					if strings.Contains(n, "rep") {
						nodeToAppend = dictNodeReplaced[n]
					} else {
						nInt, _ := strconv.Atoi(n)
						nodeToAppend = NodeSignal{value: nInt}
					}

					node.childs = append(node.childs, nodeToAppend)
				}
			}

			dictNodeReplaced[name] = node
			itReplacement++
		}
	}
}

func (d *Day13) Step1() (Result, error) {
	pairResult := make([]bool, 0)
	for i := 0; i < len(d.nodes); i += 2 {
		v, _ := compareNodes(d.nodes[i], d.nodes[i+1])
		pairResult = append(pairResult, v)
	}

	tt := 0
	for k, p := range pairResult {
		if p {
			tt += k + 1
		}
	}

	return Result{Value: tt}, nil
}

func compareNodes(left, right NodeSignal) (bool, string) {
	idxL := len(left.childs)
	idxR := len(right.childs)

	var idx int
	if idxL > idxR {
		idx = idxL
	} else {
		idx = idxR
	}

	for k := 0; k < idx; k++ {
		if len(right.childs) <= k {
			return false, "right run out"
		} else if len(left.childs) <= k {
			return true, "left run out break"
		}

		leftChild := left.childs[k]
		rightChild := right.childs[k]

		if leftChild.value != -1 && rightChild.value != -1 {
			if leftChild.value > rightChild.value {
				return false, "right side is smaller"
			} else if leftChild.value < rightChild.value {
				return true, "left side is smaller break"
			}
		} else {
			if leftChild.value != -1 {
				leftChild = NodeSignal{childs: []NodeSignal{leftChild}}
			} else if rightChild.value != -1 {
				rightChild = NodeSignal{childs: []NodeSignal{rightChild}}
			}

			ok, msg := compareNodes(leftChild, rightChild)
			if !ok {
				return false, msg
			} else if ok && strings.Contains(msg, "break") {
				return true, msg
			}
		}
	}

	return true, "left run out"
}

type inputSignal struct {
	vS string
	vI int
}

func (d *Day13) Step2() (Result, error) {
	reg := regexp.MustCompile(`(?m)(?P<input>[\d]+|\]+)`)
	lines := d.lines
	lines = append(lines, "[[2]]")
	lines = append(lines, "[[6]]")

	inputsSignal := make([]inputSignal, 0)
	for _, l := range lines {
		var value int
		sub := reg.FindAllString(l, -1)
		if len(sub) == 0 || sub[0] == "]" {
			value = 0
		} else {
			v, _ := strconv.Atoi(sub[0])
			value = v
		}

		inputsSignal = append(inputsSignal, inputSignal{
			vS: l,
			vI: value,
		})
	}

	sort.Slice(inputsSignal, func(i, j int) bool {
		if inputsSignal[i].vI < inputsSignal[j].vI {
			return true
		}

		if inputsSignal[i].vI == inputsSignal[j].vI && len(inputsSignal[i].vS) < len(inputsSignal[j].vS) {
			return true
		}

		return false
	})

	tt := 1
	for k, i := range inputsSignal {
		fmt.Println("----")
		fmt.Println(i.vS)

		if i.vS == "[[2]]" || i.vS == "[[6]]" {
			tt *= k + 1
		}
	}

	return Result{Value: tt}, nil
}
