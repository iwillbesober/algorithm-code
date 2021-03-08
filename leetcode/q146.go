package leetcode

// 运用你所掌握的数据结构，设计和实现一个 LRU (最近最少使用) 缓存机制 。
//
//
//
// 实现 LRUCache 类：
//
//
// LRUCache(int capacity) 以正整数作为容量 capacity 初始化 LRU 缓存
// int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
// void put(int key, int value) 如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字-值」。当缓存容量达到上
// 限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。
//
//
//
//
//
//
// 进阶：你是否可以在 O(1) 时间复杂度内完成这两种操作？
//
//
//
// 示例：
//
//
// 输入
// ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
// [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
// 输出
// [null, null, null, 1, null, -1, null, -1, 3, 4]
//
// 解释
// LRUCache lRUCache = new LRUCache(2);
// lRUCache.put(1, 1); // 缓存是 {1=1}
// lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
// lRUCache.get(1);    // 返回 1
// lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
// lRUCache.get(2);    // 返回 -1 (未找到)
// lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
// lRUCache.get(1);    // 返回 -1 (未找到)
// lRUCache.get(3);    // 返回 3
// lRUCache.get(4);    // 返回 4
//
//
//
//
// 提示：
//
//
// 1 <= capacity <= 3000
// 0 <= key <= 3000
// 0 <= value <= 104
// 最多调用 3 * 104 次 get 和 put
//
// Related Topics 设计
// 👍 1217 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
type LRUCache struct {
	cmap     map[int]*Node
	dlist    *DoubleLinkedList
	capacity int
}

func (this *LRUCache) setMap(key int, val *Node) {
	if this.cmap == nil {
		this.cmap = make(map[int]*Node, this.capacity)
	}
	this.cmap[key] = val
}

type Node struct {
	key, val  int
	pre, next *Node
}

type DoubleLinkedList struct {
	size       int
	head, last *Node
}

func (l *DoubleLinkedList) GetSize() int {
	return l.size
}

func (l *DoubleLinkedList) AddFirst(node *Node) {
	if l.size == 0 {
		node.pre = nil
		node.next = nil
		l.head = node
		l.last = node
	} else {
		h := l.head
		h.pre = node
		l.head = node
		l.head.pre = nil
		l.head.next = h
	}
	l.size++
}

func (l *DoubleLinkedList) Remove(node *Node) {
	if l.size == 0 || node == nil {
		return
	}
	pre := node.pre
	next := node.next
	if pre != nil {
		pre.next = next
	}
	if next != nil {
		next.pre = pre
	}
	if l.head == node {
		l.head = next
	}
	if l.last == node {
		l.last = pre
	}
	l.size--
}

func (l *DoubleLinkedList) RemoveLast() *Node {
	if l.size == 0 {
		return nil
	}
	last := l.last
	l.Remove(l.last)
	return last
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		dlist:    new(DoubleLinkedList),
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	v, ok := this.cmap[key]
	if !ok {
		return -1
	}
	this.Put(v.key, v.val)
	return v.val
}

func (this *LRUCache) Put(key int, value int) {
	node := &Node{
		key: key,
		val: value,
	}
	v, ok := this.cmap[key]
	if ok {
		this.dlist.Remove(v)
		this.dlist.AddFirst(node)
		this.setMap(key, node)
	} else {
		if this.dlist.GetSize() == this.capacity {
			last := this.dlist.RemoveLast()
			delete(this.cmap, last.key)
		}
		this.setMap(key, node)
		this.dlist.AddFirst(node)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
//leetcode submit region end(Prohibit modification and deletion)