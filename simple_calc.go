package simple_calc

import "github.com/Drawell/simple_calc/internal"

func Evaluate(text string) (float64, error) {
	return internal.Evaluate(text)
}
