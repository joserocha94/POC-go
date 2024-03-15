package procedures

import (
	"fmt"

	structs 	 "github.com/joserocha94/POC-go/structs"
	structs_ext  "github.com/joserocha94/POC-go/structs/extended"
)


var MAX_DISTANCE int = 999
var MIN_DISTANCE int = 0

var predecessors [] *structs_ext.Parent
var distances [] *Pair


type Pair struct {
    Id, Distance interface{}
}


func initDistancesQueue(g structs.Graph, q *structs.Queue, s structs_ext.Node){

	for _,v := range g.Nodes{

		distance := MIN_DISTANCE
		if v.Id == s.Id {
			distance = MIN_DISTANCE
		} else {
			distance = MAX_DISTANCE
		}
		new_queuenode := structs_ext.QueueNode { Base : *v, Distance : distance } 
		q.Enqueue(new_queuenode)
	}
}


// initialize predecessors
func initPredecessors(g structs.Graph){
	println(">>> Generating predecessors...")

	for _,v := range g.Nodes{
		new_parent := structs_ext.Parent{ v.Id, "-"}
		predecessors = append(predecessors, &new_parent)
	}

	println(">>> Generating predecessors done")
}


// initialize distance vector
func initDistances(q *structs.Queue){
	println(">>> Building distances vector...")

	if !(q.IsEmpty()) {
		curr := q.Head
		for (curr != nil) {
			new_pair := Pair{curr.Base.Id, curr.Distance}
			distances = append(distances, &new_pair)
			curr = curr.Next
		}
	}
	println(">>> Building distances vector done")
}


// initialize queue with max distances
// initialize vector predecessors
// initialize vector distances
func initialize(g structs.Graph, q *structs.Queue, s structs_ext.Node) {

	// queue
	initDistancesQueue(g, q, s)

	// predecessors
	initPredecessors(g)

	// distances vector
	initDistances(q)

}


// print predecessors
func printPredecessors(){
	for _,v := range predecessors {
		fmt.Println("(" + v.Node + ", " + v.Parent + ")" )
	}
}


// print distances
func printDistances(){
	for _,v := range distances {
		fmt.Println(v.Id, v.Distance)
	}
}


// print Queue 
// do to Queueu methods framework ..
func printQueue(q *structs.Queue){
	fmt.Println(q.Print())
}



func Dijkstra(g structs.Graph, q structs.Queue, s structs_ext.Node) {

	// initialize queue with max distances
	// initialize vector predecessors
	initialize(g, &q, s)

	println()
	println(">>> Setup builded, program ready")
	println()
	printPredecessors()
	printQueue(&q)

	// to do:

	//	- refactor initialize
	//	- dividir em 3 metodos:
	//		- predecessors
	//		- distances
	//		- encher a queue

	// 	- priority queue vs queue ?
	for i := range q.Count() {
		q.Dequeue()
		//printQueue(&q)
		i = i
	}

	println(">>> TO DO ..")
	println(">>> TO DO ..")

}