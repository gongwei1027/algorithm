package sortalgorithm

type Myheap interface {
	sink(i int)
	swim(i int)
	insert(v int)
	pop() int
}

type selfHeap struct {
	arr []int
	N   int
}

func newHeap() selfHeap {
	N := 0
	var arr []int = []int{}
	arr = append(arr, 0)
	return selfHeap{
		arr: arr,
		N:   N,
	}
}

func (h *selfHeap) sink(i int) {
	for {
		left := 2 * i
		right := 2*i + 1
		min := left
		if left > h.N {
			break
		}
		if left < h.N && h.arr[right] < h.arr[left] {
			min = right
		}
		if right < h.N && h.arr[i] < h.arr[min] {
			break
		}
		h.swap(h.arr, i, min)
		i = min
	}

}

func (h *selfHeap) swim(i int) {
	for i > 1 && h.arr[i] < h.arr[i/2] {
		h.swap(h.arr, i, i/2)
		i = i / 2
	}
}

func (h *selfHeap) swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func (h *selfHeap) insert(v int) {
	h.arr = append(h.arr, v)
	h.N++
	h.swim(h.N)
}

func (h *selfHeap) pop() int {
	if h.N == 0 {
		return 0
	}
	min := h.arr[1]
	h.swap(h.arr, 1, h.N)
	h.N--
	h.sink(1)
	return min
}
