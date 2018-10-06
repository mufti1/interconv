package interconv

import "errors"

// ParseInt function to convert other type data to int
func ParseInt(val interface{}) (int, error) {
	number, ok := val.(int)
	if ok == false {
		return -1, errors.New("Cannot convert to int")
	}
	return number, nil
}
