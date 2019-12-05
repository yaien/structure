package rules

import (
	"reflect"
	"testing"

	"github.com/yaien/structure/core"
)

func TestPattern(t *testing.T) {
	testcases := []struct {
		pattern string
		value   string
		valid   bool
	}{
		{pattern: "^[a-z]+$", value: "hola", valid: true},
		{pattern: "^[0-9]+$", value: "hola", valid: false},
		{pattern: "^[a-b]+\\.com$", value: "hola", valid: false},
		{pattern: "^[a-z]+\\.com$", value: "hola.com", valid: true},
	}
	for _, testcase := range testcases {
		validate, err := Pattern([]string{testcase.pattern})
		if err != nil {
			t.Error(err)
			continue
		}
		item := &core.Item{Value: reflect.ValueOf(testcase.value)}
		err = validate(item)
		valid := err == nil
		if valid != testcase.valid {
			t.Errorf("expected validation pattern:'%s' with value '%s' to be '%v', received: '%v'",
				testcase.pattern, testcase.value, testcase.valid, valid)
		}
	}
}
