package main

import (
	"fmt"
	"math"
)

type TailAndSize struct {
	Tail *Node
	Size int
}

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func (list *LinkedList) Add(value int) {
	if list.Head == nil {
		list.Head = &Node{Value: value}
		return
	}
	current := list.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = &Node{Value: value}
}

func (list *LinkedList) Display() {
	current := list.Head
	for current != nil {
		fmt.Printf("%d -> ", current.Value)
		current = current.Next
	}
	fmt.Println("nil")
}

func main() {
	common := &Node{Value: 7, Next: &Node{Value: 8, Next: &Node{Value: 9}}}
	headA := &Node{Value: 1, Next: &Node{Value: 2, Next: &Node{Value: 3, Next: common}}}
	headB := &Node{Value: 4, Next: &Node{Value: 5, Next: common}}

	intersection := intersectionNodes(headA, headB)

	// Displaying the result
	if intersection != nil {
		fmt.Printf("Intersection at node with value: %d\n", intersection.Value)
	} else {
		fmt.Println("No intersection")
	}

}

func intersectionNodes(headA *Node, headB *Node) *Node {
	if headA == nil || headB == nil {
		return nil
	}

	res1 := getTailAndSize(headA)
	res2 := getTailAndSize(headB)

	if res1.Tail != res2.Tail {
		return nil
	}

	var longer, shorter *Node
	if res1.Size < res2.Size {
		shorter = headA
		longer = headB
	} else {
		shorter = headB
		longer = headA
	}

	longer = getKthNode(longer, int(math.Abs(float64(res1.Size)-float64(res2.Size))))

	for shorter != longer {
		shorter = shorter.Next
		longer = longer.Next
	}

	return longer
}

func getTailAndSize(list *Node) TailAndSize {
	if list == nil {
		return TailAndSize{}

	}

	size := 1
	currNode := list
	for currNode.Next != nil {
		size++
		currNode = currNode.Next
	}

	res := TailAndSize{
		Size: size,
		Tail: currNode,
	}
	return res
}

func getKthNode(head *Node, k int) *Node {
	currNode := head
	for k > 0 && currNode != nil {
		currNode = currNode.Next
		k--
	}

	return currNode
}
