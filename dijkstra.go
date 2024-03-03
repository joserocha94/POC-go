package main

import (
	"fmt"

	structs "github.com/joserocha94/POC-go/structs"
)

var MAX_DISTANCE int = 999
var MIN_DISTANCE int = 0


func Dijkstra(g structs.Graph, q structs.Queue, s structs.Node) {

	g.Print()
	fmt.Println()

	// every node from [s] has distance = +00
	for _,v := range g.Nodes{

		distance := MIN_DISTANCE
		if v.Id == s.Id {
			distance = MIN_DISTANCE
		} else {
			distance = MAX_DISTANCE
		}
		new_queuenode := structs.QueueNode { Base : *v, Distance : distance } 
		q.Enqueue(new_queuenode)
	}

	fmt.Println(q.Print())

}