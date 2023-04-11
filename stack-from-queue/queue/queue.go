package queue

type Queue struct {
	items []int
}

func New() *Queue {
	return &Queue{
		items: make([]int, 0),
	}
}

func (d Queue) Values() []int {
	return d.items
}

// func (d *Queue) Enqueue(item int) {
// 	d.items = append([]int{item}, d.items...)
// }

func (d *Queue) Enqueue(item int) {
	d.items = append(d.items, item)
}

func (d *Queue) Dequeue() int {
	if len(d.items) == 0 {
		return -1
	}
	front := d.items[0]
	d.items = d.items[1:]
	return front
}

func (d *Queue) IsEmpty() bool {
	return len(d.items) == 0
}

func (d *Queue) Size() int {
	return len(d.items)
}
