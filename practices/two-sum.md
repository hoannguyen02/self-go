Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Output: Because nums[0] + nums[1] == 9, we return [0, 1].
```
package main

import "fmt"

func twoSum(numbers []int, target int, c chan []int) {
	seenNum := make(map[int]int)
	for index, a := range numbers {
		numToFind := target - a;
		val, ok := seenNum[numToFind]
		if ok {
			c <- []int{val, index}
			return
		} else {
			seenNum[a]=index
		}
	}
	

}

func main() {
	numbers := []int{1, 7, 2, 15}
	target := 9
	c := make(chan []int)
	go twoSum(numbers, target, c)
	fmt.Println(<-c)
}
```