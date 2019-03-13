package interconv

// ParseInt function to convert other type data to int
func ParseInt(val interface{}) (int, error) {
	number, err := ParseFloat64(val)
	if err != nil {
		return 0, err
	}
	return int(number), nil
}

// ParseInt8 function to convert other type data to int8
func ParseInt8(val interface{}) (int8, error) {
	number, err := ParseFloat64(val)
	if err != nil {
		return int8(0), err
	}
	return int8(number), nil
}

// ParseInt16 function to convert other type data to int16
func ParseInt16(val interface{}) (int16, error) {
	number, err := ParseFloat64(val)
	if err != nil {
		return int16(0), err
	}
	return int16(number), nil
}
