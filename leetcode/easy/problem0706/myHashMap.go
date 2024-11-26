package problem0706

type MyHashMap struct {
	buckets []*Node
}

type Node struct {
	key   int
	value int
	next  *Node
}

func Constructor() MyHashMap {
	return MyHashMap{
		buckets: make([]*Node, 16),
	}
}

func (this *MyHashMap) hash(key int) int {
	return key % len(this.buckets)
}

func (this *MyHashMap) Put(key int, value int) {
	h := this.hash(key)
	b := this.buckets[h]

	if b == nil {
		this.buckets[h] = &Node{key: key, value: value}
		return
	}

	for b != nil {
		if b.key == key {
			b.value = value
			return
		}
		if b.next == nil {
			b.next = &Node{key: key, value: value}
			return
		}
		b = b.next
	}
}

func (this *MyHashMap) Get(key int) int {
	h := this.hash(key)
	b := this.buckets[h]

	for b != nil {
		if b.key == key {
			return b.value
		}
		b = b.next
	}

	return -1
}

func (this *MyHashMap) Remove(key int) {
	h := this.hash(key)
	b := this.buckets[h]

	if b == nil {
		return
	}

	if b.key == key {
		this.buckets[h] = b.next
		return
	}

	for b.next != nil {
		if b.next.key == key {
			b.next = b.next.next
			return
		}
		b = b.next
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
