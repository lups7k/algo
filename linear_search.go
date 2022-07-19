package algo

func LinearSearch(list []int, target int) (result int) {
	// Returns the index position of the target if found, else returns None

	for i, item := range list {
		if item == target {
			return i
		}
	}
	return -1
}
