package arraysort

func quickSort(arr []int) {
	_quickSort(arr, 0, len(arr)-1)
}

func _quickSort(arr []int, left, right int) {
	if left < right {
		partitionIndex := partition(arr, left, right)
		_quickSort(arr, left, partitionIndex-1)
		_quickSort(arr, partitionIndex+1, right)
	}
}

func partition(arr []int, left, right int) int {
	pivot := left
	index := pivot + 1

	for i := index; i <= right; i++ {
		if arr[i] < arr[pivot] {
			swap(arr, i, index)
			index++
		}
	}
	swap(arr, pivot, index-1)
	return index - 1
}
