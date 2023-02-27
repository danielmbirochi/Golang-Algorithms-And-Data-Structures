package deque

type FixedSizeDeque struct {
	items [5]int
	front int
	back  int
}

func (d *FixedSizeDeque) EnqueueFront(item int) bool {
	if d.IsFull() {
		return false
	}
	d.front = (d.front - 1 + len(d.items)) % len(d.items)
	d.items[d.front] = item
	return true
}

func (d *FixedSizeDeque) EnqueueBack(item int) bool {
	if d.IsFull() {
		return false
	}
	d.items[d.back] = item
	d.back = (d.back + 1) % len(d.items)
	return true
}

func (d *FixedSizeDeque) DequeueFront() (int, bool) {
	if d.IsEmpty() {
		return 0, false
	}
	front := d.items[d.front]
	d.front = (d.front + 1) % len(d.items)
	return front, true
}

func (d *FixedSizeDeque) DequeueBack() (int, bool) {
	if d.IsEmpty() {
		return 0, false
	}
	d.back = (d.back - 1 + len(d.items)) % len(d.items)
	rear := d.items[d.back]
	return rear, true
}

func (d *FixedSizeDeque) IsEmpty() bool {
	return d.front == d.back
}

func (d *FixedSizeDeque) IsFull() bool {
	return (d.back+1)%len(d.items) == d.front
}

func (d *FixedSizeDeque) Size() int {
	if d.IsEmpty() {
		return 0
	}
	return (d.back - d.front + len(d.items)) % len(d.items)
}
