package structs

import "fmt"

type Graph struct {
	Nodes []*Node
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
func (g *Graph) AddNode(new_node Node){
	g.Nodes = append(g.Nodes, &new_node)
}

// indicates if a given node [id] is a node in the graph [g]
func (g *Graph) Contains(node Node) bool {

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
			fmt.Printf("\n\t-> %v", v.Id)
		}
	}
}

// adds edge to the graph 
// between Nodes [source] -> [target]
func (g *Graph) AddEdge(source Node, target Node) {
	
	//check if the Nodes exists
	if g.Contains(source) && g.Contains(target) {
		source.Adjacent = append(source.Adjacent, &target)
	} else {
		fmt.Println("Invalid Nodes")
	}

}

