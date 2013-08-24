package euler

import (
	"strconv"
)

const gridSize = 20
var gridMap map[string]int

/**
* Project Euler #15 :
* 
* Starting in the top left corner of a 2x2 grid, and only
* being able to move to the right and down, there are exactly
* 6 routes to the bottom right corner.
* How many such routes are there through a 20x20 grid?
*/
func Euler15() int {
	gridMap = make(map[string]int)
	return trace(0,0)
}

func trace(x int, y int) int {
	i := 0
	key := strconv.Itoa(x) + "x" + strconv.Itoa(y)
	g := gridMap[key]
	if g != 0 {
		return g
	}
	if x < gridSize {
		i += trace(x + 1, y)
	}
	if y < gridSize {
		i += trace(x, y + 1)
	}
	if x == gridSize && y == gridSize {
		return 1
	}
	gridMap[key] = i

	return i
}
