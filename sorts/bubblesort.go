package sorts

//冒泡排序
func BubbleSort(arr []int) []int {
	//是否交换
	IsSwapped := true
	for IsSwapped {
		IsSwapped = false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i+1] < arr[i] {
				arr[i+1], arr[i] = arr[i], arr[i+1]
				IsSwapped = true
			}
		}
	}
	return arr
}
