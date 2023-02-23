package deque

type DynamicSizedDeque struct {
	items []int
}

func EnqueueFront(d *DynamicSizedDeque, item int) {
	d.items = append([]int{item}, d.items...)
}

func EnqueueRear(d *DynamicSizedDeque, item int) {
	d.items = append(d.items, item)
}

func DequeueFront(d *DynamicSizedDeque) int {
	if len(d.items) == 0 {
		return -1
	}
	front := d.items[0]
	d.items = d.items[1:]
	return front
}

func DequeueRear(d *DynamicSizedDeque) int {
	if len(d.items) == 0 {
		return -1
	}
	rear := d.items[len(d.items)-1]
	d.items = d.items[:len(d.items)-1]
	return rear
}

func IsEmpty(d *DynamicSizedDeque) bool {
	return len(d.items) == 0
}

func Size(d *DynamicSizedDeque) int {
	return len(d.items)
}
