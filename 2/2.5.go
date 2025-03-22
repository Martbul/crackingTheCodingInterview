package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

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
	list1 := LinkedList{}
	list1.Add(7)
	list1.Add(1)
	list1.Add(6)

	list2 := LinkedList{}
	list2.Add(5)
	list2.Add(9)
	list2.Add(2)

	sum := sumLists(list1.Head, list2.Head)

	fmt.Print(sum)
}

func sumLists(headA *Node, headB *Node) int {
	curr1 := reverseList(headA)
	curr2 := reverseList(headB)

	var strBuilder1 strings.Builder
	var strBuilder2 strings.Builder
	var b1 bytes.Buffer
	for curr1 != nil {
		fmt.Println(curr1.Value)
		b1.WriteString(strconv.Itoa(curr1.Value))
		curr1 = curr1.Next
	}

	fmt.Println(b1.String())
	for curr2 != nil {
		strBuilder2.WriteString(strconv.Itoa(curr2.Value))
		curr2 = curr2.Next
	}

	n1, _ := strconv.Atoi(strBuilder1.String())
	n2, _ := strconv.Atoi(strBuilder2.String())
	fmt.Println(n1)
	sumValues := n1 + n2

	sumString := strconv.Itoa(sumValues)
	fmt.Println(sumString)

	for i := len(sumString); i > 0; i-- {
		s := string(sumString[i])
		fmt.Println(s)
	}
	return 1
}

func reverseList(list *Node) *Node {
	currNode := list
	var prevNode *Node
	var newHead *Node

	for currNode != nil {
		nextNode := currNode.Next
		if prevNode != nil {
			currNode.Next = prevNode
		}
		newHead = currNode
		currNode = nextNode
	}

	return newHead

}
