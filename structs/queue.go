package structs

type Item struct {
	Id string
	Next *Item
}

type Queue struct {
	Head *Item
	Tail *Item
}

func (q *Queue) IsEmpty() bool {

	if (q.Head == nil) {
		return true
	} else {
		return false
	}
}

func (q *Queue) Enqueue(Identifier string) {
	
	new_item := Item { Id : Identifier } 
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
		item := q.Head
		q.Head = q.Head.Next
		return item.Id
	}
}

func (q *Queue) Print() string {
	
	output := ""

	if !(q.IsEmpty()) {
		curr := q.Head
		output += "[" + curr.Id + "] "
		for (curr.Next != nil) {
			output += "[" + curr.Next.Id + "] "
			curr = curr.Next
		}
		output += "\n\n"
	}
	
	return output
}