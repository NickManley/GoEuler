package euler

import (
	"math/big"
	"strconv"
)

/**
* Project Euler #16 :
* 
* 2^15 = 32768 and the sum of its digits is
* 3 + 2 + 7 + 6 + 8 = 26.
*
* What is the sum of the digits of the number 2^1000?
*/

func Euler16() int {
	x := big.NewInt(2)
	y := big.NewInt(1000)
	z := big.NewInt(0)
	z.Exp(x, y, nil)
	total := 0
	for _,v := range z.String() {
		num, _ := strconv.Atoi(string(v))
		total += num
	}
	return total
}