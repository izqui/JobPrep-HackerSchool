package main

import "fmt"

func main() {

	fmt.Println(sum([]int{1, 2, 3, 10, 2882, 21, 2, 12, 3, 10, 1}))
}

func sum(array []int) int {

	if len(array) == 2 {

		return array[0] + array[1]
	} else {

		return array[0] + sum(array[1:])
	}
}
