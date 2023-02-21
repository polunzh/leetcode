package kthlargestelementinastream

import (
	"container/heap"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}

func (h *IntHeap) Pop() interface{} {
	top := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]

	return top
}

type KthLargest struct {
	data *IntHeap
	k    int
}

func Constructor(k int, nums []int) KthLargest {
	kth := KthLargest{
		k:    k,
		data: &IntHeap{},
	}

	for i := 0; i < len(nums); i++ {
		heap.Push(kth.data, nums[i])

		if kth.data.Len() > k {
			heap.Pop(kth.data)
		}
	}

	return kth
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.data, val)

	if this.data.Len() > this.k {
		heap.Pop(this.data)
	}

	return (*this.data)[0]
}
