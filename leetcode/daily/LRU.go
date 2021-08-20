package daily

/**
 *  @ClassName:LRU
 *  @Description:最近最少未使用
 *  @Author:jackey
 *  @Create:2021/8/20 下午8:08
 */

// 是一个map加双向链表
type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*DLinkedNode
	head, tail *DLinkedNode
}

type DLinkedNode struct {
	key, value int
	prev, post *DLinkedNode
}

//初始化
func initDLinkedNode(key, value int) *DLinkedNode {
	return &DLinkedNode{
		key:   key,
		value: value,
	}
}

// 构建LRU
func Constructor(cap int) LRUCache {
	l := LRUCache{
		size:     0,
		capacity: cap,
		cache:    make(map[int]*DLinkedNode),
		head:     initDLinkedNode(0, 0),
		tail:     initDLinkedNode(0, 0),
	}
	l.tail.prev = l.head
	l.head.post = l.tail
	return l
}

// 使用
func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	node := this.cache[key]
	this.moveToHead(node)
	return node.value
}

// 添加
func (this *LRUCache) Put(key, value int) {
	if _, ok := this.cache[key]; !ok {
		node := initDLinkedNode(key, value)
		// 添加进表头
		this.cache[key] = node
		this.addToHead(node)
		this.size++
		if this.size > this.capacity {
			removed := this.removeTail()
			delete(this.cache, removed.key)
			this.size++
		}
	} else {
		node := this.cache[key]
		node.value = value
		this.moveToHead(node)
	}
}

// 最近使用的一定要添加到前面去
func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = this.head
	node.post = this.head.post
	node.post.prev = node
	this.head.post = node
}

// 去除
func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.post = node.post
	node.post.prev = node.prev
}

// 移动到头部
func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

// 移出尾部
func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}
