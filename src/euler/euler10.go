package euler

import (

)

/**
* Project Euler #10 :
* 
* The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
* Find the sum of all the primes below two million.
*/

func Euler10() int {
	return sum( primeSieve(2000000) )
}