package hamming

import (
	"errors"
)

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("strings must be equal length")
	}

	if len(a) == 0 || len(b) == 0 {
		return 0, nil
	}

	// we know both lengths are equal now
	// so we can do this
	hammingDist := 0
	for i := range len(a) {
		if a[i] != b[i] {
			hammingDist += 1
		}
	}

	return hammingDist, nil
}
