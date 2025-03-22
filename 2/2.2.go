package main

//import "fmt"

//type Node struct {
//	Value int
//	Next  *Node
//}/

//type LinkedList struct {
//	Head *Node
//}

//func main() {
//	list := LinkedList{}
//	list.Add(1)
//	list.Add(2)
///	list.Add(3)
//	list.Add(4)
//	list.Add(5)
//	list.Add(6)
//	list.Add(7)
//	list.Add(8)
//	list.Add(9)
//	list.Add(10)
//	list.Add(11)
//	list.Add(12)
//	list.Add(13)
//	list.Add(14)
//	list.Add(15)
//
//	nodeVal := list.KtoLast(6)
//
//	fmt.Print(nodeVal)
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
//	for current != nil {
///		fmt.Printf("%d -> ", current.Value)
//		current = current.Next
//	}
//	fmt.Println("nil")
//}

//func (list *LinkedList) KtoLast(k int) int {
//I WILL REVERSE THE LINKED LIST AND RETURN THE K NODE
//	currNode := list.Head
///	var prevNode *Node
//	for currNode != nil {
//
///		nextNode := currNode.Next
//
///		currNode.Next = prevNode
//
//		prevNode = currNode
//
//		currNode = nextNode
//	}

//	list.Head = prevNode

//	list.Display()
//	node := list.Head
//	for i := 0; i < k; i++ {
//		node = node.Next
//	}
//	return node.Value

//}
