package structs

type QueueNode struct {
	Base Node
	Next *QueueNode
}

type Queue struct {
	Head *QueueNode
	Tail *QueueNode
}

func (q *Queue) IsEmpty() bool {

	if (q.Head == nil) {
		return true
	} else {
		return false
	}
}

func (q *Queue) Enqueue(Identifier string) {
	
	new_item := QueueNode { 
		Base : Node {id : Identifier} }
		 
	if (q.IsEmpty()) {
		q.Head = &new_item
		q.Tail = &new_item
	} else {
		q.Tail.Next = &new_item
		q.Tail = &new_item
	}
}

func (q *Queue) Dequeue() string {

	if q.IsEmpty() {
		q.Tail.Next = nil
		q.Tail = nil
		return ""
	} else {
		queueNode := q.Head
		q.Head = q.Head.Next
		return queueNode.Base.id
	}
}

func (q *Queue) Print() string {
	
	output := ""

	if !(q.IsEmpty()) {
		curr := q.Head
		output += "[" + curr.Base.id + "] "
		for (curr.Next != nil) {
			output += "[" + curr.Next.Base.id + "] "
			curr = curr.Next
		}
		output += "\n\n"
	}
	
	return output
}