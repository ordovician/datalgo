package main

import "fmt"

// Find returns the index of first item in array xs which satisfies condition given
// as predicate function. If no matching value is found then -1 is returned
func Find(xs []int, condition func(int) bool) int {
	for i, x := range xs {
		if condition(x) {
			return i
		}
	}
	return -1
}

func main() {
	xs := []int{4, 6, 8, 10, 7}

	i := Find(xs, func(x int) bool {
		return x >= 8
	})

	if i == -1 {
		fmt.Println("No value >= 8 found in array")
	} else {
		fmt.Printf("Found value >= 8 at index %d\n", i)
	}
}
