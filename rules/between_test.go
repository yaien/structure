package rules

import (
	"reflect"
	"testing"

	"github.com/yaien/structure/core"
)

func TestBetweenRule(t *testing.T) {

	testcases := []struct {
		value reflect.Value
		fail  bool
	}{
		{value: reflect.ValueOf(2), fail: true},
		{value: reflect.ValueOf(8), fail: true},
		{value: reflect.ValueOf(6), fail: false},
		{value: reflect.ValueOf(5), fail: false},
		{value: reflect.ValueOf(3.2), fail: true},
		{value: reflect.ValueOf(6.1), fail: true},
		{value: reflect.ValueOf(5.5), fail: false},
		{value: reflect.ValueOf("minie"), fail: false},
		{value: reflect.ValueOf("winnyx"), fail: false},
		{value: reflect.ValueOf("miniemouse"), fail: true},
		{value: reflect.ValueOf("mi"), fail: true},
	}

	validate, err := Between([]string{"4", "6"})

	if err != nil {
		t.Fatal(err)
	}

	for _, testcase := range testcases {
		item := &core.Item{
			Value: testcase.value,
		}
		err := validate(item)
		fail := err != nil
		if testcase.fail != fail {
			t.Errorf("expected validation failed 'between:4,6' to be '%v' for value '%v'", testcase.fail, testcase.value)
		}
	}
}
