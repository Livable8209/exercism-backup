package darts

import "fmt"

func Score(x, y float64) int {
	pointResult := isPointInOrOnRadius(x, y, 10)
	fmt.Printf("Point (%v,%v) in Circle with r = 10: %v\n", x, y, pointResult)
	return 0
}

func isPointInOrOnRadius(x, y, r float64) bool {
	// i could do this all in one line
	// but i rather not because this is
	// much more clear for anyone who
	// might read this (probably just myself)
	radiusSquared := r * r
	xSquared := x * x
	ySquared := y * y

	pointsSum := xSquared + ySquared

	return pointsSum <= radiusSquared
}
