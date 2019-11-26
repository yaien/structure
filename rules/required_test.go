package rules

import (
	"reflect"
	"testing"

	"github.com/yaien/structure/core"
)

func TestRequired(t *testing.T) {
	testcases := []struct {
		value interface{}
		pass  bool
	}{
		{value: "", pass: false},
		{value: 0, pass: false},
		{value: nil, pass: false},
		{value: 1, pass: true},
		{value: 0.0, pass: false},
		{value: "x", pass: true},
	}

	validate, err := Required([]string{})

	if err != nil {
		t.Fatal(err)
	}

	for _, testcase := range testcases {
		item := &core.Item{Value: reflect.ValueOf(testcase.value)}
		err := validate(item)
		pass := err == nil
		if pass != testcase.pass {
			t.Errorf("expected required validation for value '%v' pass to be '%v',", testcase.value, testcase.pass)
		}
	}
}
