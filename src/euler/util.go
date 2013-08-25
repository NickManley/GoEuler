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
	sqLimit := int(math.Ceil(math.Sqrt(float64(limit))))
	array := make([]bool, limit)
	output := make([]int, 0)
	
	for i := 0; i < limit; i++ {
		array[i] = true
	}

	for i := 2; i <= sqLimit; i++ {
		if array[i] {
			for j := i*i; j < limit; j += i {
				array[j] = false
			}
		}
	}

	for i := 2; i < limit; i++ {
		if array[i] {
			output = append(output, i)
		}
	}
	return output
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

func max(nums []int) int {
	result := nums[0]
	for _,v := range nums {
		if v > result {
			result = v
		}
	}
	return result
}
