package linkedlist

//链表节点
type DNode struct {
	val        int
	next, prev *DNode
}

//双向链表
type DoubleLinkedList struct {
	head, tail *DNode
}

func NewDoubleLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{
		head: nil,
		tail: nil,
	}
}

func NewDNode(val int) *DNode {
	return &DNode{
		val:  val,
		next: nil,
		prev: nil,
	}
}

//头插
func (dl *DoubleLinkedList) AddToHead(val int) {
	n := NewDNode(val)
	n.next = dl.head
	dl.head = n
}

//尾插
func (dl *DoubleLinkedList) AddToEnd(val int) {
	n := NewDNode(val)
	if dl.tail == nil {
		if dl.head == nil {
			dl.head = n
		}
		dl.tail = n
		return
	}
	dl.tail.next = n
	n.prev = dl.tail
	dl.tail = n

}
