package main

import (
	"fmt"

	structs "github.com/joserocha94/POC-go/structs"
)
	

func main() {

    fmt.Println("Starting...\n")
	// nodes
	new_nodeA := structs.Node { Id: "A", Adjacent : [] *structs.Node{} }
	new_nodeB := structs.Node { Id: "B", Adjacent : [] *structs.Node{} }
	new_nodeC := structs.Node { Id: "C", Adjacent : [] *structs.Node{} }
	new_nodeD := structs.Node { Id: "D", Adjacent : [] *structs.Node{} }
	
	// graph
	test := &structs.Graph {}
	
	// add nodes to graph
	fmt.Println(test.Count())
	test.AddNode(new_nodeA)
	test.AddNode(new_nodeB)
	test.AddNode(new_nodeC)
	fmt.Println(test.Count())

	// check functions
	fmt.Println(test.Contains(new_nodeA))
	fmt.Println(test.Contains(new_nodeD))

	// queue
	test_queue := structs.Queue {}
	new_queuenodeA := structs.QueueNode { Base : new_nodeA }
	new_queuenodeB := structs.QueueNode { Base : new_nodeB }
	new_queuenodeC := structs.QueueNode { Base : new_nodeC }
	new_queuenodeD := structs.QueueNode { Base : new_nodeD }

	// check queue functions
	test_queue.Enqueue(new_queuenodeA)
	test_queue.Enqueue(new_queuenodeB)
	test_queue.Enqueue(new_queuenodeC)
	fmt.Println(test_queue.Count())
	test_queue.Dequeue()
	fmt.Println(test_queue.Count())	
	fmt.Println(test_queue.Print())

	test_queue.Enqueue(new_queuenodeD)
	fmt.Println(test_queue.Count())	
	fmt.Println(test_queue.Print())

	fmt.Println("Finish")
}

