package structs

type QueueNode struct {
	Base Node
	Distance int
	Next *QueueNode
}