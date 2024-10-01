package containers

// Node 定义链表节点
type Node[T any] struct {
	data T
	prev *Node[T] // 节点前驱指针
	next *Node[T] // 节点后继指针
}

// GetData 返回节点的数据
func (n *Node[T]) GetData() T {
	return n.data
}

// Next 返回当前节点的下一个节点
func (n *Node[T]) Next() *Node[T] {
	return n.next
}

// Prev 返回当前节点的前一个节点
func (n *Node[T]) Prev() *Node[T] {
	return n.prev
}

// List 定义链表结构
type List[T any] struct {
	head *Node[T] // 链表头指针
	tail *Node[T] // 链表尾指针
	size int      // 节点个数
}

// NewList 返回一个新的空链表
func NewList[T any]() *List[T] {
	return &List[T]{}
}

// 头插法
// InsertAtHead 在链表头部插入新节点
// （方法）需要绑定接收者
func (list *List[T]) InsertAtHead(data T) {
	// 新建节点
	newNode := &Node[T]{data: data}
	// 头空：链表为空
	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		newNode.next = list.head
		list.head.prev = newNode
		list.head = newNode
	}
	list.size++
}

// 尾插法
// InsertAtTail 在链表尾部插入新节点
// （方法）需要绑定接收者
func (list *List[T]) InsertAtTail(data T) {
	// 新建节点
	newNode := &Node[T]{data: data}
	// 尾空：链表空
	if list.tail == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		newNode.prev = list.tail
		list.tail.next = newNode
		list.tail = newNode
	}
	list.size++
}

// 指定节点之后插入
// InsertAfter
// （方法）需要绑定接收者
func (list *List[T]) InsertAfter(node *Node[T], data T) {
	// 判断节点合法性
	if node == nil {
		return
	}
	// 新建节点
	newNode := &Node[T]{data: data}
	newNode.prev = node
	newNode.next = node.next
	// 判断节点是否为尾节点
	if node.next != nil {
		newNode.next.prev = node
	} else {
		list.tail = newNode
	}
	node.next = newNode
	list.size++
}

// 头删法
// RemoveAtHead
// （方法）需要绑定接收者
func (list *List[T]) RemoveAtHead() {
	// 链表合法性判断（空？）
	if list.head == nil {
		return
	}
	// 头指针后移（注意可能指向空）
	list.head = list.head.next
	if list.head != nil {
		list.head.prev = nil
	} else {
		list.tail = nil
	}
	list.size--
}

// 尾删法
// RemoveAtTail
// （方法）需要绑定接收者
func (list *List[T]) RemoveAtTail() {
	// 链表合法性判断（空？）
	if list.tail == nil {
		return
	}
	// 尾指针前移
	list.tail = list.tail.prev
	if list.tail != nil {
		list.tail.next = nil
	} else {
		list.head = nil
	}
	list.size--
}

// 删除指定节点
func (list *List[T]) RemoveNode(node *Node[T]) {
	// 节点合法性判断
	if node == nil {
		return
	}
	// 判断前面是否有节点
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		list.head = node.next
	}
	// 判断后面是否有节点
	if node.next != nil {
		node.next.prev = node.prev
	} else {
		list.tail = node.prev
	}
	list.size--
}

func (list *List[T]) Size() int {
	return list.size
}

// 清空链表
func (list *List[T]) Clear() {
	list.head = nil
	list.tail = nil
	list.size = 0
}

// 遍历：顺序
func (list *List[T]) TraverseForward(f func(data T)) {
	for node := list.head; node != nil; node = node.next {
		f(node.data)
	}
}

func (list *List[T]) TraverseBackward(f func(data T)) {
	for node := list.tail; node != nil; node = node.prev {
		f(node.data)
	}
}

func (list *List[T]) Head() *Node[T] {
	return list.head
}

func (list *List[T]) Tail() *Node[T] {
	return list.tail
}
