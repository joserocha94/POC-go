package main

import (
	"fmt"

	structs "github.com/joserocha94/POC-go/structs"
)
	

func main() {

    fmt.Println("Starting...\n")

	test := &structs.Graph{}
	for i := 0; i < 3; i++ {
		test.AddNode(i*10);
	}

	
	test.Print()
	test.AddEdge(0,10)
	test.AddEdge(0,20)
	test.AddEdge(10,20)
	fmt.Println()

	test.Print()
	fmt.Println()

	queue := &structs.Queue{}
	queue.Id = 5

	fmt.Println("\n%d", structs.Contador)

	fmt.Println("\nFinish")
}

