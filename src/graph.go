package main

import "fmt"

type Graph struct {
	nodes []*Node
}

// adds node [k] to graph [g]
func (g *Graph) AddNode(k int) {
	g.nodes = append(g.nodes, &Node{id : k})
}

// indicates if a given node [id] is a node in the graph [g]
func (g *Graph) Contains(id int) bool {

	for _,v := range g.nodes{
		if v.id == id {
			return true
		} 
	}
	return false
}

// print graph [g] nodes
// to do : adjacency list
func (g *Graph) Print() {

	for i,v := range g.nodes {
		fmt.Printf("\nNode [%d] : %d", i, v.id)
		for _,v := range v.adjacent {
			fmt.Printf("\n\t-> %v", v.id)
		}
	}
}

// adds edge to the graph 
// between nodes [source] -> [target]
func (g *Graph) AddEdge(source int, target int) {
	
	//check if the nodes exists
	if g.Contains(source) && g.Contains(target) {
		
		source_node := g.GetNode(source)
		target_node := g.GetNode(target)
		source_node.adjacent = append(source_node.adjacent, target_node)

	} else {
		fmt.Println("Invalid Nodes")
	}
}

// Get node [id] from graph [g]
func (g *Graph) GetNode(id int) *Node {

	for i,v := range g.nodes {
		if v.id == id {
			return g.nodes[i]
		}
	}	
	return nil
}