package algo

func mergeSort(list []int) (result []int) {
	/*
	 Sorts a list in ascending order
	 Returns a new sorted list

	 Divide -> Find the midpoint of the list and divide into sublists
	 Conquer -> Recursively sort the sublists created in previous step
	 Combine -> Merge the sorted sublists created in previous step

	 Overall -> O(kn log n)
	*/

	if len(list) <= 1 {
		return list
	}

	var leftHalf, rightHalf []int = split(list)
	var left []int = mergeSort(leftHalf)
	var right []int = mergeSort(rightHalf)

	mergeReturn := merge(left, right)
	return mergeReturn
}

func split(list []int) (resul []int, result []int) {
	// divide unsorted list at midpoint into sublists
	// returns two sublists - left and right
	// overall -> O(k log n)

	mid := len(list) / 2
	left := list[:mid]
	right := list[mid:]

	return left, right
}

func merge(left []int, right []int) (result []int) {
	// Merges two lists, sorting them in the process
	// Returns a new merged list

	var l []int
	i := 0 // right list
	j := 0 // left list

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			l = append(l, left[i])
			i += 1
		} else {
			l = append(l, right[j])
			j += 1
		}
	}

	for i < len(left) {
		l = append(l, left[i])
		i += 1
	}
	for j < len(right) {
		l = append(l, right[j])
		j += 1
	}
	return l
}

func verifySorted(list []int) (result bool) {
	n := len(list)

	if n == 0 || n == 1 {
		return true
	}

	return list[0] < list[1] && verifySorted(list[1:])
}
