package main

import (
	"fmt"

	procedures 	"github.com/joserocha94/POC-go/procedures"
	structs 	"github.com/joserocha94/POC-go/structs"
	structs_ext "github.com/joserocha94/POC-go/structs/extended"
)
	

func main() {

    fmt.Println("Starting...\n")


	// build nodes
	new_nodeA := structs_ext.Node { Id: "A", Adjacent : [] *structs_ext.Edge{} }
	new_nodeB := structs_ext.Node { Id: "B", Adjacent : [] *structs_ext.Edge{} }
	new_nodeC := structs_ext.Node { Id: "C", Adjacent : [] *structs_ext.Edge{} }

		
	// build edges
	new_edgeAB := structs_ext.Edge { Node: &new_nodeB, Distance : 5}
	new_edgeAC := structs_ext.Edge { Node: &new_nodeC, Distance : 7}


	// build adjacency list
	new_nodeA.Adjacent = append(new_nodeA.Adjacent, &new_edgeAB)
	new_nodeA.Adjacent = append(new_nodeA.Adjacent, &new_edgeAC)


	// queue
	test_queue := structs.Queue {}


	// graph
	test := structs.Graph {}

	
	// add nodes to graph
	test.AddNode(new_nodeA)
	test.AddNode(new_nodeB)
	test.AddNode(new_nodeC)


	// run dijkstra
	procedures.Dijkstra(test, test_queue, new_nodeA)

	fmt.Println("\nFinish")
}

