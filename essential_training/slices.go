package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{16, 8, 42, 4, 23, 15}
	fmt.Println("My solution")
	sort.Ints(nums)
	fmt.Println(nums[len(nums)-1])

	fmt.Println("--------")
	fmt.Println("Tutorial solution")

	max := nums[0]
	for _, value := range nums[1:] {
		if value > max {
			max = value
		}
	}
	fmt.Println(max)
}
