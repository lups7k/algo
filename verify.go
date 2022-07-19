package algo

import "fmt"

// verify algorithms

func Verify(index any) {
	if index != nil {
		fmt.Println("Target found at index: ", index)
	} else {
		fmt.Println("Target not found in list")
	}
}

func VerifySorted(list []int) (result bool) {
	n := len(list)

	if n == 0 || n == 1 {
		return true
	}

	return list[0] < list[1] && VerifySorted(list[1:])
}
