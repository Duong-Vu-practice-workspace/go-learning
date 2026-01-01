package __custom_error_interface

import (
	"fmt"
)

type divideError struct {
	dividend float64
}

// ?
func (dError divideError) Error() string {
	return fmt.Sprintf("can not divide %v by zero", dError.dividend)
}
func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, divideError{dividend: dividend}
	}
	return dividend / divisor, nil
}
