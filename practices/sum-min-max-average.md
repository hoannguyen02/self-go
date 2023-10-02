```
package main

import (
	"log"
)

func max(arr []int) int {
	if len(arr) == 0 {
		return -1;
	}
	max := arr[0]
	for _,number := range arr {
		if number > max {
			max = number
		}
	}
	return max
}

func min(arr []int) int {
	if len(arr) == 0 {
		return -1;
	}
	min := arr[0]
	for _,number := range arr {
		if number < min {
			min = number
		}
	}
	return min
}

func sum(arr []int) int {
	if len(arr) == 0 {
		return -1;
	}
	sum := 0
	for _,number := range arr {
		sum += number;
	}
	return sum
}

func average(arr []int) int {
	count := len(arr);
	if count == 0 {
		return -1;
	}
	return sum(arr) / count
}

func main() {
	numbers := []int{1, 3, 4, 5}
	
	log.Print("Sum: ", sum(numbers))
	log.Print("Max: ", max(numbers))
	log.Print("Min: ", min(numbers))
	log.Print("Average: ", average(numbers))
}
```