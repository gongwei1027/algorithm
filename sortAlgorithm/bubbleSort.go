package sortalgorithm

type Sort interface {
	BubbleSort(nums []int) []int
	FastSort(nums []int) []int
	InsertSort(arr []int) []int
	MergeSort(arr []int) []int
	HeapSort(arr []int) []int
	CountSort(arr []int) []int
}

type sortF struct{}

type SortService struct {
	Sort Sort
}

func NewSortService() SortService {
	return SortService{
		Sort: &sortF{},
	}
}

//support BubbleSort
func (s *sortF) BubbleSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] <= nums[j] {
				continue
			} else {
				tmp := nums[j]
				nums[j] = nums[i]
				nums[i] = tmp
			}
		}
	}
	return nums
}

func (s *sortF) FastSort(nums []int) []int {
	fastSortIter(0, len(nums)-1, nums)
	return nums
}

func fastSortIter(left int, right int, nums []int) {
	if left < right {
		key := nums[left]
		i := left
		j := right
		for i < j {
			for i < j && key < nums[j] {
				j--
			}
			if i < j {
				nums[i] = nums[j]
				i++
			}
			for i < j && key > nums[i] {
				i++
			}
			if i < j {
				nums[j] = nums[i]
				j--
			}
		}
		nums[i] = key
		fastSortIter(left, i-1, nums)
		fastSortIter(i+1, right, nums)
	}
}

func (s *sortF) InsertSort(arr []int) []int {
	preIndex := 0
	curent := 0
	for i := 1; i < len(arr); i++ {
		curent = arr[i]
		preIndex = i - 1
		for preIndex >= 0 && arr[preIndex] > curent {
			arr[preIndex+1] = arr[preIndex]
			preIndex--
		}
		arr[preIndex+1] = curent
	}
	return arr
}

func (s *sortF) MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	medium := len(arr) / 2
	left := arr[0:medium]
	right := arr[medium:]

	return merge(s.MergeSort(left), s.MergeSort(right))

}

func merge(left []int, right []int) []int {
	var result []int
	for len(left) != 0 && len(right) != 0 {
		if left[0] < right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	for len(left) != 0 {
		result = append(result, left[0])
		left = left[1:]
	}
	for len(right) != 0 {
		result = append(result, right[0])
		right = right[1:]
	}
	return result
}

func (s *sortF) HeapSort(arr []int) []int {
	myHeap := newHeap()
	for _, a := range arr {
		myHeap.insert(a)
	}
	var re []int = []int{}
	for myHeap.N != 0 {
		re = append(re, myHeap.pop())
	}
	return re
}

func (s *sortF) CountSort(arr []int) []int {
	max := arr[0]
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	tmp := make([]int, max+1)
	index := 0
	for i := 0; i < len(arr)-1; i++ {
		tmp[arr[i]]++
	}
	var re []int = []int{}
	for index <= max {
		for tmp[index] > 0 {
			re = append(re, index)
			tmp[index]--
		}
		index++
	}
	return re
}
