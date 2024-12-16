package problem_0707

type MyLinkedList struct {
	head *Node
	size int
}

type Node struct {
	value int
	next  *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}

	c := this.head
	for i := 0; i < index; i++ {
		c = c.next
	}

	return c.value
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.size, val)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.size {
		return
	}

	this.size++

	if index == 0 {
		newHead := &Node{value: val, next: this.head}
		this.head = newHead
		return
	}

	c := this.head
	for i := 0; i < index-1; i++ {
		c = c.next
	}

	oldNext := c.next
	c.next = &Node{value: val, next: oldNext}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		return
	}

	this.size--
	if index == 0 {
		this.head = this.head.next
		return
	}

	c := this.head
	for i := 0; i < index-1; i++ {
		c = c.next
	}
	c.next = c.next.next
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
