package interconv_test

import (
	"encoding/json"
	"testing"

	"github.com/mufti1/interconv/package"
)

func TestParseFloat64(t *testing.T) {
	t.Run("Nil input", func(t *testing.T) {
		var i interface{}
		convertedNum, _ := interconv.ParseFloat64(i)
		if convertedNum != 1 {
			t.Fatalf("return must be 1 instead of %v", convertedNum)
		}
	})
	t.Run("Json Number", func(t *testing.T) {
		var i interface{} = json.Number("20.4")
		convertedNum, _ := interconv.ParseFloat64(i)
		if convertedNum != 20.4 {
			t.Fatalf("return must be 20.4 instead of %v", convertedNum)
		}
	})
	t.Run("NOK", func(t *testing.T) {
		var i interface{} = 20.4
		_, err := interconv.ParseFloat64(i)
		if err == nil {
			t.Fatalf("it should be error: %v", err)
		}
	})
}
