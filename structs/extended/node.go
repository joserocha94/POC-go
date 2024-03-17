package structs

type Edge struct {
	Node *Node
	Distance int
}

type Node struct {
	Id string
	Adjacent []*Edge
}


