package main

import "fmt"
func main() {
	target := 4
	sl := []int{1,2,3,4}
	fmt.Println(twoSum(sl, target))
}

func twoSum(nums []int, target int) []int {
    complementMap := map[int]int{}
    for i, n := range nums {
        c := target - n
        if j, ok := complementMap[c]; ok {
            return []int{j, i}
        }
        complementMap[n] = i
    }
    return []int{}
}
