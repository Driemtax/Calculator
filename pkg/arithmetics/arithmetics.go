package arithmetics

import "errors"

func Add(summand1 float64, summand2 float64) float64 {
	return summand1 + summand2
}

func Subtract(minuend float64, subtrahend float64) float64 {
	return minuend - subtrahend
}

func Multiply(factor1 float64, factor2 float64) float64 {
	return factor1 * factor2
}

func Divide(dividend float64, divisor float64) (float64, error) {
	if divisor == 0.0 {
		err := errors.New("division by zero")
		return 0.0, err
	}

	return dividend / divisor, nil
}
