package euler

import (
	"math"
)

func Euler3() int {
	const num = 600851475143
	var largestPrimeFactor int
	for i := 2; i < int(math.Sqrt(num)); i++ {
		if num % i == 0 && IsPrime(i) {
			largestPrimeFactor = i
		}
	}
	return largestPrimeFactor
}