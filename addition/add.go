package addition

import "errors"

var ErrInvalidSummand = errors.New("Invalid summand")

func Add(x, y int) (int, error) {
	if x < 0 || y < 0 {
		return 0, ErrInvalidSummand
	} else if x == 0 && y == 0 {
		return 0, nil
	} else {
		return x + y, nil
	}
}
