package search

func BinarySearch(arr []int, target int, left int, right int) int {
	var mid int
	for left <= right {
		mid = (left + right) >> 1
		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
