package link


/**
 *  @ClassName:LRUCache
 *  @Description:leetcode 146 LRU缓存实现
 *  @Author:jackey
 *  @Create:2021/4/2 下午2:06
 */

/*双向链表按照被使用的顺序存储了这些键值对，靠近头部的键值对是最近使用的，而靠近尾部的键值对是最久未使用的。*/

/*定义LRU缓存结构体，添加size和cap变量*/
type LRUCache struct {
	size int
	capacity int
	cache map[int]*DLinkedNode
	head, tail *DLinkedNode
}

/*定义双链表结构体*/
type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

/*初始化双向链表*/
func initDLinkedNode(key, value int) *DLinkedNode {
	return &DLinkedNode{
		key: key,
		value: value,
	}
}

/*初始化LRU缓存结构体*/
func ConstructorLRU(capacity int) LRUCache {
	l := LRUCache{
		cache: map[int]*DLinkedNode{},
		head: initDLinkedNode(0, 0),
		tail: initDLinkedNode(0, 0),
		capacity: capacity,
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

/*get的操作：
 * 1、查询是否已经在缓存中，如果是，则将key移动到表头位置
 * 2、如果没有返回-1
 */
func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	node := this.cache[key]
	this.moveToHead(node)
	return node.value
}

/*cache插入操作：
 * 1、如果已经存在，则将它对应的val修改
 * 2、如果不修改则插入
 * 3、如果size已经达到cap，则需要移出队尾节点
 * 4、如果没有不够cap，则添加到表头
*/
func (this *LRUCache) Put(key int, value int)  {
	if _, ok := this.cache[key]; !ok {
		node := initDLinkedNode(key, value)
		this.cache[key] = node
		this.addToHead(node)
		this.size++
		if this.size > this.capacity {
			removed := this.removeTail()
			delete(this.cache, removed.key)
			this.size--
		}
	} else {
		node := this.cache[key]
		node.value = value
		this.moveToHead(node)
	}
}

/*添加节点*/
func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

/*移出节点*/
func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

/*移到表头*/
func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

/*去除尾节点*/
func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}