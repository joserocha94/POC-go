package main

import (
	"fmt"

	procedures 	"github.com/joserocha94/POC-go/procedures"
	structs 	"github.com/joserocha94/POC-go/structs"
	structs_ext "github.com/joserocha94/POC-go/structs/extended"
)
	

func main() {

    fmt.Println("Starting...\n")

	// nodes
	new_nodeA := structs_ext.Node { Id: "A", Adjacent : [] *structs_ext.Node{} }
	new_nodeB := structs_ext.Node { Id: "B", Adjacent : [] *structs_ext.Node{} }
	new_nodeC := structs_ext.Node { Id: "C", Adjacent : [] *structs_ext.Node{} }
		
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

	procedures.Dijkstra(test, test_queue, new_nodeA)

	fmt.Println("\nFinish")
}

