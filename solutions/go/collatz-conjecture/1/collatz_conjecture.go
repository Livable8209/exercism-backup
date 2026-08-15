package collatzconjecture

import (
	"errors"
)

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("n must be a positive integer.")
	}

	steps := 0
	for {
		if n == 1 {
			break
		}

		if n%2 == 0 {
			n /= 2 // i never remember these operators exist until i have to look them up
		} else {
			n = (n * 3) + 1
		}
		steps++ // forgot ++ was a thing
	}

	return steps, nil
}
