package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func main() {
	//headNode := &Node{data: 1}
	//headNode.next = &Node{data: 2}
	//headNode.next.next = &Node{data: 3}
	//headNode.next.next.next = &Node{data: 4}
	//for headNode.next != nil {
	//	fmt.Println(headNode.data)
	//	headNode = headNode.next
	//}
	numbers := []int{1, 2, 3}
	linkedList := createLinkedList(numbers)
	for linkedList.next != nil {
		fmt.Println(linkedList.data)
		linkedList = *linkedList.next
	}
	fmt.Println(linkedList.data)

}

func createLinkedList(numbers []int) Node {
	headNode := &Node{}
	for i, n := range numbers {
		if i == 0 {
			headNode.data = n
		}
		if i != 0 {
			node := Node{}
			node.data = n
			node.next = headNode
			headNode = &node
		}
	}
	return *headNode
}
