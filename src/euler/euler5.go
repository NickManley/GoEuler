package euler

import (

)

/**
* Project Euler #5 :
* 
* 2520 is the smallest number that can be divided by
* each of the numbers from 1 to 10 without any remainder.
* What is the smallest positive number that is evenly 
* divisible by all of the numbers from 1 to 20?
*/

func Euler5() int {
	const limit = 20
	num := limit*limit
	for true {
		isDivisible := true
		for i := 2; i <= limit; i++ {
			if num % i != 0 {
				isDivisible = false
				break
			}
		}
		if isDivisible {
			return num
		}
		num += limit
	}
	return num
}
