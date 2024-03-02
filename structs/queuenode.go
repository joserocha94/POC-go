package structs

type QueueNode struct {
	Base Node
	Next *QueueNode
}