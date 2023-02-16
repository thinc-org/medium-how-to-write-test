package example

import "errors"

func Divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("b cannot be 0")
	}

	return a / b, nil
}
