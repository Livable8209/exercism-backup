package darts

// 3 + 1 rings
// 1, 5, 10, n >= 11 (fail case)

func Score(x, y float64) int {
	if isPointInOrOnRadius(x, y, 1) {
		return 10
	}

	if isPointInOrOnRadius(x, y, 5) {
		return 5
	}

	if isPointInOrOnRadius(x, y, 10) {
		return 1
	}

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
