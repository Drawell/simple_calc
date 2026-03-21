package lib

func Evaluate(text string) (float64, error) {
	e := Expression{text, 0}
	value, err := e._ParseExpression()
	e.skipSpaces()
	if err != nil {
		return 0, err
	} else if e.ptr != len(e.text) {
		return 0, ParseError{"Invalid symbols", e.ptr}
	} else {
		return value, err
	}
}
