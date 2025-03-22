package main

//import "fmt"

//type Node struct {
//	Value int
//	Next  *Node
//}

//type LinkedList struct {
//	Head *Node
//}

//func main() {
//	list := LinkedList{}
//	list.Add(1)
//	list.Add(2)
//	list.Add(3)
//	list.Add(4)
//	list.Add(5)

//list.deleteMiddleNode(3)
//	var node = list.Head
//	for i := 0; i < 3; i++ {
//		node = node.Next
//	}
//	fmt.Println(node.Value)
//	list.deleteGivenNode(node)
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
//	for current != nil {
//		fmt.Printf("%d -> ", current.Value)
//		current = current.Next
//	}
//	fmt.Println("nil")
//}

func (list *LinkedList) deleteMiddleNode(val int) {
	currNode := list.Head
	var prevNode *Node
	for currNode != nil {
		if currNode.Value == val {
			prevNode.Next = currNode.Next
			currNode.Next = nil
		}
		prevNode = currNode
		currNode = currNode.Next
	}
}

func (list *LinkedList) deleteGivenNode(node *Node) {
	if node == nil || node.Next == nil {
		return
	}

	nextNode := node.Next
	node.Value = nextNode.Value
	node.Next = nextNode.Next
}
