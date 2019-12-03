package rules

import (
	"reflect"
	"testing"

	"github.com/yaien/structure/core"
)

func TestEqual(t *testing.T) {
	type structure struct {
		Value        interface{}
		Confirmation interface{}
	}
	testcases := []struct {
		data   structure
		target interface{}
		valid  bool
	}{
		{data: structure{"password", "password"}, valid: true},
		{data: structure{"password", "nopassword"}, valid: false},
		{data: structure{2, 3}, valid: false},
		{data: structure{3, 3}, valid: true},
		{data: structure{false, true}, valid: false},
		{data: structure{true, true}, valid: true},
	}

	validate, err := Equal([]string{"Confirmation"})
	if err != nil {
		t.Fatal(err)
	}
	for _, testcase := range testcases {
		item := &core.Item{
			Source: testcase.data,
			Value:  reflect.ValueOf(testcase.data.Value),
		}
		err := validate(item)
		valid := err == nil
		if valid != testcase.valid {
			t.Errorf("expected validation equal for %v to be '%v', received: '%v'",
				testcase.data, testcase.valid, valid)
		}
	}

}
