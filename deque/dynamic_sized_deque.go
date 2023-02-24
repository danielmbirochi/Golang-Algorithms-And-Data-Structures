package deque

type DynamicSizedDeque struct {
	items []int
}

func (d DynamicSizedDeque) Values() []int {
	return d.items
}

func (d *DynamicSizedDeque) EnqueueFront(item int) {
	d.items = append([]int{item}, d.items...)
}

func (d *DynamicSizedDeque) EnqueueBack(item int) {
	d.items = append(d.items, item)
}

func (d *DynamicSizedDeque) DequeueFront() int {
	if len(d.items) == 0 {
		return -1
	}
	front := d.items[0]
	d.items = d.items[1:]
	return front
}

func (d *DynamicSizedDeque) DequeueBack() int {
	if len(d.items) == 0 {
		return -1
	}
	rear := d.items[len(d.items)-1]
	d.items = d.items[:len(d.items)-1]
	return rear
}

func (d *DynamicSizedDeque) IsEmpty() bool {
	return len(d.items) == 0
}

func (d *DynamicSizedDeque) Size() int {
	return len(d.items)
}
