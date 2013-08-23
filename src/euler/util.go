package euler

/**
 * Contains various utility functions
 * used for multiple Project Euler problems.
 */

import (
	"math"
)

/**
* Return the sum of all elements in array
*/
func sum(x []int) int {
	total := 0
	for _, i := range x {
		total += i
	}
	return total
}

func isPrime(num int) bool {
	sq := int(math.Ceil( math.Sqrt( float64(num) )))
	prime := true
	
	// special case, 2 is prime
	if num == 2 {
		return true
	}
	
	for i := 2; i <= sq; i++ {
		if num % i == 0 {
			prime = false
			break
		}
	}
	return prime
}

func primeSieve(limit int) []int {
	var nums []int
	prime := 2
	primeIndex := 0
	for i := 2; i <= limit; i++ {
		nums = append(nums, i)
	}
	
	for true {
		prime = nums[primeIndex]
		for i := 0; i < len(nums); i++ {
			if nums[i] % prime == 0 && nums[i] > prime {
				nums = append(nums[:i], nums[i+1:]...)
			}
		}
		primeIndex++
		if primeIndex >= len(nums) {
			break
		}
	}
	return nums
}

func reverseString(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func isPalindrome(s string) bool {
	return s == reverseString(s)
}