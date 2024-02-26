package main

import "fmt"

type Graph struct {
	nodes []*Node
}

type Node struct{
	id int
	adjacent []*Node
}

func (g *Graph) AddNode(k int){
	g.nodes = append(g.nodes, &Node{id : k})
}

func (g *Graph) Print(){
	for i,v := range g.nodes{
		fmt.Printf("Node [%d] : %d \n", i, v.id)
		for _,v := range v.adjacent{
			fmt.Printf("%v", v.id)
		}
	}
}

func main() {
    fmt.Println("Starting...\n")

	test := &Graph{}
	for i := 0; i < 3; i++ {
		test.AddNode(i*10);
	}

	test.Print()

	fmt.Println("\nFinish")
}
