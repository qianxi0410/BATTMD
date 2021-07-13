package hot100

type LRUCache struct {
	head     *node
	tail     *node
	nodeMap  map[int]*node
	capacity int
	size     int
}

type node struct {
	prev  *node
	next  *node
	key   int
	value int
}

func __Constructor(capacity int) LRUCache {
	lru := LRUCache{
		head:     new(node),
		tail:     new(node),
		nodeMap:  make(map[int]*node),
		capacity: capacity,
	}
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.nodeMap[key]; ok {
		this.removeNode(node)
		this.appendAfterHead(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.nodeMap[key]; ok {
		node.value = value
		this.removeNode(node)
		this.appendAfterHead(node)
		return
	}
	if this.size == this.capacity {
		toDel := this.tail.prev
		delete(this.nodeMap, toDel.key)
		this.removeNode(toDel)
		this.size--
	}
	this.size++
	p := &node{
		key:   key,
		value: value,
	}
	this.nodeMap[key] = p
	this.appendAfterHead(p)
}

func (this *LRUCache) appendAfterHead(node *node) {
	node.next = this.head.next
	node.prev = this.head
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *node) {
	prev := node.prev
	prev.next = node.next
	node.next.prev = prev
}
