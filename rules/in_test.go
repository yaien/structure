package rules

import (
	"reflect"
	"testing"

	"github.com/yaien/structure/core"
)

func TestIn(t *testing.T) {
	testcases := []struct {
		params []string
		value  interface{}
		valid  bool
	}{
		{params: []string{"red", "white", "blue"}, value: "white", valid: true},
		{params: []string{"2", "3", "4"}, value: 2, valid: true},
		{params: []string{"true", "false"}, value: 0, valid: false},
		{params: []string{"red", "white"}, value: "blue", valid: false},
		{params: []string{"true", "false"}, value: true, valid: true},
		{params: []string{"2.2", "3.2"}, value: 2.2, valid: true},
	}

	for _, testcase := range testcases {
		validate, err := In(testcase.params)
		if err != nil {
			t.Fatal(err)
		}
		item := &core.Item{Value: reflect.ValueOf(testcase.value)}
		err = validate(item)
		valid := err == nil
		if valid != testcase.valid {
			t.Errorf("expected validation for options %v given %v to be '%v', received: '%v'", testcase.params, testcase.value, testcase.valid, valid)
		}
	}
}
