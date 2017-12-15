package gotest

import "errors"

func Division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New(`divide zero`)
	}
	return a / b, nil
}
