package euler

import (
	"math"
	"sort"
)

/**
* Project Euler #12 :
* 
* The sequence of triangle numbers is generated by adding 
* the natural numbers. So the 7th triangle number would be
* 1 + 2 + 3 + 4 + 5 + 6 + 7 = 28. The first ten terms would be:
* 1, 3, 6, 10, 15, 21, 28, 36, 45, 55, ...
*
* Let us list the factors of the first seven triangle numbers:
*
*     1: 1
*     3: 1,3
*     6: 1,2,3,6
*    10: 1,2,5,10
*    15: 1,3,5,15
*    21: 1,3,7,21
*    28: 1,2,4,7,14,28
*
* We can see that 28 is the first triangle number to have
* over five divisors. What is the value of the first triangle
* number to have over five hundred divisors?
*/

func Euler12() int {
	i := 1
	for true {
		t := triangleNumber(i)
		if len(listFactors(t)) >= 500 {
			return t
		}
		i++
	}
	return 0
}

func triangleNumber(num int) int {
	var total int
	for i := 0; i <= num; i++ {
		total += i;
	}
	return total
}

func listFactors(num int) []int {
	limit := int(math.Ceil( math.Sqrt(float64(num)) ))
	var result []int
	for i := 1; i <= limit; i++ {
		if num % i == 0 {
			result = append(result, i, num/i)
		}
	}
	sort.Ints(result)
	return removeDuplicates(result)
}

func removeDuplicates(values []int) []int {
	var result []int
	for _,v := range values {
		i := len(result)-1
		if i <= 0 {
			result = append(result, v)
		} else if result[i] != v {
			result = append(result, v)
		}
	}
	return result
}


