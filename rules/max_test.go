package rules

import (
	"reflect"
	"testing"

	"github.com/yaien/structure/core"
)

func TestMaxRule(t *testing.T) {

	testcases := []struct {
		value reflect.Value
		fail  bool
	}{
		{value: reflect.ValueOf(8), fail: true},
		{value: reflect.ValueOf(6), fail: false},
		{value: reflect.ValueOf(3), fail: false},
		{value: reflect.ValueOf(6.1), fail: true},
		{value: reflect.ValueOf(5.5), fail: false},
		{value: reflect.ValueOf("minie"), fail: false},
		{value: reflect.ValueOf("winnyx"), fail: false},
		{value: reflect.ValueOf("miniemouse"), fail: true},
	}

	validate, err := Max([]string{"6"})

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
			t.Errorf("expected validation failed 'max:6' to be ['%v'] for value %v", testcase.fail, testcase.value)
		}
	}
}
