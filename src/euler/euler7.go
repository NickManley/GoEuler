package euler

import (

)

func Euler7() int {
	x := primeSieve(1000000)
	return x[10000]
}