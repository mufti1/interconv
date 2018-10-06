package interconv_test

import (
	"testing"

	interconv "github.com/mufti1/interconv/package"
)

func TestParseInt(t *testing.T) {
	t.Run("NOK", func(t *testing.T) {
		var expectedResult interface{} = 1
		actualResult, err := interconv.ParseInt(expectedResult)
		if err != nil {
			t.Fatalf("Expects to be %v, but got %d", expectedResult, actualResult)
		}
	})

}
