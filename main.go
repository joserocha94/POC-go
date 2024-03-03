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
		
	// graph
	test := structs.Graph {}
	
	// add nodes to graph
	test.AddNode(new_nodeA)
	test.AddNode(new_nodeB)
	test.AddNode(new_nodeC)

	// queue
	test_queue := structs.Queue {}
	//new_queuenodeA := structs.QueueNode { Base : new_nodeA }
	//new_queuenodeB := structs.QueueNode { Base : new_nodeB }
	//new_queuenodeC := structs.QueueNode { Base : new_nodeC }
	//new_queuenodeD := structs.QueueNode { Base : new_nodeD }

	Dijkstra(test, test_queue, new_nodeA)

	fmt.Println("\n\nFinish")
}

