package lib

import (
	"strconv"
	"unicode"
)

func (e *Expression) consumeOperand() (float64, error) {
	e.skipSpaces()
	start := e.ptr
	for ; e.ptr < len(e.text) && (unicode.IsDigit(rune(e.text[e.ptr])) || e.text[e.ptr] == '.'); e.ptr++ {
	}

	if start == e.ptr {
		return 0, ParseError{"Unable to get operand", e.ptr}
	}

	operand, _ := strconv.ParseFloat(e.text[start:e.ptr], 64)
	return operand, nil
}

func execute(lOperand float64, operator string, rOperand float64) (float64, error) {
	switch operator {
	case "+":
		return lOperand + rOperand, nil
	case "-":
		return lOperand - rOperand, nil
	case "*":
		return lOperand * rOperand, nil
	case "/":
		if rOperand == 0 {
			return 0, ParseError{"Zero division occurs", 0}
		} else {
			return lOperand / rOperand, nil
		}
	default:
		return lOperand, nil
	}
}

func (e *Expression) nextChar() uint8 {
	e.skipSpaces()
	return e.text[e.ptr]
}

func (e *Expression) consume(chars []uint8) (uint8, bool) {
	e.skipSpaces()
	if e.ptr >= len(e.text) {
		return ' ', false
	}

	for _, char := range chars {
		if e.text[e.ptr] == char {
			e.ptr++
			return char, true
		}
	}
	return ' ', false
}

func (e *Expression) skipSpaces() {
	for ; e.ptr < len(e.text) && e.text[e.ptr] == ' '; e.ptr++ {
	}
}
