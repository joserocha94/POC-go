package main

import "fmt"
	

func main() {

    fmt.Println("Starting...\n")

	test := &Graph{}
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

	queue := &Queue{}
	queue.id = 5

	faker := &xpto{}

	fmt.Printf("\n %d", queue.id)

	fmt.Println("\nFinish")
}
