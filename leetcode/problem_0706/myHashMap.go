package problem_0706

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

func (m *MyHashMap) hash(key int) int {
	return key % len(m.buckets)
}

func (m *MyHashMap) Put(key int, value int) {
	h := m.hash(key)
	b := m.buckets[h]

	if b == nil {
		m.buckets[h] = &Node{key: key, value: value}
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

func (m *MyHashMap) Get(key int) int {
	h := m.hash(key)
	b := m.buckets[h]

	for b != nil {
		if b.key == key {
			return b.value
		}
		b = b.next
	}

	return -1
}

func (m *MyHashMap) Remove(key int) {
	h := m.hash(key)
	b := m.buckets[h]

	if b == nil {
		return
	}

	if b.key == key {
		m.buckets[h] = b.next
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
