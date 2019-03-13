package interconv_test

import (
	"encoding/json"
	"reflect"
	"testing"

	interconv "github.com/mufti1/interconv/package"
)

func TestParseInt(t *testing.T) {
	t.Run("Json Number", func(t *testing.T) {
		var i interface{} = json.Number("20")
		convertedNum, _ := interconv.ParseInt(i)
		if convertedNum != 20 {
			t.Fatalf("return must be 20.4 instead of %v", convertedNum)
		}
	})
	t.Run("NOK", func(t *testing.T) {
		var i interface{} = 20
		_, err := interconv.ParseInt(i)
		if err == nil {
			t.Fatalf("it should be error: %v", err)
		}
	})
	t.Run("should be int", func(t *testing.T) {
		var expectedDataType = reflect.Int
		var i interface{} = 20

		val, err := interconv.ParseInt(i)
		actualDataType := reflect.TypeOf(val).Kind()

		if err == nil {
			t.Fatalf("it should be error: %v", err)
		}

		if actualDataType != expectedDataType {
			t.Fatalf("Expects to be type of %v, but got type of %v.", expectedDataType, actualDataType)
		}
	})

}

func TestParseInt8(t *testing.T) {
	t.Run("Json Number", func(t *testing.T) {
		var i interface{} = json.Number("20")
		convertedNum, _ := interconv.ParseInt8(i)
		if convertedNum != 20 {
			t.Fatalf("return must be 20.4 instead of %v", convertedNum)
		}
	})
	t.Run("NOK", func(t *testing.T) {
		var i interface{} = 20
		_, err := interconv.ParseInt8(i)
		if err == nil {
			t.Fatalf("it should be error: %v", err)
		}
	})
	t.Run("should be int", func(t *testing.T) {
		var expectedDataType = reflect.Int8
		var i interface{} = 20

		val, err := interconv.ParseInt8(i)
		actualDataType := reflect.TypeOf(val).Kind()

		if err == nil {
			t.Fatalf("it should be error: %v", err)
		}

		if actualDataType != expectedDataType {
			t.Fatalf("Expects to be type of %v, but got type of %v.", expectedDataType, actualDataType)
		}
	})
}

func TestParseInt16(t *testing.T) {
	t.Run("Json Number", func(t *testing.T) {
		var i interface{} = json.Number("20")
		convertedNum, _ := interconv.ParseInt16(i)
		if convertedNum != 20 {
			t.Fatalf("return must be 20.4 instead of %v", convertedNum)
		}
	})
	t.Run("NOK", func(t *testing.T) {
		var i interface{} = 20
		_, err := interconv.ParseInt16(i)
		if err == nil {
			t.Fatalf("it should be error: %v", err)
		}
	})
	t.Run("should be int", func(t *testing.T) {
		var expectedDataType = reflect.Int16
		var i interface{} = 20

		val, err := interconv.ParseInt16(i)
		actualDataType := reflect.TypeOf(val).Kind()

		if err == nil {
			t.Fatalf("it should be error: %v", err)
		}

		if actualDataType != expectedDataType {
			t.Fatalf("Expects to be type of %v, but got type of %v.", expectedDataType, actualDataType)
		}
	})
}
