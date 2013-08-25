package euler

import (
	"strconv"
)

/**
* Project Euler #20 :
* 
* 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
* and the sum of the digits in the number 10! is
* 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.
*
* Find the sum of the digits in the number 100!
*/

func Euler20() int {
	x := factorial(100)
	result := 0
	for _, v := range x.String() {
		digit, _ := strconv.Atoi(string(v))
		result += digit
	}
	return result
}
