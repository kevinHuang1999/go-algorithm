package sorts

type maxHeap struct {
	heap []int
	size int
}

//构建最大堆
func buildMaxHeap(arr []int) *maxHeap {
	h := &maxHeap{heap: arr, size: len(arr)}
	//从第一个非叶子节点开始调整
	for i := len(arr) / 2; i >= 0; i-- {
		h.Heapify(i)
	}
	return h
}

//堆排序
func HeapSort(arr []int) []int {
	h := buildMaxHeap(arr)
	for i := len(h.heap) - 1; i >= 1; i-- {
		h.heap[0], h.heap[i] = h.heap[i], h.heap[0]
		h.size--
		h.Heapify(0)
	}
	return h.heap
}

//向下调整
func (h *maxHeap) Heapify(i int) {
	//左右节点
	l, r := (i>>1)+1, (i>>1)+2
	max := i
	//左右节点最大值上提
	if l < h.size && h.heap[l] > h.heap[max] {
		max = l
	}
	if r < h.size && h.heap[r] > h.heap[max] {
		max = r
	}
	if max != i {
		h.heap[i], h.heap[max] = h.heap[max], h.heap[i]
		//递归调整
		h.Heapify(max)
	}
}
