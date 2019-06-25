package queue

var q []string

//Push add new element to the queue.
func Push(str string) {
	q = append(q, str)
}

//Pop returns first element added to the queue.
func Pop() string {
	lq := len(q)
	if lq == 0 {
		return ""
	}

	popElem := q[0]
	q = q[1:lq]
	return popElem
}
