package euler

import (
	"strconv"
)

/**
* Project Euler #4 :
*
* A palindromic number reads the same both ways.
* The largest palindrome made from the product of
* two 2-digitnumbers is 9009 = 91 × 99.
* Find the largest palindrome made from the product
* of two 3-digit numbers.
 */

func Euler4() int {
	var largestPali int
	for i := 999; i >= 100; i-- {
		for j := 999; j >= 100; j-- {
			isPali := isPalindrome(strconv.Itoa(i*j))
			if  isPali && i*j > largestPali {
				largestPali = i * j
			}
		}
	}
	return largestPali
}
