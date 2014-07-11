package main

import (
	"fmt"
	"reflect"
)

func main() {

	fmt.Println(lastIndexOf(3, []int{4, 3, 2, 1, 2, 6}))                      //1
	fmt.Println(lastIndexOf(3, []int{4, 3, 2, 1, 3, 6}))                      //4
	fmt.Println(lastIndexOf(3, []int{4, 1, 2, 1, 2, 6}))                      //-1
	fmt.Println(lastIndexOf(3, []int{4, 3, 2, 1, 3, 6, 7, 3, 2, 1, 3, 3, 3})) //something

}

func lastIndexOf(t int, array []int) int {

	var lastIndex func(array []int, t, i, temp int) int

	lastIndex = func(array []int, t, i, temp int) int {

		if reflect.DeepEqual(array, []int{}) {

			return temp

		} else if array[0] == t {

			temp = i
		}

		return lastIndex(array[1:], t, i+1, temp)
	}

	return lastIndex(array, t, 0, -1)
}
