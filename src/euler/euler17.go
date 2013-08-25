package euler

import (

)

/**
* Project Euler #17 :
* 
* If the numbers 1 to 5 are written out in words:
* one, two, three, four, five,
* then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.
*
* If all the numbers from 1 to 1000 (one thousand) inclusive were
* written out in words, how many letters would be used?
*
* NOTE: Do not count spaces or hyphens. For example, 342
* (three hundred and forty-two) contains 23 letters and 115
* (one hundred and fifteen) contains 20 letters.
* The use of "and" when writing out numbers is in compliance
* with British usage.
*/

func Euler17() int {
	total := 0
	for i := 1; i <= 1000; i++ {
		total += len(numberToName(i))
	}
	return total
}

func numberToName(num int) string {
	return parseHundred(num) + parseTen(num) + parseOne(num)
}

func parseHundred(num int) string {
	var result string
	var i int = num / 100
	var r int = num % 100
	switch {
		case i == 1: result += "onehundred"
		case i == 2: result += "twohundred"
		case i == 3: result += "threehundred"
		case i == 4: result += "fourhundred"
		case i == 5: result += "fivehundred"
		case i == 6: result += "sixhundred"
		case i == 7: result += "sevenhundred"
		case i == 8: result += "eighthundred"
		case i == 9: result += "ninehundred"
		case i == 10: result += "onethousand"
	}
	if r > 0 && i > 0 {
		result += "and"
	}
	return result
}

func parseTen(num int) string {
	var i int = (num - num/100*100) / 10
	switch {
		case i == 2: return "twenty"
		case i == 3: return "thirty"
		case i == 4: return "forty"
		case i == 5: return "fifty"
		case i == 6: return "sixty"
		case i == 7: return "seventy"
		case i == 8: return "eighty"
		case i == 9: return "ninety"
	}
	return ""
}

func parseOne(num int) string {
	var i int = (num - num/100*100)
	var r int = i % 10
	if i - 10 < 10 {
		switch {
			case i == 10: return "ten"
			case i == 11: return "eleven"
			case i == 12: return "twelve"
			case i == 13: return "thirteen"
			case i == 14: return "fourteen"
			case i == 15: return "fifteen"
			case i == 16: return "sixteen"
			case i == 17: return "seventeen"
			case i == 18: return "eighteen"
			case i == 19: return "nineteen"
		}
	}
	switch {
		case r == 1: return "one"
		case r == 2: return "two"
		case r == 3: return "three"
		case r == 4: return "four"
		case r == 5: return "five"
		case r == 6: return "six"
		case r == 7: return "seven"
		case r == 8: return "eight"
		case r == 9: return "nine"
	}
	return ""
}