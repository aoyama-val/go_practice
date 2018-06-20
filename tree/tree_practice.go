package main

import (
	"fmt"
)

type Tree struct {
	Value    string
	Children []Tree
}

func NewTree() *Tree {
	return &Tree{}
}

func (t *Tree) Print() {
	t.printSub(0, "", true)
}

func (t *Tree) printSub(level int, prefix string, isLastChild bool) {
	var corner string
	var appendPrefix string
	if level == 0 {
		appendPrefix = ""
		fmt.Printf("%s\n", t.Value)
	} else {
		if isLastChild {
			corner = "`"
			appendPrefix = "    "
		} else {
			corner = "|"
			appendPrefix = "|   "
		}
		fmt.Printf("%s%s---%s\n", prefix, corner, t.Value)
	}
	for i, child := range t.Children {
		if i == len(t.Children)-1 {
			child.printSub(level+1, prefix+appendPrefix, true)
		} else {
			child.printSub(level+1, prefix+appendPrefix, false)
		}
	}
}

func main() {
	tree := Tree{
		Value: "root",
	}
	tree.Children = append(
		tree.Children,
		Tree{
			Value: "sub1",
			Children: []Tree{
				Tree{
					Value: "sub11",
					Children: []Tree{
						Tree{
							Value: "sub111",
						},
					},
				},
				Tree{Value: "sub12"},
				Tree{Value: "sub13"},
			}})
	tree.Children = append(tree.Children, Tree{Value: "sub2", Children: []Tree{Tree{Value: "sub21"}}})
	tree.Children = append(tree.Children, Tree{Value: "sub3", Children: []Tree{Tree{Value: "sub31"}}})

	tree.Print()
}
