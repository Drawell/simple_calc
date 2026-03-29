package calc

type Expression struct {
	text string
	ptr  int
}

func (e *Expression) parseExpression() (float64, error) {
	_, isOpen := e.consume([]uint8{'('})
	lValue, err := e.parseTerm()
	if err != nil {
		return 0, err
	}

	for operator, ok := e.consume([]uint8{'+', '-'}); ok; operator, ok = e.consume([]uint8{'+', '-'}) {
		rValue, err := e.parseTerm()
		if err != nil {
			return 0, err
		}

		lValue, err = execute(lValue, string(operator), rValue)
		if err != nil {
			return 0, err
		}
	}

	_, isClosed := e.consume([]uint8{')'})
	if isOpen && !isClosed {
		return 0, ParseError{"There is not closing bracket", e.ptr}
	} else if !isOpen && isClosed {
		return 0, ParseError{"There is not open bracket for closing", e.ptr - 1}
	}

	return lValue, nil
}

func (e *Expression) parseTerm() (float64, error) {
	lValue, err := e.parseFactor()
	if err != nil {
		return 0, err
	}

	for operator, ok := e.consume([]uint8{'*', '/'}); ok; operator, ok = e.consume([]uint8{'*', '/'}) {
		rValue, err := e.parseFactor()
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

func (e *Expression) parseFactor() (float64, error) {
	e.skipSpaces()
	if e.nextChar() == '(' {
		return e.parseExpression()
	} else {
		return e.consumeOperand()
	}
}
