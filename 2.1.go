package main

//
//import "fmt"

// Node represents a node in the linked list
//type Node struct {
//	Value int
//	Next  *Node
//}

// LinkedList represents a linked list
//type LinkedList struct {
//	Head *Node
//}/

//func main() {
//	list := LinkedList{}
//	list.Add(3)
//	list.Add(4)
//	list.Add(1)
//	list.Add(4)
//	list.Add(4)
///	list.Add(5)
//	list.Add(3)
//	list.Add(4)
////	list.Add(8)
//	list.Add(6)
//	list.Add(6)
//	list.Add(6)
//	list.Add(2)
//	list.Add(3)
//
//	list.removeDublicates()

//	list.Display()
//}

// Add appends a new node to the end of the linked list
//func (list *LinkedList) Add(value int) {
//	if list.Head == nil {
//		list.Head = &Node{Value: value}
//		return
//	}
//	current := list.Head
//	for current.Next != nil {
//		current = current.Next
//	}
//	current.Next = &Node{Value: value}
//}

// Display prints the values of the linked list
//func (list *LinkedList) Display() {
//	current := list.Head
///	for current != nil {
//		fmt.Printf("%d -> ", current.Value)
//		current = current.Next
//	}
//	fmt.Println("nil")
//}

// for singly linked list
//func (list *LinkedList) removeDublicates() {
//	nodes := make(map[int]bool)
//	currNode := list.Head
//	var prevNode *Node
//	for currNode != nil {
//		currVal := currNode.Value
//		copyCurrNode := currNode
///		_, ok := nodes[currVal]
//		if !ok {
//			nodes[currVal] = true
//
//			prevNode = copyCurrNode
//		} else {
//
//			//REMOVE THE NODE FROM THE LIST
//			prevNode.Next = currNode.Next
///			//`currNode.Next = nil

//		}

//		currNode = copyCurrNode.Next
//	}
//}
