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
	

	fmt.Println(test.Count())
	test.AddNode(new_nodeA)
	test.AddNode(new_nodeB)
	test.AddNode(new_nodeC)
	fmt.Println(test.Count())

	fmt.Println(new_nodeD.Id)
	
	fmt.Println(test.Contains(new_nodeA))
	fmt.Println(test.Contains(new_nodeD))

	fmt.Println("\nFinish")
}

