package maps

import (
	"fmt"
	"hash/fnv"
)

var defaultCapacity uint64 = 1 << 4

type node struct {
	key  interface{}
	val  interface{}
	next *node
}

type HashMap struct {
	capacity uint64
	size     uint64
	table    []*node
}

func NewHashMap() *HashMap {
	return &HashMap{
		capacity: defaultCapacity,
		table:    make([]*node, defaultCapacity),
	}
}

//获取节点值
func (h *HashMap) Get(key interface{}) interface{} {
	node := h.getNode(h.hashKey(key), key)
	if node != nil {
		return node.val
	}
	return nil
}

func (h *HashMap) getNode(hash uint64, key interface{}) *node {
	node := h.table[hash]
	if node != nil {
		if node.key != key {
			for e := node.next; e != nil; e = e.next {
				if e.key == key {
					return e
				}
			}
			return nil
		}
		return node
	}
	return nil
}

func (h *HashMap) ContainsKey(key interface{}) bool {
	return h.getNode(h.hashKey(key), key) != nil
}

//新增节点
func (h *HashMap) Put(key interface{}, val interface{}) interface{} {
	return h.putValue(h.hashKey(key), key, val)
}

func (h *HashMap) putValue(hash uint64, key interface{}, value interface{}) interface{} {
	//未初始化，则初始化
	if h.capacity == 0 {
		h.capacity = defaultCapacity
		h.table = make([]*node, defaultCapacity)
	}
	Node := h.table[hash]
	var e *node = nil
	var p = Node
	if Node == nil {
		h.table[hash] = NewNode(key, value)
	} else {
		for e = Node; e != nil; e = e.next {
			if e.key == key {
				break
			}
			p = e
		}
		if e == nil {
			p.next = NewNode(key, value)
		}
	}
	if e != nil {
		oldValue := e.val
		e.val = value
		return oldValue
	}
	h.size++
	return nil
}

//创建下一个节点
func NewNode(key interface{}, value interface{}) *node {
	return &node{
		key:  key,
		val:  value,
		next: nil,
	}
}

//扩容
func (h *HashMap) resize() {
	h.capacity <<= 1
	tmpTable := h.table
	h.table = make([]*node, h.capacity)
	for i := 0; i < len(tmpTable); i++ {
		n := tmpTable[i]
		if n == nil {
			continue
		}
		h.table[h.hashKey(n.key)] = n
	}
}

//获取hash值
func (h *HashMap) hashKey(key interface{}) uint64 {
	f := fnv.New64a()
	_, _ = f.Write([]byte(fmt.Sprintf("%v", key)))
	hashValue := f.Sum64()
	return (h.capacity - 1) & (hashValue ^ (hashValue >> 32))
}
