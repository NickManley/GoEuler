package euler

import (
	"math"
)

/**
* Project Euler #3 :
* 
* The prime factors of 13195 are 5, 7, 13 and 29.
* What is the largest prime factor of the number 600851475143 ?
*/

func Euler3() int {
	const num = 600851475143
	var largestPrimeFactor int
	for i := 2; i < int(math.Sqrt(num)); i++ {
		if num % i == 0 && isPrime(i) {
			largestPrimeFactor = i
		}
	}
	return largestPrimeFactor
}