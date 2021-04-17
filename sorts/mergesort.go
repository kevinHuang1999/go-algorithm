package sorts

func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	var mid = len(arr) / 2
	var left = MergeSort(arr[:mid])
	var right = MergeSort(arr[mid:])
	return merge(left, right)
}

//合并两个有序数组
func merge(a []int, b []int) []int {
	var sum = make([]int, len(a)+len(b))
	var l = 0
	var r = 0
	for l < len(a) && r < len(b) {
		if a[l] <= b[r] {
			sum[l+r] = a[l]
			l++
		} else {
			sum[l+r] = b[r]
			r++
		}
	}
	for l < len(a) {
		sum[l+r] = a[l]
		l++
	}
	for r < len(b) {
		sum[l+r] = b[r]
		r++
	}
	return sum
}
