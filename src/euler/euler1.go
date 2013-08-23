package euler

/**
* Project Euler #1 :
* 
* If we list all the natural numbers below 10 that are multiples
* of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
* Find the sum of all the multiples of 3 or 5 below 1000.
*/

func Euler1() int {
	return sum( naturalNumbers(1000, []int{3,5}) )
}

/**
* Returns a list of natural numbers from
* 1 to limit which are multiples of the values
* in mod. 
*/
func naturalNumbers(limit int, mods []int) []int {
	var result []int
	for i := 1; i < limit; i++ {
		
		isMult := false
		for _ , j := range mods {
			if i % j == 0 {
				isMult = true
				break
			}
		}
		
		if isMult {
			result = append(result, i)
		}
	}
	return result
}
