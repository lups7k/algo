package algo

func BinarySearch(list []int, target int) (result int) {
	first := 0
	last := len(list) - 1

	for first <= last {
		mid := (first + last) / 2
		if list[mid] == target {
			return mid
		}
		if list[mid] < target {
			first = mid + 1
		} else {
			last = mid - 1
		}
	}

	if first == len(list) || list[first] != target {
		return -1
	}

	return 0
}
