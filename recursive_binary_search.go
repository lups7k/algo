package algo

import "fmt"

func recursiveBinarySearch(list []int, target int) (result bool) {
	if len(list) == 0 {
		return false
	} else {
		mid := len(list) / 2

		if list[mid] == target {
			return true
		} else {
			if list[mid] < target {
				return recursiveBinarySearch(list[mid+1:], target)
			} else {
				return recursiveBinarySearch(list[:mid], target)
			}
		}
	}
}

func verify(index any) {
	if index != nil {
		fmt.Println("Target found at index: ", index)
	} else {
		fmt.Println("Target not found in list")
	}
}