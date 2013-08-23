package euler

import (

)

func Euler7() int {
	x := primeSieve(200000)
	return x[10000]
}