package fourTwo

import (
	"github.com/martbul/4/tree"
)

// 1 4 5 7 9 11 33 56 57 61 69 70 75 80 100
func MinimalTree(sortedArr []int) {
	treeDepth := len(sortedArr) / 2
	count := 1
	rootVal := sortedArr[treeDepth]
	root := &tree.Node{Value: rootVal}
	for treeDepth > 0 {
		c1 := &tree.Node{Value: sortedArr[treeDepth-count]}
		c2 := &tree.Node{Value: sortedArr[treeDepth+count]}
		root.Children = append(root.Children, c1, c2)
		count++
		treeDepth--
	}
}

func printTree(root *tree.Node) {

}
