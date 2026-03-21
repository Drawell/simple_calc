package lib

type Expression struct {
	text string
	ptr  int
}

func (e *Expression) _ParseExpression() (float64, error) {
	_, isOpen := e.consume([]uint8{'('})
	lValue, err := e._ParseTerm()
	if err != nil {
		return 0, err
	}

	for operator, ok := e.consume([]uint8{'+', '-'}); ok; operator, ok = e.consume([]uint8{'+', '-'}) {
		rValue, err := e._ParseTerm()
		if err != nil {
			return 0, err
		}

		lValue, err = execute(lValue, string(operator), rValue)
		if err != nil {
			return 0, err
		}
	}

	_, isClosed := e.consume([]uint8{')'})
	if isOpen != isClosed {
		return 0, ParseError{"There is not closing bracket", e.ptr}
	}

	return lValue, nil
}

func (e *Expression) _ParseTerm() (float64, error) {
	lValue, err := e._ParseFactor()
	if err != nil {
		return 0, err
	}

	for operator, ok := e.consume([]uint8{'*', '/'}); ok; operator, ok = e.consume([]uint8{'*', '/'}) {
		rValue, err := e._ParseFactor()
		if err != nil {
			return 0, err
		}

		lValue, err = execute(lValue, string(operator), rValue)
		if err != nil {
			return 0, err
		}
	}

	return lValue, nil
}

func (e *Expression) _ParseFactor() (float64, error) {
	e.skipSpaces()
	if e.nextChar() == '(' {
		return e._ParseExpression()
	} else {
		return e.consumeOperand()
	}
}
