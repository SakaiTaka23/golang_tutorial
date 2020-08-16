package gotest

import (
	"errors"
)

// Division 割り算関数
func Division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数は0以外でなければなりません")
	}

	return a / b, nil
}
