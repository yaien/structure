package rules

import (
	"reflect"
	"testing"

	"github.com/yaien/structure/core"
)

func TestMinRule(t *testing.T) {

	testcases := []struct {
		value reflect.Value
		fail  bool
	}{
		{value: reflect.ValueOf(4), fail: true},
		{value: reflect.ValueOf(6), fail: false},
		{value: reflect.ValueOf(8), fail: false},
		{value: reflect.ValueOf(6.1), fail: false},
		{value: reflect.ValueOf("minie"), fail: true},
		{value: reflect.ValueOf("miniex"), fail: false},
		{value: reflect.ValueOf("miniemouse"), fail: false},
	}

	validate, err := Min([]string{"6"})

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
			t.Errorf("expected validation failed 'min:6' to be ['%v'] for value %v", testcase.fail, testcase.value)
		}
	}
}
