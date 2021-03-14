package main

import (
	"fmt"
)

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6}

	// New Slice
	sliceTwo := []int{7, 8, 9, 10, 11, 12, 13}

	// Append New Slice to Old Slice
	slice = append(slice, sliceTwo...)
	fmt.Println(slice)

}
