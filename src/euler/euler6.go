package euler

import (

)

/**
* Project Euler #6 :
* 
* The sum of the squares of the first ten natural numbers is,
* 1^2 + 2^2 + ... + 10^2 = 385
* The square of the sum of the first ten natural numbers is,
* (1 + 2 + ... + 10)^2 = 55^2 = 3025
* Hence the difference between the sum of the squares of the
* first ten natural numbers and the square of the sum is
* 3025 − 385 = 2640.
* Find the difference between the sum of the squares of the
* first one hundred natural numbers and the square of the sum.
*/

func Euler6() int {
	const limit = 100
	return squareOfSum(limit) - sumOfSquares(limit)
}

func sumOfSquares(limit int) int {
	var squares []int
	for i := 1; i <= limit; i++ {
		squares = append(squares, i*i)  
	}
	return sum(squares)
}

func squareOfSum(limit int) int {
	var total int
	for i := 1; i <= limit; i++ {
		total += i
	}
	return total*total
}