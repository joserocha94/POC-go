package main

import (
	"fmt"

	structs 	"github.com/joserocha94/POC-go/structs"
	structs_ext "github.com/joserocha94/POC-go/structs/extended"

)

var MAX_DISTANCE int = 999
var MIN_DISTANCE int = 0

var predecessors [] *structs_ext.Parent


// initialize queue with max distances
// initialize vector predecessors
func initialize(g structs.Graph, q *structs.Queue, s structs.Node) {
	for _,v := range g.Nodes{

		//distances
		distance := MIN_DISTANCE
		if v.Id == s.Id {
			distance = MIN_DISTANCE
		} else {
			distance = MAX_DISTANCE
		}
		new_queuenode := structs.QueueNode { Base : *v, Distance : distance } 
		q.Enqueue(new_queuenode)

		//predecessors
		new_parent := structs_ext.Parent{ v.Id, "-"}
		predecessors = append(predecessors, &new_parent)
	}
}


func Dijkstra(g structs.Graph, q structs.Queue, s structs.Node) {


	// initialize queue with max distances
	// initialize vector predecessors
	initialize(g, &q, s)

	fmt.Println()
	for _,v := range predecessors{
		fmt.Println("(" + v.Node + ", " + v.Parent + ")" )
	}

	fmt.Println(q.Print())

}