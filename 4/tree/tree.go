package tree

import (
	"fmt"
)

type Node struct {
	Value    int
	Children []*Node
}

type Tree struct {
	Root *Node
}

func (t *Tree) Insert(parentValue, value int) {
	if t.Root == nil {
		t.Root = &Node{Value: parentValue}
	}
	insertNode(t.Root, parentValue, value)
}

func insertNode(node *Node, parentValue, value int) {
	if node.Value == parentValue {
		node.Children = append(node.Children, &Node{Value: value})
		return
	}
	for _, child := range node.Children {
		insertNode(child, parentValue, value)
	}
}

func (t *Tree) Traverse(node *Node, depth int) {
	if node != nil {
		fmt.Println(string('|'+depth*'-'), node.Value)
		for _, child := range node.Children {
			t.Traverse(child, depth+1)
		}
	}
}
func main() {
	tree := &Tree{}
	tree.Insert(1, 2)
	tree.Insert(1, 3)
	tree.Insert(1, 4)
	tree.Insert(2, 5)
	tree.Insert(2, 6)
	tree.Insert(3, 7)
	tree.Insert(3, 8)

	fmt.Println("Tree Traversal:")
	tree.Traverse(tree.Root, 0)
}
