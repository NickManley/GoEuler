package euler

import (
	"math"
)

/**
* Project Euler #9 :
* 
* A Pythagorean triplet is a set of three natural numbers,
* a < b < c, for which, a^2 + b^2 = c^2
* There exists exactly one Pythagorean triplet for which
* a + b + c = 1000.
* Find the product abc.
*/

func Euler9() int {
	for a := float64(2); a < 1000; a++ {
		for b := float64(a+1); b < 1000; b++ {
			c := math.Sqrt(a*a + b*b)
			if a+b+c == 1000 {
				return int(a*b*c)
			}
		}
	}
	return 0
}