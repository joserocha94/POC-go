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

	fmt.Println("\nQueue: \n")

	queue := &structs.Queue{}

	queue.Enqueue("Primeiro")
	queue.Enqueue("Segundo")
	queue.Enqueue("Terceiro")
	queue.Enqueue("Quarto")
	queue.Enqueue("Quinto")

	fmt.Print(queue.Print())

	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())

	fmt.Println("\nFinish")
}

