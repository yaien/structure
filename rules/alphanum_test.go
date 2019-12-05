package rules

import (
	"reflect"
	"testing"

	"github.com/yaien/structure/core"
)

func TestAlphaNum(t *testing.T) {
	testcases := []struct {
		value string
		valid bool
	}{
		{value: "omgd344", valid: true},
		{value: "omgd", valid: true},
		{value: "OMFDS45", valid: true},
		{value: "OMG", valid: true},
		{value: "dnsonf%55", valid: false},
		{value: "dsDSj8$%#$%", valid: false},
	}

	validate, err := AlphaNum([]string{})
	if err != nil {
		t.Fatal(err)
	}
	for _, testcase := range testcases {
		item := &core.Item{Value: reflect.ValueOf(testcase.value)}
		err := validate(item)
		valid := err == nil
		if valid != testcase.valid {
			t.Errorf("expected validation for '%s' to be '%v', received: '%v'",
				testcase.value, testcase.valid, valid)
		}
	}
}
