package euler

import (

)

/**
* Project Euler #14 :
* 
* The following iterative sequence is defined for the set
* of positive integers:
* n → n/2 (n is even)
* n → 3n + 1 (n is odd)
*
* Using the rule above and starting with 13,
* we generate the following sequence:
* 13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
*
* It can be seen that this sequence (starting at 13 and finishing at 1)
* contains 10 terms. Although it has not been proved yet (Collatz Problem),
* it is thought that all starting numbers finish at 1.
*
* Which starting number, under one million, produces the longest chain?
* NOTE: Once the chain starts the terms are allowed to go above one million.
*/

func Euler14() int {
	var result int
	var longestChain int
	for i := 999999; i >= 2; i-- {
		chainLen := len(collatz(i))
		if chainLen > longestChain {
			longestChain = chainLen
			result = i
		}
	}
	return result
}

func collatz(start int) []int {
	var nums []int
	nums = append(nums, start)
	for true {
		lastNum := nums[len(nums)-1]
		if lastNum == 1 {
			return nums
		}
		nums = append(nums, collatzNum(lastNum))
	}
	return nil
}

func collatzNum(num int) int {
	if num % 2 == 0 {
		return num / 2
	}
	return 3*num + 1
}