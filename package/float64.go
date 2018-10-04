package interconv

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// ParseFloat64 function to convert other type data to float64
func ParseFloat64(val interface{}) (float64, error) {
	switch val.(type) {
	case nil:
		return 1, nil
	case json.Number:
		return val.(json.Number).Float64()
	default:
		return -1, fmt.Errorf("unable to casting number %v (type %v)", val, reflect.TypeOf(val))
	}
}
