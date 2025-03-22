package main

//type Node struct {
//	Value int
//	Next  *Node
//}

//type LinkedList struct {
//	Head *Node
//}

//func main() {
//	list := LinkedList{}
//	list.Add(3)
//	list.Add(5)
//	list.Add(1)
//	list.Add(8)
//	list.Add(5)
//	list.Add(10)
//	list.Add(2)
//	list.Add(4)
//	list.Add(10)
//list.partition(5)
//list.partitionScndTry(5)
//	list.partitionFromBook(5)
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
//
//	func (list *LinkedList) Display() {
//		current := list.Head
//		for current != nil {
//			fmt.Printf("%d -> ", current.Value)
//			current = current.Next
//		}
//		fmt.Println("nil")
//	}/

func (list *LinkedList) partitionFromBook(val int) {
	var beforeStart *Node
	var beforeEnd *Node
	var afterStart *Node
	var afterEnd *Node

	currNode := list.Head

	for currNode != nil {
		nextNode := currNode.Next
		currNode.Next = nil

		if currNode.Value < val {
			if beforeStart == nil {
				beforeStart = currNode
				beforeEnd = beforeStart
			} else {
				beforeEnd.Next = currNode
				beforeEnd = currNode
			}
		}

		if currNode.Value >= val {
			if afterStart == nil {
				afterStart = currNode
				afterEnd = afterStart
			} else {
				afterEnd.Next = currNode
				afterEnd = currNode
			}
		}
		currNode = nextNode

	}
	// Merge the two lists
	if beforeStart == nil {
		// If there's no "before" list, the head is the start of the "after" list
		list.Head = afterStart
		return
	}

	beforeEnd.Next = afterStart
	list.Head = beforeStart
}
