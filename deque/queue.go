package deque

type Queue []int

func (q *Queue) Enqueue(v int) {
	*q = append(*q, v)
}

func (q *Queue) Dequeue() (int, bool) {
	if len(*q) == 0 {
		return 0, false
	}

	// Without parentheses, Go first applies the slices operation `[1:]` to `q`, and then tries to
	// deference the result with `*`. This doesn't make sense because `q` is a pointer to a slice,
	// not the slice itself, so I can't apply the slice operation to `q`.
	// By adding the parentheses, i'm changing the order of operations. Go now deferences `q` with `*` to get
	// the slice that `q` points to, and then applies the slice operation to that slice.
	value := (*q)[0]
	*q = (*q)[1:]

	return value, true
}

func (q *Queue) Peak() int {
	return (*q)[0]
}
