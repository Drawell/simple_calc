package calc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEvaluate(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		expected   float64
	}{
		{"1+1", "1 + 1", 2.},
		{"1+1+1", "1 + 1 + 1", 3.},
		{"2-1", "2 - 1", 1.},
		{"2*3", "2 * 3", 6.},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := Evaluate(testCase.expression)
			if err != nil || (res-testCase.expected)*(res-testCase.expected) > 0.001 {
				t.Errorf("Eval expr %v, expected %v, but get %v, '%v'",
					testCase.expression, testCase.expected, res, err)
			}
		})
	}
}

func TestEvaluate_testify(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		expected   float64
	}{
		{"5/2", "5 / 2", 2.5},
		{"2 * (2 + 2)", "2 * (2 + 2)", 8.},
		{"many_spaces", "    1    +  1  ", 2.},
		{"many_brackets", "(((1 + 1)) + 1)", 3.},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := Evaluate(testCase.expression)
			assert.Nil(t, err)
			assert.Equal(t, res, testCase.expected)
		})
	}
}

func TestEvaluateInvalid(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		errorMsg   string
	}{
		{"zero div", "2 / 0", "Zero division occurs at 0"},
		{"zero div 2", "1 + (2 / 0)", "Zero division occurs at 0"},
		{"zero div 3", "2 * (2 / 0)", "Zero division occurs at 0"},
		{"no closing bracket", "(1 + 1", "There is not closing bracket at 6"},
		{"no opening bracket", "1 + 1)", "There is not open bracket for closing at 5"},
		{"invalid_symbol_left_operand", "a + 1", "Unable to get operand at 0"},
		{"invalid_symbol_right_operand", "1 + a", "Unable to get operand at 4"},
		{"invalid_operation", "1 a 1", "Invalid symbols at 2"},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := Evaluate(testCase.expression)
			if err.Error() != testCase.errorMsg || (res-0)*(res-0) > 0.001 {
				t.Errorf("Eval expr %v, expected error '%v' and but get '%v'",
					testCase.expression, testCase.errorMsg, err)
			}
		})
	}
}

func BenchmarkEvaluate(b *testing.B) {
	expression := "(1 + 5) * 2 / 3"
	for b.Loop() {
		_, err := Evaluate(expression)
		if err == nil {
			b.Errorf("Benchmark error %v", err)
		}
	}
}
