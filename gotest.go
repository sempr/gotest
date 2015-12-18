package gotest

import (
	"errors"
)

func Division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("b should not be 0")
	}

	return a / b, nil
}
