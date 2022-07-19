package algo

func RecursiveBinarySearch(list []int, target int) (result bool) {
	if len(list) == 0 {
		return false
	} else {
		mid := len(list) / 2

		if list[mid] == target {
			return true
		} else {
			if list[mid] < target {
				return RecursiveBinarySearch(list[mid+1:], target)
			} else {
				return RecursiveBinarySearch(list[:mid], target)
			}
		}
	}
}
