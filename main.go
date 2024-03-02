package main

import (
	"fmt"

	structs "github.com/joserocha94/POC-go/structs"
)
	

func main() {

    fmt.Println("Starting...\n")

	test := &structs.Graph{}

	test.AddNode("A");
	test.AddNode("B");
	test.AddNode("C");
	
	test.Print()
	test.AddEdge("A", "B")
	test.AddEdge("A", "C")
	test.AddEdge("B", "C")
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

