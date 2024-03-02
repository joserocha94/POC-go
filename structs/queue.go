package structs

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

func (q *Queue) Enqueue(new_node QueueNode) {
	
	if (q.IsEmpty()) {
		q.Head = &new_node
		q.Tail = &new_node
	} else {
		q.Tail.Next = &new_node
		q.Tail = &new_node
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
		return queueNode.Base.Id
	}
}

func (q *Queue) Print() string {
	
	output := ""

	if !(q.IsEmpty()) {
		curr := q.Head
		output += "[" + curr.Base.Id + "] "
		for (curr.Next != nil) {
			output += "[" + curr.Next.Base.Id + "] "
			curr = curr.Next
		}
		output += "\n\n"
	}
	
	return output
}