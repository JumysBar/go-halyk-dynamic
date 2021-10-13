package main

import "fmt"

func main() {
	fmt.Println(Max([]int{1, 22, 3, 4, 5, 66, 7, 8, -9}))
}

func Max(numbers []int) int {
	var max int

	for _, element := range numbers {
		if element > max {
			max = element
		}
	}
	return max
}

func MaxIndex(numbers []int) int {
	var max int
	var maxIndex int

	for index, element := range numbers {
		if element > max {
			max = element
			maxIndex = index
		}
	}
	return maxIndex
}

func MaxWithError(numbers []int) (int, error) {
	if len(numbers) == 0 {
		return 0, fmt.Errorf("slice must be non empty")
	}
	var max int

	for _, element := range numbers {
		if element < 0 {
			return 0, fmt.Errorf("got negative number")
		}

		if element > max {
			max = element
		}
	}
	return max, nil
}
