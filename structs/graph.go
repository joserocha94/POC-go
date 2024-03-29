package structs

import (
	"fmt"
	structs_ext "github.com/joserocha94/POC-go/structs/extended"
)

type Graph struct {
	Nodes []*structs_ext.Node
}

// adds node [k] to graph [g]
/*
func (g *Graph) AddNode(k string) {
	g.Nodes = append(g.Nodes, &Node{Id : k})
}
*/

// indicates the graph number of elements
func (g *Graph) Count() int {
	return len(g.Nodes)
}

// adds node [new_node] to graph [g]
func (g *Graph) AddNode(new_node structs_ext.Node){
	g.Nodes = append(g.Nodes, &new_node)
}

// indicates if a given node [id] is a node in the graph [g]
func (g *Graph) Contains(node structs_ext.Node) bool {

	for _,v := range g.Nodes{
		if v.Id == node.Id {
			return true
		} 
	}
	return false
}

// print graph [g] Nodes
// to do : adjacency list
func (g *Graph) Print() {

	for i,v := range g.Nodes {
		fmt.Printf("\nNode [%d] : %s", i, v.Id)
		for _,v := range v.Adjacent {
			fmt.Printf("\n\t-> %v", v.Node.Id)
		}
	}
}

// adds edge to the graph 
// between Nodes [source] -> [target]
// to do
func (g *Graph) AddEdge(source structs_ext.Node, target structs_ext.Node) {
	
	//check if the Nodes exists
	if g.Contains(source) && g.Contains(target) {
		//source.Adjacent = append(source.Adjacent, &target)
	} else {
		fmt.Println("Invalid Nodes")
	}

}

